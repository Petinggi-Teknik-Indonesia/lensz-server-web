package handler

import (
	"net/http"
	"strconv"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/service"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	service *service.OrganizationService
}

func NewOrganizationHandler(service *service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{service: service}
}

func (h *OrganizationHandler) Create(c *gin.Context) {
	var req model.Organization
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateOrganization(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (h *OrganizationHandler) GetAll(c *gin.Context) {
	orgs, err := h.service.GetAllOrganizations(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orgs)
}

func (h *OrganizationHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	org, err := h.service.GetOrganizationByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, org)
}

func (h *OrganizationHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req model.Organization
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = uint(id)
	if err := h.service.UpdateOrganization(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *OrganizationHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteOrganization(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
