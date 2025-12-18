package handler

import (
	"net/http"
	"strconv"

	"lensz-server-web/internal/model"
	"lensz-server-web/internal/service"

	"github.com/gin-gonic/gin"
)

type ScannerCRUDHandler struct {
	service *service.ScannerService
}

func NewScannerCRUDHandler(service *service.ScannerService) *ScannerCRUDHandler {
	return &ScannerCRUDHandler{service: service}
}

// POST /api/scanners
func (h *ScannerCRUDHandler) Create(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scanner := model.Scanner{
		Name: req.Name,
	}

	if err := h.service.Create(c, &scanner); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, scanner)
}

// GET /api/scanners
func (h *ScannerCRUDHandler) GetAll(c *gin.Context) {
	scanners, err := h.service.GetAll(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, scanners)
}

// GET /api/scanners/:id
func (h *ScannerCRUDHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	scanner, err := h.service.GetByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, scanner)
}

// DELETE /api/scanners/:id
func (h *ScannerCRUDHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.Delete(c, uint(id)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
