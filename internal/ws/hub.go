package ws

import "sync"

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client] = true
			h.mu.Unlock()
		case client := <-h.Unregister:
			h.mu.Lock()
			delete(h.Clients, client)
			close(client.Send)
			h.mu.Unlock()
		case msg := <-h.Broadcast:
			h.mu.Lock()
			for c := range h.Clients {
				select {
				case c.Send <- msg:
				default:
					delete(h.Clients, c)
					close(c.Send)
				}
			}
			h.mu.Unlock()
		}
	}
}
