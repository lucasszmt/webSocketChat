package main

import (
	"fmt"
	"github.com.br/lucasszmt/webSocketChat/cmd/server"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

func main() {
	//start the hub that controls the messaging and the clients that uses the service
	hub := server.NewHub()
	go hub.Run()

	router := gin.New()
	router.GET("/", func(context *gin.Context) {
		home(context.Writer, context.Request)
	})
	router.GET("/ws", func(context *gin.Context) {
		server.ServeWebSocket(context.Writer, context.Request, hub)
	})
	router.GET("/clients", func(context *gin.Context) {
		connections := make(map[string]bool)
		for key, _ := range hub.Clients {
			connections[key.Conn.RemoteAddr().String()] = true
		}
		fmt.Println(connections)
		context.JSON(200, connections)
	})

	if err := router.Run(":8080"); err != nil {
		log.Println(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/ws")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
