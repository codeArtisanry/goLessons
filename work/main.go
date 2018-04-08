package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"

	"github.com/astaxie/beego/config"
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
	"github.com/mvdan/xurls"
	"github.com/zpnk/go-bitly/bitly"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

type Sms struct {
	Number string
	Text   string
}

// read the config file
var iniconf, _ = config.NewConfig("ini", "../conf/app.conf")

func main() {
	// Make a new pool,
	// Context{} is a struct that will be the context for the request.
	pool := work.NewWorkerPool(Sms{}, 10, "burstsms_sms", redisPool)

	// Map the name of jobs to handler functions
	pool.Job("send_sms", (*Sms).SendSms)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	pool.Stop()
}

func (s *Sms) post(number string, text string) string {
	client := &http.Client{}
	URL := "https://api.transmitsms.com/send-sms.json"
	v := url.Values{"to": {number}, "message": {text}}
	//pass the values to the request's body
	req, _ := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(iniconf.String("burstsms_api_key"), iniconf.String("burstsms_api_sec"))
	resp, _ := client.Do(req)
	bodyText, _ := ioutil.ReadAll(resp.Body)
	result := string(bodyText)
	return result
}

func bitly_url(url string) string {
	client := bitly.NewClient(iniconf.String("bitly_auth_token"))
	link := &bitly.Link{client}
	short_link, err := link.Lookup(url)
	if err != nil {
		log.Fatal(err)
	}
	return short_link[0].AggregateLink
}

func (s *Sms) find_and_short_urls() {
	urls := xurls.Strict.FindAllString(s.Text, -1)
	for _, element := range urls {
		s.Text = strings.Replace(s.Text, element, bitly_url(element), -1)
	}
}

func (s *Sms) SendSms(job *work.Job) error {
	fmt.Println("Starting job...")
	if _, ok := job.Args["Number"]; ok {
		s.Number = job.ArgString("Number")
		s.Text = job.ArgString("Text")
		//s.find_and_short_urls()
		result := s.post(s.Number, s.Text)
		fmt.Println(result)
	} else {
		log.Fatal("Couldn't parse the job data!")
	}
	return nil
}
