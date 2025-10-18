package handler

import (
	"net/http"
	"strconv"

	"lensz-server-web/internal/model"
	"lensz-server-web/internal/service"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	service *service.HistoryService
	glassesService *service.GlassesService
}

func NewHistoryHandler(historyService *service.HistoryService, glassesService *service.GlassesService) *HistoryHandler {
	return &HistoryHandler{
		service: historyService,
		glassesService: glassesService,
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
		RFID   string              `json:"rfid" binding:"required"`
		Status model.GlassesStatus `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.glassesService.UpdateGlassesStatusByRFID(c, req.RFID, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status updated successfully"})
}
