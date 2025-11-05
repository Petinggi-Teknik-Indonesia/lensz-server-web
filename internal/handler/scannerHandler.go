package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"lensz-server-web/internal/ws"
	"github.com/gin-gonic/gin"
	"sync"
)

type ScannerHandler struct {
	hub *ws.Hub
	isRegistering bool
	mu sync.Mutex
}

const registerTimeout = 30 * time.Second // example duration

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

	h.mu.Lock()
	if h.isRegistering {
		h.mu.Unlock()
		c.JSON(http.StatusConflict, gin.H{"message": "Registration in progress"})
		return
	}
	h.isRegistering = true
	h.mu.Unlock()

	msg := map[string]interface{}{
		"type": "rfid_scanned",
		"payload": map[string]string{
			"rfid": body.RFID,
		},
	}
	b, _ := json.Marshal(msg)
	h.hub.Broadcast <- b

	// start auto timeout watcher
	go h.autoTimeout()

	c.JSON(http.StatusOK, gin.H{"message": "Scan received"})
}

func (h *ScannerHandler) autoTimeout() {
	time.Sleep(registerTimeout)
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.isRegistering {
		h.isRegistering = false
		// optionally tell FE
		msg := map[string]string{"type": "registration_timeout"}
		b, _ := json.Marshal(msg)
		h.hub.Broadcast <- b
	}
}

func (h *ScannerHandler) CompleteRegistration(c *gin.Context) {
	h.mu.Lock()
	h.isRegistering = false
	h.mu.Unlock()
	c.JSON(http.StatusOK, gin.H{"message": "Registration completed"})
}

func (h *ScannerHandler) CancelRegistration(c *gin.Context) {
	h.mu.Lock()
	h.isRegistering = false
	h.mu.Unlock()
	c.JSON(http.StatusAccepted, gin.H{"message": "Registration canceled"})
}
