package main

import (
	"log"
	"regexp"
	"time"

	"github.com/imroc/req"
)

/**
 * 将域名解析Ip地址
 * 获得域名对应的所有Ip地址
 */
var regip = regexp.MustCompile("\\d+\\.\\d+\\.\\d+\\.\\d+")
var ms float32
var remote_ip string

// func va(ip, port string) {
// 	urli := url.URL{}
// 	urlproxy, _ := urli.Parse("https://" + ip + ":" + port)
// 	client := http.Client{
// 		Transport: &http.Transport{
// 			Proxy:           http.ProxyURL(urlproxy),
// 			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 		},
// 		Timeout: time.Millisecond * 500,
// 	}

// 	tr := &http.Transport{
// 		Proxy:           http.ProxyURL(proxyUrl),
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}
// 	client := &http.Client{
// 		Transport: tr,
// 		Timeout:   time.Millisecond * 500,
// 	}

// 	t1 := time.Now().UnixNano()

// 	resp, err := client.Get("https://myip.ipip.net")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	t2 := time.Now().UnixNano()
// 	ms = float32((t2 - t1) / 1000.0 / 1000.0)
// 	log.Println("time", ms)

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Printf("%s\n", body)

// 	ips := regip.FindAllString(string(body), 1)
// 	if len(ips) == 1 {
// 		remote_ip = ips[0]
// 		log.Println(remote_ip)
// 		err = nil
// 		return
// 	}

// 	resp.Body.Close()

// }
func main() {
	// va("185.89.217.99", "52065")
	req.SetTimeout(10 * time.Second)
	req.SetProxyUrl("http://115.166.128.230:65309")
	t1 := time.Now().UnixNano()
	// r, err := req.Get("http://myip.ipip.net")
	r, err := req.Get("http://ifconfig.me/ip")
	// r, err := req.Get("http://curlmyip.com")
	// r, err := req.Get("http://httpbin.org/ip")
	t2 := time.Now().UnixNano()
	if err != nil {
		return
	}
	ms = float32((t2 - t1) / 1000.0 / 1000.0)
	log.Println("time", ms)
	log.Println("body", r.String())
	ips := regip.FindAllString(r.String(), 1)
	if len(ips) == 1 {
		remote_ip = ips[0]
		log.Println(remote_ip)
		err = nil
		return
	}
	return

}
