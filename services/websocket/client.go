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
	c.Pong()
	for {
		err := c.conn.SetReadDeadline(time.Now().Add(time.Second * 60))
		if err != nil {
			log.Println("TIME OVER")
			return err
		}
		_, msg, msgerr := c.conn.ReadMessage()
		if msgerr != nil {
			log.Println("Erro de leitura")
			return msgerr
		}
		fmt.Println(string(msg))
	}
}

func (c Client) Pong() {
	c.conn.SetPongHandler(func(appData string) error {
		log.Println("Recebi pong")
		c.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		return nil
	})
}

func (c Client) Ping() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			log.Println("Pingueango!")
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
