package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler handles health check requests.
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}