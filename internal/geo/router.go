package geo

import (
	"net/http"

	"github.com/akpatri/srt/internal/service"
	"github.com/gin-gonic/gin"
)

// Router handles routing for geo-related endpoints
type Router struct {
	geoService service.GeoService
}

// NewRouter creates a new Router instance
func NewRouter(geoService service.GeoService) *Router {
	return &Router{geoService: geoService}
}

// RegisterRoutes registers the geo-related routes
func (r *Router) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/snap", r.SnapLocation)
	router.GET("/search", r.SearchNearbyStops)
}

// SnapLocation handles the snapping of locations to the nearest stop
func (r *Router) SnapLocation(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	if latitude == "" || longitude == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "latitude and longitude are required"})
		return
	}

	snapResult, err := r.geoService.SnapLocation(latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, snapResult)
}

// SearchNearbyStops handles searching for nearby stops
func (r *Router) SearchNearbyStops(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")

	if latitude == "" || longitude == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "latitude and longitude are required"})
		return
	}

	stops, err := r.geoService.SearchNearbyStops(latitude, longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stops)
}
