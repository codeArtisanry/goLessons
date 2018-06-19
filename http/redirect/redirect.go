package main

import (
    "log"
    "net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {

    http.Redirect(w, r, "http://www.baidu.com", 301)
}

func main() {
    http.HandleFunc("/", redirect)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}