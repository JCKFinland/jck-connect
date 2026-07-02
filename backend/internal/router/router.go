package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Recovery middleware only for now
	r.Use(gin.Recovery())

	// Health endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "jck-connect",
		})
	})

	// API v1
	api := r.Group("/api/v1")
	{
		_ = api
	}

	return r
}