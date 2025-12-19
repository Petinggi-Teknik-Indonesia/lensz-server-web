package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"lensz-server-web/internal/model"
	"lensz-server-web/internal/service"
	"lensz-server-web/internal/ws"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	service        *service.HistoryService
	glassesService *service.GlassesService
	scannerService *service.ScannerService
	hub            *ws.Hub
}

func NewHistoryHandler(historyService *service.HistoryService, glassesService *service.GlassesService, hub *ws.Hub) *HistoryHandler {
	return &HistoryHandler{
		service:        historyService,
		glassesService: glassesService,
		hub:            hub,
	}
}

// GET /api/glasses/:id/history
func (h *HistoryHandler) GetByGlassesID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid glasses ID"})
		return
	}

	histories, err := h.service.GetHistoryByGlassesID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, histories)
}

// PATCH /api/glasses/status
func (h *HistoryHandler) UpdateStatusByRFID(c *gin.Context) {
	var req struct {
		RFID       string              `json:"rfid" binding:"required"`
		Status     model.GlassesStatus `json:"status" binding:"required"`
		DeviceName string              `json:"deviceName" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.scannerService.GetByName(c, req.DeviceName); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid device"})
		return
	}

	if err := h.glassesService.UpdateGlassesStatusByRFID(c, req.RFID, req.Status, 0); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ✅ Broadcast WebSocket update
	msg := map[string]interface{}{
		"type": "glasses_status_updated",
		"payload": map[string]interface{}{
			"rfid":   req.RFID,
			"status": req.Status,
		},
	}
	b, _ := json.Marshal(msg)
	h.hub.Broadcast <- b

	c.JSON(http.StatusOK, gin.H{"message": "status updated successfully"})
}
