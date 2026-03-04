package handler

import (
	"net/http"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/service"
	"github.com/akpatri/srt/pkg/errors"
	"github.com/gin-gonic/gin"
)

// RouteHandler handles route-related HTTP requests
type RouteHandler struct {
	routeService service.RouteService
}

// NewRouteHandler creates a new RouteHandler
func NewRouteHandler(routeService service.RouteService) *RouteHandler {
	return &RouteHandler{routeService: routeService}
}

// CreateRoute handles the creation of a new route
func (h *RouteHandler) CreateRoute(c *gin.Context) {
	var route domain.Route
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid request payload"))
		return
	}

	if err := h.routeService.CreateRoute(&route); err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, route)
}

// UpdateRoute handles the updating of an existing route
func (h *RouteHandler) UpdateRoute(c *gin.Context) {
	var route domain.Route
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid request payload"))
		return
	}

	if err := h.routeService.UpdateRoute(&route); err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, route)
}

// DeleteRoute handles the deletion of a route
func (h *RouteHandler) DeleteRoute(c *gin.Context) {
	routeID := c.Param("id")

	if err := h.routeService.DeleteRoute(routeID); err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetRoutes handles fetching all routes
func (h *RouteHandler) GetRoutes(c *gin.Context) {
	orgID := c.Query("org_id")
	routes, err := h.routeService.GetAllRoutes(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, routes)
}

// GetRoute handles fetching a specific route by ID
func (h *RouteHandler) GetRoute(c *gin.Context) {
	routeID := c.Param("id")
	route, err := h.routeService.GetRoute(routeID)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.NewNotFoundError("Route not found"))
		return
	}

	c.JSON(http.StatusOK, route)
}
