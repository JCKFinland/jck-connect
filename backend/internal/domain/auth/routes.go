package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/JCKFinland/jck-connect/backend/internal/domain/auth/handler"
)

// RegisterRoutes registers authentication routes.
func RegisterRoutes(
	router *gin.RouterGroup,
	handler *handler.Handler,
) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", handler.Login)
	}
}
