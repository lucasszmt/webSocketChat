package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ServeWebsocket(c *gin.Context, h *Hub) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	client := &Client{conn: conn, Hub: h}
	go func() {
		err := client.ReadPump()
		if err != nil {
			conn.Close()
			log.Println(err)
			return
		}
	}()
	go client.Ping()
}
