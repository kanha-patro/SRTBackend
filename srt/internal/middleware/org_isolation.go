package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/akpatri/srt/internal/repository"
)

// OrgIsolationMiddleware ensures that all requests are scoped to the organization
func OrgIsolationMiddleware(orgRepo repository.OrgRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		orgID := c.Param("orgID") // Assuming orgID is passed as a URL parameter

		// Check if the organization exists
		if !orgRepo.Exists(orgID) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
			c.Abort()
			return
		}

		// Set the organization ID in the context for further processing
		c.Set("orgID", orgID)
		c.Next()
	}
}