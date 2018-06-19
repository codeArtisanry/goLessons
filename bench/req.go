package main

import (
	"log"
	"net/http"
	"time"

	"github.com/imroc/req"
	"github.com/mozillazg/request"
)

func main() {

	//size is 26MB
	start := time.Now()
	c := new(http.Client)
	reqs := request.NewRequest(c)
	resp, _ := reqs.Get("http://httpbin.org/get")
	// resp.Json()
	resp.Body.Close()

	timeCommon := time.Now()

	// test for github.com/imroc/req

	r := req.New()
	// r.SetFlags(req.LstdFlags | req.Lcost)
	resp2, _ := r.Get("http://ip.cn/?uyxelc=yxth33")

	timeCommon2 := time.Now()

	log.Printf("read common cost time %v\n", timeCommon.Sub(start))
	log.Printf("read common cost time %v\n", timeCommon2.Sub(timeCommon))
	log.Println(resp2.Cost())

	log.Println(resp2.Dump())
	// mux := http.NewServeMux()
	// mux.Handle("/", http.RedirectHandler("http://baidu.com", 301))
	// http.ListenAndServe(":8080", mux)
}
