package main

// 188.24.51.81 - - [01/Feb/2018:14:49:16 CST] "GET http://udn-plus.cedexis-test.com/img/35062/iuni2.html?rnd=-1-1-13960-0-0-35062-3705136164-_CgJqMRAUGEYiBQgBEIhtKKTI3-YNMJvPXDjp8MrTBUDW_tnzDEoQCAMQtAEYhEQgACirjoCgBFAAWgoIABAAGAAgACgAYABqGmJ1dHRvbi13b3JrZXIyLmFtcy5odi5wcm9kggEQCAMQtAEYhEQgACiwjoCgBIgBvLHMiQU HTTP/1.1" 200 0 1008 1412 "http://stardust-rain.tumblr.com/ask" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:57.0) Gecko/20100101 Firefox/57.0" "-" "-" "-" "-" LLNW
// 195.142.179.194 - - [01/Feb/2018:14:54:25 CST] "GET http://udn-plus.cedexis-test.com/img/35935/r20.gif?rnd=0-1-13960-0-0-35935-2572944071-_CgJqMRAUGEYiBQgBEIhtKMeF8MoJMJzPXDie88rTBUDc2MgyShEIBBDWARiokQIgACiwkYCgBFAAWgoIABAAGAAgACgAYABqGmJ1dHRvbi13b3JrZXIxLmFtcy5odi5wcm9kggERCAQQ1gEYqJECIAAosJGAoASIAcOdzpUM HTTP/1.1" 200 0 43 445 "http://bigboy1977.tumblr.com/post/153677022974" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36" "-" "-" "-" "-" LLNW

package main 

import (
    "fmt"
    "regexp"
    "os"
    "bufio"
)

type myRegexp struct {
    *regexp.Regexp
}

func (r *myRegexp) FindStringSubmatchMap(s string) map[string]string {
    captures := make(map[string]string)

    match := r.FindStringSubmatch(s)
    if match == nil {
        return captures
    }

    for i, name := range r.SubexpNames() {
        // 
        if i == 0 {
            continue
        }
        captures[name] = match[i]

    }
    return captures
}

func main() {
    re2str := `^(?P<remote_addr>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) .* ` + 
              `\[(?P<time_local>.*?)\] ` +
              `"(?P<request>.*?)" ` + 
              `(?P<status>[^ ]*) `  + 
              `(?P<request_time>[^ ]*) ` +
              `(?P<body_bytes_sent>[^ ]*) ` +
              `(?P<bytes_sent>[^ ]*) ` +
              `"(?P<http_referer>[^"]*)" ` +
              `"(?P<http_user_agent>[^"]*)" ` +
              `"(?P<http_x_forwarded_for>[^"]*)" ` +
              `(?P<connection>[^ ]*) ` +
              `"(?P<hit>[^"]*)" ` +
              `"(?P<server_addr>[^"]*)" ` +
              `(?P<cdn>.*)` 

    re2 := myRegexp{regexp.MustCompile(re2str)}

    for i := 0; i< 10; i++ {
        inFile, err := os.Open("test.log")

        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        } else {
            defer inFile.Close()
        }

        scanner := bufio.NewScanner(inFile)
        scanner.Split(bufio.ScanLines)       
        for scanner.Scan() {
              line := scanner.Text()
              re2.FindStringSubmatchMap(line)
        }
    }
}