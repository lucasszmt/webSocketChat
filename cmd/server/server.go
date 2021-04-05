package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//upgrader of the endpoint, change the options as needed
var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//ServeWebSocket Upgrades de desired route, to an websocket listening point
//use it in any route you'd like
func ServeWebSocket(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatal("Upgrade error:", err)
	}
	defer conn.Close()


}

