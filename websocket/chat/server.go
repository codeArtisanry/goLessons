package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"io/ioutil"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var core = &Core{make(map[*websocket.Conn]*Client), sync.Mutex{}}

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/publish", servePublish)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	//w.Write([]byte("Hello"))
	http.ServeFile(w, r, "index.html")
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	// 1. Security: Client Validation
	// origin := r.Header.Get("Origin")
	// w.Header().Set("Access-Control-Allow-Origin", origin)
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// 2. Open socket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	core.Register(conn, "1")
	defer core.Remove(conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func servePublish(w http.ResponseWriter, r *http.Request) {
	// 1. Security: Validate sender
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// 2. Broadcast
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, c := range core.Clients {
		if c.UserID == "1" {
			c.Conn.WriteMessage(websocket.TextMessage, b)
		}
	}
}

type Client struct {
	Conn   *websocket.Conn
	UserID string
}

type Core struct {
	Clients map[*websocket.Conn]*Client
	lock    sync.Mutex
}

func (c *Core) Register(conn *websocket.Conn, userID string) error {
	c.lock.Lock()
	c.Clients[conn] = &Client{conn, userID}
	c.lock.Unlock()

	fmt.Println("Client subscribed")
	return nil
}

func (C *Core) Remove(conn *websocket.Conn) {
	defer conn.Close()
	delete(core.Clients, conn)

	fmt.Println("Client unsubscribed")
}
