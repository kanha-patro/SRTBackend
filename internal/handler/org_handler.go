package handler

import (
	"net/http"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/service"
	"github.com/gin-gonic/gin"
)

// OrgHandler handles organization-related HTTP requests
type OrgHandler struct {
	orgService service.OrgService
}

// NewOrgHandler creates a new OrgHandler
func NewOrgHandler(orgService service.OrgService) *OrgHandler {
	return &OrgHandler{orgService: orgService}
}

// RegisterOrg handles the registration of a new organization
func (h *OrgHandler) RegisterOrg(c *gin.Context) {
	var org domain.Org
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.orgService.RegisterOrg(&org); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, org)
}

// ApproveOrg handles the approval of an organization by an admin
func (h *OrgHandler) ApproveOrg(c *gin.Context) {
	orgID := c.Param("id")

	if err := h.orgService.ApproveOrg(orgID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization approved"})
}

// SuspendOrg handles the suspension of an organization
func (h *OrgHandler) SuspendOrg(c *gin.Context) {
	orgID := c.Param("id")

	if err := h.orgService.SuspendOrg(orgID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization suspended"})
}

// GetActiveOrgs retrieves all active organizations
func (h *OrgHandler) GetActiveOrgs(c *gin.Context) {
	orgs, err := h.orgService.GetActiveOrgs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orgs)
}

// UpdateOrg handles the dynamic update of an organization's details
func (h *OrgHandler) UpdateOrg(c *gin.Context) {
	var org domain.Org
	if err := c.ShouldBindJSON(&org); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	orgID := c.Param("id")
	if err := h.orgService.UpdateOrg(orgID, &org); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, org)
}
