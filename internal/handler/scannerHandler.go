package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"lensz-server-web/internal/service"
	"lensz-server-web/internal/ws"
	"sync"

	"github.com/gin-gonic/gin"
)

type ScannerHandler struct {
	hub               *ws.Hub
	isRegistering     bool
	mu                sync.Mutex
	service           *service.ScannerService
	drawerScanService *service.DrawerScanService
	heartbeatStore *service.ScannerHeartbeatStore

	registrationRFID string
	lastHeartbeatAt  time.Time

	searchCancel context.CancelFunc
	searchRFID   string
}


func NewScannerHandler(
	hub *ws.Hub,
	scannerService *service.ScannerService,
	drawerScanService *service.DrawerScanService,
	heartbeat *service.ScannerHeartbeatStore,
) *ScannerHandler {
	return &ScannerHandler{
		hub:               hub,
		service:           scannerService,
		drawerScanService: drawerScanService,
		heartbeatStore:    heartbeat,
	}
}

func (h *ScannerHandler) Register(c *gin.Context) {
	var req struct {
		RFID       string `json:"rfid" binding:"required"`
		DeviceName string `json:"deviceName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.GetByName(c, req.DeviceName); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid device"})
		return
	}

	h.mu.Lock()
	if h.isRegistering {
		h.mu.Unlock()
		c.JSON(http.StatusConflict, gin.H{"message": "Registration in progress"})
		return
	}

	h.isRegistering = true
	h.registrationRFID = req.RFID
	h.lastHeartbeatAt = time.Now()
	h.mu.Unlock()

	// 🔴 broadcast registration started
	h.broadcastRegistration("registration_started")

	// 🔴 start WS-only watchdog
	go h.registrationWatchdog()

	c.JSON(http.StatusOK, gin.H{"message": "Registration started"})
}

func (h *ScannerHandler) broadcastRegistration(event string) {
	msg := map[string]interface{}{
		"type": event,
		"payload": gin.H{
			"rfid": h.registrationRFID,
		},
	}

	b, _ := json.Marshal(msg)
	h.hub.Broadcast <- b
}

func (h *ScannerHandler) registrationWatchdog() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		h.mu.Lock()
		if !h.isRegistering {
			h.mu.Unlock()
			return
		}

		if time.Since(h.lastHeartbeatAt) > 30*time.Second {
			h.cancelRegistration("registration_timeout")
			h.mu.Unlock()
			return
		}
		h.mu.Unlock()

		h.broadcastRegistration("registration_waiting")
	}
}

func (h *ScannerHandler) cancelRegistration(event string) {
	h.isRegistering = false
	h.registrationRFID = ""

	msg := map[string]interface{}{
		"type": event,
	}

	b, _ := json.Marshal(msg)
	h.hub.Broadcast <- b
}

func (h *ScannerHandler) HandleWSMessage(msg []byte) {
	var data struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(msg, &data); err != nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	if !h.isRegistering {
		return
	}

	switch data.Type {

	case "registration_heartbeat":
		h.lastHeartbeatAt = time.Now()

	case "registration_confirm":
		h.cancelRegistration("registration_confirmed")

	case "registration_cancel":
		h.cancelRegistration("registration_cancelled")

	case "search_ack":
		if h.searchCancel != nil {
			h.searchCancel()
			h.searchCancel = nil
		}

	}

}

func (h *ScannerHandler) Search(c *gin.Context) {
	var req struct {
		RFID       string `json:"rfid" binding:"required"`
		DeviceName string `json:"deviceName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.GetByName(c, req.DeviceName); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid device"})
		return
	}

	if h.searchCancel != nil {
		h.searchCancel()
		h.searchCancel = nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	h.searchCancel = cancel
	h.searchRFID = req.RFID

	go h.searchBroadcaster(ctx, req.RFID)

	c.JSON(http.StatusOK, gin.H{"message": "Search started"})

}

func (h *ScannerHandler) searchBroadcaster(ctx context.Context, rfid string) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeout := time.After(30 * time.Second)

	for {
		select {
		case <-ctx.Done():
			// 🔴 cancelled by new search or ACK
			return

		case <-timeout:
			msg := map[string]interface{}{
				"type": "search_timeout",
			}
			b, _ := json.Marshal(msg)
			h.hub.Broadcast <- b
			return

		case <-ticker.C:
			msg := map[string]interface{}{
				"type": "rfid_search",
				"payload": gin.H{
					"rfid": rfid,
				},
			}
			b, _ := json.Marshal(msg)
			h.hub.Broadcast <- b
		}
	}
}

func (h *ScannerHandler) Scan(c *gin.Context) {
	var req struct {
		RFID       string `json:"rfid" binding:"required"`
		DeviceName string `json:"deviceName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1️⃣ Validate device
	if _, err := h.service.GetByName(c, req.DeviceName); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid device"})
		return
	}

	// 2️⃣ Try drawer scan logic
	progress, hwResp :=
		h.drawerScanService.TryHandleRFID(c, req.RFID)

	// 3️⃣ Broadcast progress to frontend ONLY if scan succeeded
	if progress != nil {
		msg := map[string]interface{}{
			"type": "drawer_check_progress",
			"payload": gin.H{
				"counted":  progress.Counted,
				"expected": progress.Expected,
			},
		}

		b, _ := json.Marshal(msg)
		h.hub.Broadcast <- b
	}

	// 4️⃣ Respond to hardware with meaningful message
	c.JSON(http.StatusOK, hwResp)
}

func (h *ScannerHandler) ScannerHeartbeat(c *gin.Context) {
	var req struct {
		DeviceName string `json:"deviceName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate device exists
	if _, err := h.service.GetByName(c, req.DeviceName); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid device"})
		return
	}

	// 🔴 record heartbeat
	h.heartbeatStore.Beat(req.DeviceName)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
