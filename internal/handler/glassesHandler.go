package handler

import (
	"net/http"
	"strconv"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/service"

	"github.com/gin-gonic/gin"
)

type GlassesHandler struct {
	service *service.GlassesService
}

func NewGlassesHandler(service *service.GlassesService) *GlassesHandler {
	return &GlassesHandler{service: service}
}

// POST /glasses
func (h *GlassesHandler) Create(c *gin.Context) {
	var req model.Glasses
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateGlasses(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

// GET /glasses/:id
func (h *GlassesHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	glasses, err := h.service.GetGlassesByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, glasses)
}

// GET /glasses
func (h *GlassesHandler) GetAll(c *gin.Context) {
	glasses, err := h.service.GetAllGlasses(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, glasses)
}

// PUT /glasses/:id
func (h *GlassesHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req model.Glasses
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = uint(id)

	if err := h.service.UpdateGlasses(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// DELETE /glasses/:id
func (h *GlassesHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteGlasses(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
