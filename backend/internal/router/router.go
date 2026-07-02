package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	// Development mode for v0.1.0
	gin.SetMode(gin.DebugMode)

	router := gin.New()

	// Prevent server crashes from panics
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

	// API Version 1
	v1 := router.Group("/api/v1")
	{
		// Future routes:
		// /auth
		// /users
		// /transactions
		// /payments
		// /vtpass
	}

	return router
}