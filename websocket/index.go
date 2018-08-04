package main

import (
	"io/ioutil"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("views/index.html")
	check(err)
	_, err = w.Write(dat)
	check(err)
}
