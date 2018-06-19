package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/", handler)
	go http.ListenAndServe(":80", nil)
	// http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	for {
	}
}
