package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Ambil role dari context
		roleValue, exists := c.Get("user_role")

		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Access forbidden",
			})
			return
		}

		// Pastikan role bertipe string
		userRole, ok := roleValue.(string)

		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Invalid user role",
			})
			return
		}

		// Periksa apakah role sesuai
		if userRole != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "You don't have permission to access this resource",
			})
			return
		}

		// Jika role sesuai, lanjutkan ke handler
		c.Next()
	}
}
