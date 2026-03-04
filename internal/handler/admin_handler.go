package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/akpatri/srt/internal/service"
	"github.com/akpatri/srt/internal/observability"
)

type AdminHandler struct {
	orgService service.OrgService
	logger     observability.Logger
}

func NewAdminHandler(orgService service.OrgService, logger observability.Logger) *AdminHandler {
	return &AdminHandler{
		orgService: orgService,
		logger:     logger,
	}
}

// ApproveOrg handles the approval of organization registration
func (h *AdminHandler) ApproveOrg(c *gin.Context) {
	orgID := c.Param("org_id")
	if err := h.orgService.ApproveOrg(orgID); err != nil {
		h.logger.Error("Failed to approve organization", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve organization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Organization approved"})
}

// SuspendOrg handles the suspension of an organization
func (h *AdminHandler) SuspendOrg(c *gin.Context) {
	orgID := c.Param("org_id")
	if err := h.orgService.SuspendOrg(orgID); err != nil {
		h.logger.Error("Failed to suspend organization", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to suspend organization"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Organization suspended"})
}

// MonitorActiveTrips handles the monitoring of active trips
func (h *AdminHandler) MonitorActiveTrips(c *gin.Context) {
	trips, err := h.orgService.GetActiveTrips()
	if err != nil {
		h.logger.Error("Failed to retrieve active trips", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve active trips"})
		return
	}
	c.JSON(http.StatusOK, trips)
}

// ForceStopTrip handles the force stopping of a trip
func (h *AdminHandler) ForceStopTrip(c *gin.Context) {
	tripID := c.Param("trip_id")
	if err := h.orgService.ForceStopTrip(tripID); err != nil {
		h.logger.Error("Failed to force stop trip", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to force stop trip"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Trip force stopped"})
}

// RevokeOTPSession handles the revocation of an OTP session
func (h *AdminHandler) RevokeOTPSession(c *gin.Context) {
	sessionID := c.Param("session_id")
	if err := h.orgService.RevokeOTPSession(sessionID); err != nil {
		h.logger.Error("Failed to revoke OTP session", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke OTP session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP session revoked"})
}