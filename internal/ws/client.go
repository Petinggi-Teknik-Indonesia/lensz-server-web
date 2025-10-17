package ws

import (
	"time"
	"log"
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadPump() {
	defer func() {
		log.Println("📴 Client disconnected (ReadPump)")
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		if _, _, err := c.Conn.ReadMessage(); err != nil {
			log.Println("⚠️ ReadPump error:", err)
			break
		}
	}
}

func (c *Client) WritePump() {
	defer func() {
		log.Println("📴 Client disconnected (WritePump)")
		c.Conn.Close()
	}()
	for msg := range c.Send {
		log.Println("📨 Sending message to client:", string(msg))
		c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("⚠️ WritePump error:", err)
			return
		}
	}
	_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}
