package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

var addr = url.URL{
	Scheme: "ws",
	Host:   "0.0.0.0:8080",
	Path:   "/ws",
}

func connectToWebSocket(addr url.URL) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(addr.String(), nil)
	if err != nil {
		log.Fatal("Dial:", err)
	}
	return conn
}

func main() {
	conn := connectToWebSocket(addr)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(msg))
		time.Sleep(time.Minute)
	}
}
