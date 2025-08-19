package main

import (
	"blog-api/src/config"
	"blog-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
