package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// New creates and configures the Gin router.
func New(appEnv string) *gin.Engine {
	// Configure Gin mode
	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	// Global middleware
	router.Use(gin.Recovery())

	// Root endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "jck-connect",
			"version":     "v0.1.0",
			"status":      "running",
		})
	})

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "jck-connect",
		})
	})

	// API v1
	api := router.Group("/api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "API v1 is running",
			})
		})
	}

	return router
}