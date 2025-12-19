package handler

import (
	"encoding/json"
	"strconv"

	"lensz-server-web/internal/ws"

	"lensz-server-web/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type DrawerScanHandler struct {
	service *service.DrawerScanService
	hub     *ws.Hub
}

func NewDrawerScanHandler(
	service *service.DrawerScanService,
	hub *ws.Hub,
) *DrawerScanHandler {
	return &DrawerScanHandler{service: service, hub: hub}
}

// POST /api/drawers/:id/check
func (h *DrawerScanHandler) Start(c *gin.Context) {
	drawerID, _ := strconv.Atoi(c.Param("id"))
	sessionID := uuid.NewString()

	session, err := h.service.StartSession(c, uint(drawerID), sessionID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"sessionId": session.ID,
		"expected":  session.ExpectedTotal,
	})
}

// POST /api/drawers/check/:sessionId/stop
func (h *DrawerScanHandler) Stop(c *gin.Context) {
	sessionID := c.Param("sessionId")

	result, err := h.service.StopSession(c, sessionID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	msg := map[string]interface{}{
		"type": "drawer_check_result",
		"payload": gin.H{
			"scanned":   result.Scanned,
			"expected":  result.Expected,
			"missing":   result.Missing,
			"mislabels": result.Mislabels,
		},
	}

	b, _ := json.Marshal(msg)
	h.hub.Broadcast <- b

	c.JSON(200, gin.H{"status": "completed"})
}
