package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RBACMiddleware is a middleware for role-based access control.
func RBACMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoles, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "roles not found"})
			c.Abort()
			return
		}

		roles := userRoles.([]string)
		if !hasRequiredRole(roles, requiredRoles) {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// hasRequiredRole checks if the user has at least one of the required roles.
func hasRequiredRole(userRoles, requiredRoles []string) bool {
	for _, userRole := range userRoles {
		for _, requiredRole := range requiredRoles {
			if strings.EqualFold(userRole, requiredRole) {
				return true
			}
		}
	}
	return false
}