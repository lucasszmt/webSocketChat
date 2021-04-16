package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	u := url.URL{
		Scheme: "ws",
		Host:   "0.0.0.0:8080",
		Path:   "/ws",
	}
	client := websocket.DefaultDialer
	conn, _, err := client.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.SetPingHandler(func(string) error {
		fmt.Print("Recebi um ping, enviando um pong!")
		err := conn.WriteMessage(websocket.PongMessage, []byte{})
		if err != nil {
			return err
		}
		return nil
	})
	go func() {
		mt, msg, err := conn.ReadMessage()
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
		err = conn.WriteMessage(websocket.TextMessage, line)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
