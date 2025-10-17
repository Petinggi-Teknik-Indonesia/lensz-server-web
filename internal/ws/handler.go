package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("🔌 New WebSocket connection attempt...")
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("❌ Upgrade failed:", err)
			return
		}

		client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
		hub.Register <- client

		log.Println("✅ WebSocket connected:", conn.RemoteAddr())

		go client.WritePump()
		go client.ReadPump()
	}
}
