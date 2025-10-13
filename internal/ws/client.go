package ws

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		if _, _, err := c.Conn.ReadMessage(); err != nil {
			break
		}
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()
	for msg := range c.Send {
		c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			return
		}
	}

	// When channel is closed, send close message to client
	_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}
