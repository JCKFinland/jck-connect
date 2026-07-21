package purchase

import "github.com/gin-gonic/gin"

// RegisterRoutes registers purchase routes.
func RegisterRoutes(
	router gin.IRoutes,
	handler *Handler,
) {
	router.POST(
		"/purchases",
		handler.Purchase,
	)
}
