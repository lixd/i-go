package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.DefaultQuery("auth", "false")
		if auth != "true" {
			c.JSON(http.StatusForbidden, gin.H{"status": "403", "message": "auth false"})
		}
	}
}
