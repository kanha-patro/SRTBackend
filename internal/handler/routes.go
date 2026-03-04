package handler

import "github.com/gin-gonic/gin"

// SetupRoutes registers top-level HTTP routes. Handlers are created with nil
// dependencies for now; real wiring happens in application bootstrap when
// concrete services/repositories are available.
func SetupRoutes(router *gin.Engine) {
	// Health
	router.GET("/health", HealthCheckHandler)

	// Trip routes (handlers created with nils - safe until wired properly)
	tripHandler := NewTripHandler(nil, nil)
	trips := router.Group("/trips")
	{
		trips.POST("/start", tripHandler.StartTrip)
		trips.POST("/location", tripHandler.UpdateLocation)
		trips.POST("/end", tripHandler.EndTrip)
		trips.GET("/active", tripHandler.GetActiveTrips)
	}

	// Route endpoints
	routeHandler := NewRouteHandler(nil)
	routes := router.Group("/routes")
	{
		routes.GET("", routeHandler.GetRoutes)
		routes.GET(":id", routeHandler.GetRoute)
	}

	// Location endpoints
	locationHandler := NewLocationHandler(nil)
	loc := router.Group("/locations")
	{
		loc.POST("/update", locationHandler.UpdateLocation)
		loc.GET("/active", locationHandler.GetActiveLocations)
	}

	// User endpoints
	userHandler := NewUserHandler(nil, nil)
	router.GET("/shuttles", userHandler.GetActiveShuttles)
}
