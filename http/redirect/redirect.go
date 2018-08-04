package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://www.baidu.com", 301)
	http.Redirect(w, r, "/", 302)
}

func main() {

	redirectHandler := http.RedirectHandler("https://307.temporaryredirect.com", http.StatusTemporaryRedirect)
	http.Handle("/307", redirectHandler)

	http.Handle("/301", http.RedirectHandler("/hello", http.StatusMovedPermanently))

	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
