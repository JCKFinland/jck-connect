package product

import (
	"github.com/gin-gonic/gin"

	producthandler "github.com/JCKFinland/jck-connect/backend/internal/domain/product/handler"
)

// RegisterRoutes registers Product routes.
func RegisterRoutes(
	router gin.IRoutes,
	handler *producthandler.Handler,
) {

	router.GET(
		"/products",
		handler.List,
	)

	router.GET(
		"/products/:id",
		handler.GetByID,
	)
}