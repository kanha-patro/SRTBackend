package handler

import (
	"net/http"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/observability"
	"github.com/akpatri/srt/internal/service"
	"github.com/akpatri/srt/pkg/errors"
	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	locationService service.LocationService
}

func NewLocationHandler(locationService service.LocationService) *LocationHandler {
	return &LocationHandler{locationService: locationService}
}

// UpdateLocation handles the incoming location updates from drivers.
func (h *LocationHandler) UpdateLocation(c *gin.Context) {
	var location domain.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		observability.LogError(err)
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid location data"))
		return
	}

	ctx := c.Request.Context()
	tripID := location.TripID

	if err := h.locationService.UpdateLocation(ctx, tripID, location); err != nil {
		observability.LogError(err)
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError("Failed to update location"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location updated successfully"})
}

// GetActiveLocations retrieves all active locations for users to track.
func (h *LocationHandler) GetActiveLocations(c *gin.Context) {
	ctx := c.Request.Context()
	orgID := c.Query("org_id")

	activeLocations, err := h.locationService.GetActiveLocations(ctx, orgID)
	if err != nil {
		observability.LogError(err)
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError("Failed to retrieve active locations"))
		return
	}

	c.JSON(http.StatusOK, activeLocations)
}
