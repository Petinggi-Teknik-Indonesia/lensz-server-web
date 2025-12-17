package handler

import (
	"lensz-server-web/internal/model"
	"lensz-server-web/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// POST /api/users/register
func (h *UserHandler) Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.RoleID = 1

	if err := h.service.Register(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registration submitted, awaiting admin verification"})
}

// POST /api/users/admin-register
func (h *UserHandler) AdminRegister(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AdminRegister(c, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "admin registered successfully"})
}

// POST /api/users/login
func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(c, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// PATCH /api/users/verify/:email
func (h *UserHandler) VerifyUser(c *gin.Context) {
	email := c.Param("email")

	user, err := h.service.VerifyUser(c, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user verified successfully",
		"user":    user,
	})
}

// DELETE /api/users/reject/:email
func (h *UserHandler) CancelUser(c *gin.Context) {
	email := c.Param("email")

	if err := h.service.CancelUser(c, email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user rejected and deleted"})
}

// GET /api/users/unverified
func (h *UserHandler) GetAllUnverified(c *gin.Context) {
	users, err := h.service.GetAllUnverified(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET /api/users/verified
func (h *UserHandler) GetAllVerified(c *gin.Context) {
	users, err := h.service.GetAllVerified(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET /api/users/unverified/org/:id
func (h *UserHandler) GetUnverifiedByOrg(c *gin.Context) {
	orgID, _ := strconv.Atoi(c.Param("id"))
	users, err := h.service.GetUnverifiedByOrg(c, uint(orgID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET /api/users/verified/org/:id
func (h *UserHandler) GetVerifiedByOrg(c *gin.Context) {
	orgID, _ := strconv.Atoi(c.Param("id"))
	users, err := h.service.GetVerifiedByOrg(c, uint(orgID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
