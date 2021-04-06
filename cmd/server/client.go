package server

import "github.com/gorilla/websocket"

//Client represents a chat client
type Client struct {
	Hub *Hub

	Conn *websocket.Conn

	Send chan []byte
}

func NewClient(hub *Hub, conn *websocket.Conn, send chan []byte) *Client {
	return &Client{
		Hub:  hub,
		Conn: conn,
		Send: send,
	}
}
