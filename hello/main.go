package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("New connection from %s", r.RemoteAddr)
	w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
