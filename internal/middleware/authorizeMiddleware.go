package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRoles(allowedRoleIDs ...uint) gin.HandlerFunc {
	// Build lookup set once (NOT per request)
	allowed := make(map[uint]struct{}, len(allowedRoleIDs))
	for _, id := range allowedRoleIDs {
		allowed[id] = struct{}{}
	}

	return func(c *gin.Context) {
		// Extract role_id from JWT context
		roleIDValue, exists := c.Get("role_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "role_id missing in token",
			})
			c.Abort()
			return
		}

		roleIDFloat, ok := roleIDValue.(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid role_id format",
			})
			c.Abort()
			return
		}

		tokenRoleID := uint(roleIDFloat)

		// OR authorization: role must be in allowed list
		if _, ok := allowed[tokenRoleID]; !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "forbidden: insufficient permissions",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
