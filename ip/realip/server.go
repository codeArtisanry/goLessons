package main

import (
	"log"
	"net/http"

	"github.com/Tomasen/realip"
	"github.com/julienschmidt/httprouter"
)

func (h *Handler) ServeIndexPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clientIP := realip.FromRequest(r)
	log.Println("GET / from", clientIP)
}
