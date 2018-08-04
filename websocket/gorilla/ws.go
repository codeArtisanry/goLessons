package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":8000", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("websocket error")
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	go echo(conn)
}

func echo(conn *websocket.Conn) {
	p := []byte("12")
	message := 1
	timer := time.NewTicker(1 * time.Second)
	for {
		<-timer.C
		if err := conn.WriteMessage(message, p); err != nil {
			return
		}
	}
}
