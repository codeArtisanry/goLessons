package main

import (
	socks5 "github.com/armon/go-socks5"
)

var conf = new(socks5.Config)

func main() {
	// conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe("tcp", "0.0.0.0:8888"); err != nil {
		panic(err)
	}
}
