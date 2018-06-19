package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	fmt.Println(req.RequestURI)
	fmt.Println(req.Body)
	w.Write([]byte("Hello\n"))
}

func main() {
	http.HandleFunc("/api/download", hello)
	http.ListenAndServe(":8001", nil)
}
