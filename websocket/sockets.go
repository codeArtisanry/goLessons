package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	var messageType int = 1
	var p []byte
	// construction of Connection.Id totally has a race condition, ignore for now
	thisConnection := Connection{len(h.Connections), conn, make(chan []byte, 1), make(chan []byte, 1)}
	h.register(thisConnection)
	for {
		go receiveSomething(conn, thisConnection)
		select {
		case p = <-thisConnection.Inbound:
			h.broadcastOthers(thisConnection, p)
		case receives := <-thisConnection.Outbound:
			conn.WriteMessage(messageType, receives)
		}
	}
}

func receiveSomething(conn *websocket.Conn, thatConnection Connection) {
	var p []byte
	var err error
	_, p, err = conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	thatConnection.Inbound <- p
}
