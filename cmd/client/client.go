package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

var addr = url.URL{
	Scheme: "ws",
	Host:   "0.0.0.0:8080",
	Path:   "/ws",
}

func connectToWebSocket(addr url.URL) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(addr.String(), nil)
	if err != nil {
		log.Fatal("Dial:", err)
	}
	return conn, err
}