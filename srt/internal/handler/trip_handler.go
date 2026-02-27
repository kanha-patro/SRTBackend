package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/akpatri/srt/internal/service"
	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/observability"
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
	var request domain.StartTripRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.logger.Error("Failed to bind JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	trip, err := h.tripService.StartTrip(request)
	if err != nil {
		h.logger.Error("Failed to start trip", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start trip"})
		return
	}

	c.JSON(http.StatusOK, trip)
}

// UpdateLocation handles the request to update the driver's location
func (h *TripHandler) UpdateLocation(c *gin.Context) {
	var request domain.UpdateLocationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.logger.Error("Failed to bind JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.tripService.UpdateLocation(request)
	if err != nil {
		h.logger.Error("Failed to update location", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update location"})
		return
	}

	c.Status(http.StatusNoContent)
}

// EndTrip handles the request to end a trip
func (h *TripHandler) EndTrip(c *gin.Context) {
	var request domain.EndTripRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.logger.Error("Failed to bind JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.tripService.EndTrip(request)
	if err != nil {
		h.logger.Error("Failed to end trip", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to end trip"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetActiveTrips handles the request to get all active trips
func (h *TripHandler) GetActiveTrips(c *gin.Context) {
	trips, err := h.tripService.GetActiveTrips()
	if err != nil {
		h.logger.Error("Failed to get active trips", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get active trips"})
		return
	}

	c.JSON(http.StatusOK, trips)
}