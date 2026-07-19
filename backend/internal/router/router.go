package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JCKFinland/jck-connect/backend/internal/container"
	"github.com/JCKFinland/jck-connect/backend/internal/middleware"

	auth "github.com/JCKFinland/jck-connect/backend/internal/domain/auth"
	order "github.com/JCKFinland/jck-connect/backend/internal/domain/order"
	product "github.com/JCKFinland/jck-connect/backend/internal/domain/product"
	user "github.com/JCKFinland/jck-connect/backend/internal/domain/user"
	wallet "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet"
	purchase "github.com/JCKFinland/jck-connect/backend/internal/handler/purchase"
)

// Register registers all application routes.
func Register(
	engine *gin.Engine,
	c *container.Container,
) {

	//--------------------------------------------------
	// Health Check
	//--------------------------------------------------

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	//--------------------------------------------------
	// API v1
	//--------------------------------------------------

	api := engine.Group("/api/v1")

	//--------------------------------------------------
	// Public Routes
	//--------------------------------------------------

	auth.RegisterRoutes(
		api,
		c.AuthHandler,
	)

	//--------------------------------------------------
	// Protected Routes
	//--------------------------------------------------

	protected := api.Group("")
	protected.Use(
		middleware.Auth(c.JWTManager),
	)

	user.RegisterRoutes(
		protected,
		c.UserHandler,
	)

	wallet.RegisterRoutes(
		protected,
		c.WalletHandler,
	)

	order.RegisterRoutes(
		protected,
		c.OrderHandler,
	)

	product.RegisterRoutes(
		protected,
		c.ProductHandler,
	)

	purchase.RegisterRoutes(
		protected,
		c.PurchaseHandler,
	)
}
