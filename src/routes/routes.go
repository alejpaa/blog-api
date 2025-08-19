package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all the routes for the application.
func SetupRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
}
