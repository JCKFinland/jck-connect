package router

import (
	"github.com/gin-gonic/gin"

	"github.com/JCKFinland/jck-connect/backend/internal/middleware"
	"github.com/JCKFinland/jck-connect/backend/internal/shared"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		shared.Success(c, "service is healthy", nil)
	})

	// API v1 group
	v1 := r.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		auth.GET("/ping", func(c *gin.Context) {
			shared.Success(c, "auth module ready", nil)
		})
	}

	return r
}
