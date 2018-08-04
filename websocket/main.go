package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/assets/", assetsHandler)
	http.HandleFunc("/ws", wsHandler)

	log.Println("Server started. Port=" + app.Port()[1:])
	err := http.ListenAndServe(app.Port(), nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
