package middleware

import (
	"finpro/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ResponseJSON(c, http.StatusUnauthorized, false, "Authorization header is required", nil)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ResponseJSON(c, http.StatusUnauthorized, false, "Invalid authorization format", nil)
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			utils.ResponseJSON(c, http.StatusUnauthorized, false, "Invalid or expired token", nil)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			utils.ResponseJSON(c, http.StatusForbidden, false, "Admin access required", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
