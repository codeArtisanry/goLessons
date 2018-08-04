/*
File      : httpShare.go
Author    : Mike
E-Mail    : Mike_Zhang@live.com
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func assetsHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./"+r.URL.Path)
	fmt.Println(r.URL.Path)
}

func main() {
	http.HandleFunc("/", assetsHandler)
	//file server
	http.Handle("/files/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
