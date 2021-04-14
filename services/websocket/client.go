package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Client struct {
	conn *websocket.Conn
	Hub  *Hub
}

func (c Client) ReadPump() error {
	for {
		_, msg, msgerr := c.conn.ReadMessage()
		if msgerr != nil {
			log.Println("Erro de leitura")
			return msgerr
		}
		fmt.Println(string(msg))
		err := c.conn.SetReadDeadline(time.Now().Add(time.Second * 1))
		if err != nil {
			log.Println("TIME OVER")
			return err
		}
	}
}
