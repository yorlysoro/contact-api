// internal/auth/middleware.go
package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Expected format: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization format must be Bearer {token}"})
			c.Abort()
			return
		}

		claims, err := ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Store userID in context for use in handlers
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func SetupRoutes(router *gin.Engine, handler *Handler, authMiddleware gin.HandlerFunc) {
	api := router.Group("/api/v1")
	{
		// Protected routes
		contacts := api.Group("/contacts")
		contacts.Use(authMiddleware)
		{
			contacts.POST("/", handler.Create)
			contacts.GET("/:id", handler.GetByID)
			// Add UPDATE and DELETE here
		}
	}
}
