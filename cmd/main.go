package main

import (
	"log"
	"task_flow/internal/apps/users/routes"
	"task_flow/internal/config"
	"task_flow/pkg/database"
	"task_flow/pkg/logger"
	"task_flow/tests/health"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger.InitLogger(cfg.Env)

	// Initialize database
	db, err := database.InitDB(cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	// Register health routes
	api := router.Group("/api/v1")
	{
		health.SetupHealthRoutes(api, db)
		routes.RegisterUserRoutes(api, db)
	}

	// Start server
	log.Printf("Server is running on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
