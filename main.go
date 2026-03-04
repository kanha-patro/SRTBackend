package main

import (
	"log"

	"github.com/akpatri/srt/internal/handler"
	"github.com/akpatri/srt/internal/middleware"
	"github.com/akpatri/srt/internal/observability"
	"github.com/akpatri/srt/pkg/config"
	"github.com/akpatri/srt/pkg/database"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize database
	database.Connect()
	defer database.Close()

	// Set up observability: create logger and initialize metrics (best-effort)
	logger, err := observability.NewLogger()
	if err != nil {
		log.Printf("warning: could not initialize logger: %v", err)
	} else {
		defer logger.Sync()
	}

	if err := observability.InitMetrics(); err != nil {
		log.Printf("warning: could not initialize metrics: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Apply middleware
	// Use a simple Zap logger instance for request logging middleware.
	router.Use(middleware.LoggingMiddleware(zap.NewExample()))
	router.Use(middleware.ErrorHandler())

	// Set up routes
	handler.SetupRoutes(router)

	// Start the server
	if err := router.Run(cfg.Port); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
