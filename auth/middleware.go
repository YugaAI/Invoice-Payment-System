package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(paseto *PasetoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
			return
		}
		pToken, err := paseto.VerifyToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		userID, err := pToken.GetString("user_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid subject"})
			return
		}

		role, err := pToken.GetString("role")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
			return
		}

		c.Set("user_id", userID)
		c.Set("role", role)
		c.Next()
	}
}

func RoleGuard(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")
		if userRole != role {
			c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
