package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
)

type Client struct {
	url    string
	conn   *websocket.Conn
	dialer *websocket.Dialer
}

func NewClient(url url.URL, dialer *websocket.Dialer) *Client {
	c := &Client{url: url.String(), dialer: dialer}
	c.Dial()
	return c
}

func (c *Client) Dial() error {
	var err error
	c.conn, _, err = c.dialer.Dial(c.url, nil)
	return err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	u := url.URL{
		Scheme: "ws",
		Host:   "0.0.0.0:8080",
		Path:   "/ws",
	}
	client := NewClient(u, websocket.DefaultDialer)
	client.conn.SetPingHandler(func(string) error {
		fmt.Print("Recebi um ping, enviando um pong!")
		err := client.conn.WriteMessage(websocket.PongMessage, []byte{})
		if err != nil {
			return err
		}
		return nil
	})
	go func() {
		mt, msg, err := client.conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("MT:%d MESSAGE:%s \n", mt, msg)
	}()
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		err = client.conn.WriteMessage(websocket.TextMessage, line)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
