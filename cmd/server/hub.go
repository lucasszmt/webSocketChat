package server

import (
	"fmt"
	"log"
)

type Hub struct {
	Clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.Clients[client] = true
		case client := <-hub.unregister:
			if _, ok := hub.Clients[client]; ok {
				if err := client.Conn.Close(); err != nil {
					fmt.Println("Error closing connection: ", err)
				}
				delete(hub.Clients, client)
				close(client.Send)
			}
		case message := <-hub.broadcast:
			log.Fatal(message, "Unimplemented")
		case message := <-hub.broadcast:
			for client := range hub.Clients {
				select {
				case client.Send <- message:
				//the default case closes the connection, if it reaches the max buffer
				default:
					close(client.Send)
					delete(hub.Clients, client)
				}
			}

		}

	}
}
