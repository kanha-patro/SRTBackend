package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/akpatri/srt/internal/observability"
	"github.com/akpatri/srt/pkg/config"
	"github.com/akpatri/srt/pkg/database"
	"github.com/akpatri/srt/internal/middleware"
	"github.com/akpatri/srt/internal/handler"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	// Set up observability
	observability.Setup(cfg.Observability)

	// Initialize Gin router
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.Logging())
	router.Use(middleware.ErrorHandler())

	// Set up routes
	handler.SetupRoutes(router, db)

	// Start the server
	if err := router.Run(cfg.Server.Address); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}