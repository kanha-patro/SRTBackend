package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/akpatri/srt/internal/service"
	"github.com/akpatri/srt/internal/observability"
)

type UserHandler struct {
	userService service.UserService
	logger      observability.Logger
}

func NewUserHandler(userService service.UserService, logger observability.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

// GetActiveShuttles retrieves all active shuttles for a given organization and optional filters.
func (h *UserHandler) GetActiveShuttles(c *gin.Context) {
	orgCode := c.Query("org_code")
	routeCode := c.Query("route_code")
	nearbyLocation := c.Query("nearby_location")

	activeShuttles, err := h.userService.GetActiveShuttles(orgCode, routeCode, nearbyLocation)
	if err != nil {
		h.logger.Error("Failed to retrieve active shuttles", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve active shuttles"})
		return
	}

	c.JSON(http.StatusOK, activeShuttles)
}