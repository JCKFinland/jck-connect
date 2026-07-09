package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JCKFinland/jck-connect/backend/internal/config"

	auth "github.com/JCKFinland/jck-connect/backend/internal/domain/auth"
	authhandler "github.com/JCKFinland/jck-connect/backend/internal/domain/auth/handler"
	authservice "github.com/JCKFinland/jck-connect/backend/internal/domain/auth/service"

	order "github.com/JCKFinland/jck-connect/backend/internal/domain/order"
	orderhandler "github.com/JCKFinland/jck-connect/backend/internal/domain/order/handler"
	orderpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/order/repository/postgres"
	orderservice "github.com/JCKFinland/jck-connect/backend/internal/domain/order/service"

	user "github.com/JCKFinland/jck-connect/backend/internal/domain/user"
	userhandler "github.com/JCKFinland/jck-connect/backend/internal/domain/user/handler"
	userpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/user/repository/postgres"
	userservice "github.com/JCKFinland/jck-connect/backend/internal/domain/user/service"

	wallet "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet"
	wallethandler "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/handler"
	walletpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository/postgres"
	walletservice "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/service"

	"github.com/JCKFinland/jck-connect/backend/internal/middleware"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
	jwtpkg "github.com/JCKFinland/jck-connect/backend/pkg/jwt"

	product "github.com/JCKFinland/jck-connect/backend/internal/domain/product"

	producthandler "github.com/JCKFinland/jck-connect/backend/internal/domain/product/handler"

	productpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/product/repository/postgres"

	productservice "github.com/JCKFinland/jck-connect/backend/internal/domain/product/service"
)

// Register registers all application routes.
func Register(
	engine *gin.Engine,
	cfg *config.Config,
	db *database.Database,
) {

	//--------------------------------------------------
	// Infrastructure
	//--------------------------------------------------

	jwtManager := jwtpkg.New(
		cfg.JWTSecret,
		cfg.JWTAccessTokenDuration,
		cfg.JWTRefreshTokenDuration,
	)

	//--------------------------------------------------
	// User Domain
	//--------------------------------------------------

	userRepository := userpostgres.New(db)

	userService := userservice.New(
		userRepository,
	)

	userHandler := userhandler.New(
		userService,
	)

	//--------------------------------------------------
	// Wallet Domain
	//--------------------------------------------------

	walletRepository := walletpostgres.New(db)

	walletService := walletservice.New(
		walletRepository,
	)

	walletHandler := wallethandler.New(
		walletService,
	)

	//--------------------------------------------------
	// Product Domain
	//--------------------------------------------------

	productRepository := productpostgres.New(db)

	productService := productservice.New(
		productRepository,
	)

	productHandler := producthandler.New(
		productService,
	)

	//--------------------------------------------------
	// Order Domain
	//--------------------------------------------------

	orderRepository := orderpostgres.New(db)

	orderService := orderservice.New(
		orderRepository,
	)

	orderHandler := orderhandler.New(
		orderService,
	)

	//--------------------------------------------------
	// Auth Domain
	//--------------------------------------------------

	authService := authservice.New(
		userService,
		jwtManager,
	)

	authHandler := authhandler.New(
		authService,
	)

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
		authHandler,
	)

	//--------------------------------------------------
	// Protected Routes
	//--------------------------------------------------

	protected := api.Group("")
	protected.Use(
		middleware.Auth(jwtManager),
	)

	user.RegisterRoutes(
		protected,
		userHandler,
	)

	wallet.RegisterRoutes(
		protected,
		walletHandler,
	)

	order.RegisterRoutes(
		protected,
		orderHandler,
	)

	product.RegisterRoutes(
		protected,
		productHandler,
	)
}
