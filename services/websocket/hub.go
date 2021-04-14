package websocket

type Hub struct {
	Clients    map[*Client]string
	Register   chan *Client
	Unregister chan *Client
}

func (h Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = client.conn.RemoteAddr().String()
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
			}
		}

	}
}
