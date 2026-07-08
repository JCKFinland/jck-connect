package order

import (
	"github.com/gin-gonic/gin"

	orderhandler "github.com/JCKFinland/jck-connect/backend/internal/domain/order/handler"
)

// RegisterRoutes registers the Order HTTP endpoints.
func RegisterRoutes(
	router gin.IRoutes,
	handler *orderhandler.Handler,
) {
	router.GET(
		"/orders",
		handler.ListOrders,
	)

	router.GET(
		"/orders/:id",
		handler.GetOrder,
	)
}