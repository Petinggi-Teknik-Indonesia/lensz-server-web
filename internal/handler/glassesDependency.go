package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/service"
)

type GlassesDependencyHandler struct {
	service *service.GlassesDependencyService
}

func NewGlassesDependencyHandler(service *service.GlassesDependencyService) *GlassesDependencyHandler {
	return &GlassesDependencyHandler{service: service}
}

// -------------------- DRAWER --------------------
func (h *GlassesDependencyHandler) CreateDrawer(c *gin.Context) {
	var req model.Drawer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateDrawer(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (h *GlassesDependencyHandler) GetDrawerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	drawer, err := h.service.GetDrawerByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drawer)
}

func (h *GlassesDependencyHandler) GetAllDrawers(c *gin.Context) {
	drawers, err := h.service.GetAllDrawers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drawers)
}

func (h *GlassesDependencyHandler) UpdateDrawer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req model.Drawer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = uint(id)

	if err := h.service.UpdateDrawer(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *GlassesDependencyHandler) DeleteDrawer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteDrawer(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// -------------------- BRAND --------------------
func (h *GlassesDependencyHandler) CreateBrand(c *gin.Context) {
	var req model.Brand
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateBrand(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (h *GlassesDependencyHandler) GetBrandByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	brand, err := h.service.GetBrandByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brand)
}

func (h *GlassesDependencyHandler) GetAllBrands(c *gin.Context) {
	brands, err := h.service.GetAllBrands(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func (h *GlassesDependencyHandler) UpdateBrand(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req model.Brand
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = uint(id)

	if err := h.service.UpdateBrand(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *GlassesDependencyHandler) DeleteBrand(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteBrand(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// -------------------- COMPANY --------------------
func (h *GlassesDependencyHandler) CreateCompany(c *gin.Context) {
	var req model.Company
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateCompany(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (h *GlassesDependencyHandler) GetCompanyByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	company, err := h.service.GetCompanyByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *GlassesDependencyHandler) GetAllCompanies(c *gin.Context) {
	companies, err := h.service.GetAllCompanies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (h *GlassesDependencyHandler) UpdateCompany(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req model.Company
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ID = uint(id)

	if err := h.service.UpdateCompany(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *GlassesDependencyHandler) DeleteCompany(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteCompany(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
