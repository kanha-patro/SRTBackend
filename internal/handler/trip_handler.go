package handler

import (
	"net/http"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/observability"
	"github.com/akpatri/srt/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TripHandler struct {
	tripService service.TripService
	logger      observability.Logger
}

func NewTripHandler(tripService service.TripService, logger observability.Logger) *TripHandler {
	return &TripHandler{
		tripService: tripService,
		logger:      logger,
	}
}

// StartTrip handles the request to start a trip
func (h *TripHandler) StartTrip(c *gin.Context) {
	var trip domain.Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.tripService.StartTrip(c.Request.Context(), &trip); err != nil {
		h.logger.Error("Failed to start trip", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start trip"})
		return
	}

	c.JSON(http.StatusOK, trip)
}

// UpdateLocation handles the request to update the driver's location
func (h *TripHandler) UpdateLocation(c *gin.Context) {
	// Expect payload: { "trip_id": "...", "location": { ... } }
	var payload struct {
		TripID   string          `json:"trip_id"`
		Location domain.Location `json:"location"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.tripService.UpdateLocation(c.Request.Context(), payload.TripID, &payload.Location); err != nil {
		h.logger.Error("Failed to update location", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update location"})
		return
	}

	c.Status(http.StatusNoContent)
}

// EndTrip handles the request to end a trip
func (h *TripHandler) EndTrip(c *gin.Context) {
	var payload struct {
		TripID string `json:"trip_id"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.tripService.EndTrip(c.Request.Context(), payload.TripID); err != nil {
		h.logger.Error("Failed to end trip", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to end trip"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetActiveTrips handles the request to get all active trips
func (h *TripHandler) GetActiveTrips(c *gin.Context) {
	trips, err := h.tripService.GetActiveTrips(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get active trips", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get active trips"})
		return
	}

	c.JSON(http.StatusOK, trips)
}
