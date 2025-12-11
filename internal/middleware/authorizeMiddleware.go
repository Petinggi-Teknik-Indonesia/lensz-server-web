package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lensz-server-web/internal/model"
)

func RequireRole(db *gorm.DB, requiredRoleName string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. Extract role_id from JWT
		roleIDValue, exists := c.Get("role_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role ID missing in token"})
			c.Abort()
			return
		}

		// JWT numeric fields decode as float64
		roleIDFloat, ok := roleIDValue.(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role ID format"})
			c.Abort()
			return
		}
		tokenRoleID := uint(roleIDFloat)

		// 2. Convert required role name → database role ID
		var role model.Role
		if err := db.First(&role, "name = ?", requiredRoleName).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Required role does not exist"})
			c.Abort()
			return
		}

		requiredRoleID := role.ID

		// 3. Check authorization
		if tokenRoleID != requiredRoleID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
