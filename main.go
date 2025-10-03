package main

import (
	"log"
	"lunch-menu-api/internal/database"
	"lunch-menu-api/internal/handlers"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	if err := database.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDatabase()

	// Create Gin router
	r := gin.Default()

	// Setup CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

	// Setup routes
	handlers.SetupRoutes(r)

	// Get port from environment or default to 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Starting Restaurant Management API on port %s", port)
	log.Printf("API endpoints:")
	log.Printf("  GET    /api")
	log.Printf("  GET    /api/restaurants")
	log.Printf("  GET    /api/restaurants/:id")
	log.Printf("  GET    /api/restaurants/:id/menu")
	log.Printf("  GET    /api/menu-items/:id")

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
