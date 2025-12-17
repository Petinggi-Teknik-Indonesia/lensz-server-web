package handler

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"lensz-server-web/internal/ws"
	"github.com/gin-gonic/gin"
)

/* =======================
   EXISTING CODE (UNCHANGED)
======================= */

type ScannerHandler struct {
	hub           *ws.Hub
	isRegistering bool
	mu            sync.Mutex
}

const registerTimeout = 30 * time.Second

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

	go h.autoTimeout()

	c.JSON(http.StatusOK, gin.H{"message": "Scan received"})
}

func (h *ScannerHandler) autoTimeout() {
	time.Sleep(registerTimeout)
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.isRegistering {
		h.isRegistering = false
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

/* =======================
   HEARTBEAT (IN-MEMORY)
======================= */

type heartbeatData struct {
	LastSeen time.Time
}

var (
	heartbeatStore = make(map[uint]heartbeatData)
	heartbeatMu    sync.RWMutex
)

/*
POST /api/scanner/heartbeat
REQ: { "scanner_id": 1 }
*/
func (h *ScannerHandler) Heartbeat(c *gin.Context) {
	var body struct {
		ScannerID uint `json:"scanner_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	heartbeatMu.Lock()
	heartbeatStore[body.ScannerID] = heartbeatData{
		LastSeen: time.Now(),
	}
	heartbeatMu.Unlock()

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

/*
GET /api/scanner/devices
- FULL IN MEMORY
- NO DB
*/
func (h *ScannerHandler) GetDevices(c *gin.Context) {
	now := time.Now()
	result := []gin.H{}

	heartbeatMu.RLock()
	for id, hb := range heartbeatStore {
		status := "dead"
		if now.Sub(hb.LastSeen) <= 45*time.Second {
			status = "alive"
		}

		result = append(result, gin.H{
			"scanner_id": id,
			"status":     status,
			"last_seen":  hb.LastSeen,
		})
	}
	heartbeatMu.RUnlock()

	c.JSON(http.StatusOK, result)
}

/* =======================
   OPTIONAL CLEANER
======================= */

func StartHeartbeatCleaner() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			heartbeatMu.Lock()
			for id, hb := range heartbeatStore {
				if time.Since(hb.LastSeen) > 10*time.Minute {
					delete(heartbeatStore, id)
				}
			}
			heartbeatMu.Unlock()
		}
	}()
}
