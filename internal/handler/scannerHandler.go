package handler

import (
	"encoding/json"
	"net/http"

	"lensz-server-web/internal/ws"
	"github.com/gin-gonic/gin"
)

type ScannerHandler struct {
	hub *ws.Hub
	// optionally keep state
	isRegistering bool
}

func NewScannerHandler(hub *ws.Hub) *ScannerHandler {
	return &ScannerHandler{hub: hub}
}

func (h *ScannerHandler) Scan(c *gin.Context) {
	var body struct {
		RFID string `json:"rfid" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Option 1: block new scans if registration in progress
	if h.isRegistering {
		c.JSON(http.StatusConflict, gin.H{"message": "Registration in progress"})
		return
	}

	// Notify frontends
	msg := map[string]interface{}{
		"type": "rfid_scanned",
		"payload": map[string]string{
			"rfid": body.RFID,
		},
	}
	b, _ := json.Marshal(msg)
	h.hub.Broadcast <- b

	h.isRegistering = true

	c.JSON(http.StatusOK, gin.H{"message": "Scan received"})
}

// endpoint for frontend to mark registration done
func (h *ScannerHandler) CompleteRegistration(c *gin.Context) {
	h.isRegistering = false
	c.JSON(http.StatusOK, gin.H{"message": "Registration completed"})
}
