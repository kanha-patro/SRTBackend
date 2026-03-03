package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorHandler is a middleware that handles errors and sends appropriate responses.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Log the error using Zap
			for _, err := range c.Errors {
				zap.L().Error("Error occurred", zap.Error(err))
			}

			// Send a generic error response
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
		}
	}
}