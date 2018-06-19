package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// 假设该URL没有#fragment后缀
	values, err := url.ParseRequestURI("https://www.baidu.com/s?wd=%E6%90%9C%E7%B4%A2&rsv_spt=1&issp=1&f=8&rsv_bp=0&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_sug3=7&rsv_sug1=6")
	fmt.Println(values)
	u, err := url.Parse("http://bing.com/search/!()*[]<>?q=!()*;[]<>$&+=:#home")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url.PathEscape(u.RawPath))     // %2Fsearch%2F%21%28%29%2A%5B%5D%3C%3E
	fmt.Println(url.PathUnescape(u.RawPath))   // /search/!()*[]<> <nil>
	fmt.Println(url.QueryEscape(u.RawQuery))   // q%3D%21%28%29%2A%3B%5B%5D%3C%3E%24%26%2B%3D%3A
	fmt.Println(url.QueryUnescape(u.RawQuery)) // q=!()*;[]<>$& =: <nil>
	fmt.Println(u)                             //http://bing.com/search/%21%28%29%2A%5B%5D%3C%3E?q=!()*;[]<>$&+=:#home
	fmt.Println(u.Scheme)                      //http
	fmt.Println(u.Opaque)                      //
	fmt.Println(u.Host)                        //bing.com
	fmt.Println(u.Hostname())                  //bing.com
	fmt.Println(u.Port())                      //
	fmt.Println(u.Path)                        ///search/!()*[]<>
	fmt.Println(u.RawPath)                     ///search/!()*[]<>
	fmt.Println(u.ForceQuery)                  // false
	fmt.Println(u.RawQuery)                    //q=!()*;[]<>$&+=:
	fmt.Println(u.Fragment)                    //home
	fmt.Println(u)
	u.RawQuery = u.Query().Encode()
	fmt.Println(u)
}
