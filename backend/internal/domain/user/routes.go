package user

import (
	"github.com/gin-gonic/gin"

	"github.com/JCKFinland/jck-connect/backend/internal/domain/user/handler"
)

// RegisterRoutes registers all user routes.
func RegisterRoutes(
	router *gin.RouterGroup,
	h *handler.Handler,
) {

	users := router.Group("/users")

	{
		users.GET("/me", h.Me)

		users.PUT("/me", h.UpdateMe)
	}
}