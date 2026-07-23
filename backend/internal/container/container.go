package container

import (
	"github.com/JCKFinland/jck-connect/backend/internal/config"
	"github.com/JCKFinland/jck-connect/backend/pkg/database"
	"time"

	//--------------------------------------------------
	// Auth
	//--------------------------------------------------

	authhandler "github.com/JCKFinland/jck-connect/backend/internal/domain/auth/handler"
	authpi "github.com/JCKFinland/jck-connect/backend/internal/domain/auth/pi"
	authservice "github.com/JCKFinland/jck-connect/backend/internal/domain/auth/service"

	jwtpkg "github.com/JCKFinland/jck-connect/backend/pkg/jwt"

	//--------------------------------------------------
	// User
	//--------------------------------------------------

	userhandler "github.com/JCKFinland/jck-connect/backend/internal/domain/user/handler"
	userrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/user/repository"
	userpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/user/repository/postgres"
	userservice "github.com/JCKFinland/jck-connect/backend/internal/domain/user/service"

	//--------------------------------------------------
	// Wallet
	//--------------------------------------------------

	wallethandler "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/handler"
	walletrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository"
	walletpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository/postgres"
	walletservice "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/service"

	//--------------------------------------------------
	// Product
	//--------------------------------------------------

	producthandler "github.com/JCKFinland/jck-connect/backend/internal/domain/product/handler"
	productrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/product/repository"
	productpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/product/repository/postgres"

	productservice "github.com/JCKFinland/jck-connect/backend/internal/domain/product/service"

	//--------------------------------------------------
	// Order
	//--------------------------------------------------

	orderhandler "github.com/JCKFinland/jck-connect/backend/internal/domain/order/handler"
	orderrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/order/repository"
	orderpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/order/repository/postgres"

	orderservice "github.com/JCKFinland/jck-connect/backend/internal/domain/order/service"

	//--------------------------------------------------
	// Transaction
	//--------------------------------------------------

	transactionrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/repository"
	transactionpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/repository/postgres"
	transactionservice "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/service"

	//--------------------------------------------------
	// Purchase Use Case
	//--------------------------------------------------

	purchasehandler "github.com/JCKFinland/jck-connect/backend/internal/handler/purchase"
	purchase "github.com/JCKFinland/jck-connect/backend/internal/usecase/purchase"
)

type Container struct {
	Config *config.Config
	DB     *database.Database

	//--------------------------------------------------
	// User
	//--------------------------------------------------

	UserRepository userrepo.Repository
	UserService    userservice.Service
	UserHandler    *userhandler.Handler

	//--------------------------------------------------
	// Wallet
	//--------------------------------------------------

	WalletRepository walletrepo.Repository
	WalletService    walletservice.Service
	WalletHandler    *wallethandler.Handler

	//--------------------------------------------------
	// Product
	//--------------------------------------------------

	ProductRepository productrepo.Repository
	ProductService    productservice.Service
	ProductHandler    *producthandler.Handler

	//--------------------------------------------------
	// Order
	//--------------------------------------------------

	OrderRepository orderrepo.Repository
	OrderService    orderservice.Service
	OrderHandler    *orderhandler.Handler

	//--------------------------------------------------
	// Transaction
	//--------------------------------------------------

	TransactionRepository transactionrepo.Repository
	TransactionService    transactionservice.Service

	//--------------------------------------------------
	// Auth
	//--------------------------------------------------

	JWTManager  *jwtpkg.Manager
	AuthService authservice.Service
	AuthHandler *authhandler.Handler

	//--------------------------------------------------
	// Purchase
	//--------------------------------------------------

	PurchaseService purchase.Service
	PurchaseHandler *purchasehandler.Handler
}

func New(
	cfg *config.Config,
	db *database.Database,
) *Container {

	c := &Container{
		Config: cfg,
		DB:     db,
	}

	c.Compose()

	return c
}

// BuildTransaction wires the Transaction domain.
func (c *Container) BuildTransaction() {

	c.TransactionRepository = transactionpostgres.New(
		c.DB,
	)

	c.TransactionService = transactionservice.New(
		c.TransactionRepository,
	)
}

// BuildAuth wires the Auth domain.
func (c *Container) BuildAuth() {

	c.JWTManager = jwtpkg.New(
		c.Config.JWTSecret,
		c.Config.JWTAccessTokenDuration,
		c.Config.JWTRefreshTokenDuration,
	)

	//--------------------------------------------------
	// Pi Authentication
	//--------------------------------------------------

	piConfig := authpi.Config{
		BaseURL: "https://api.minepi.com",
		APIKey:  c.Config.PiAPIKey,
		Timeout: 10 * time.Second,
	}

	piClient := authpi.NewClient(piConfig)

	piVerifier := authpi.NewVerifier(piClient)

	//--------------------------------------------------
	// Authentication Service
	//--------------------------------------------------

	c.AuthService = authservice.New(
		c.UserService,
		c.JWTManager,
		piVerifier,
	)

	c.AuthHandler = authhandler.New(
		c.AuthService,
	)
}

func (c *Container) composeUser() {

	c.UserRepository = userpostgres.New(
		c.DB,
	)

	c.UserService = userservice.New(
		c.UserRepository,
	)

	c.UserHandler = userhandler.New(
		c.UserService,
	)
}

// Compose builds the application's dependency graph.
func (c *Container) Compose() {

	//--------------------------------------------------
	// User
	//--------------------------------------------------

	c.composeUser()

	//--------------------------------------------------
	// Wallet
	//--------------------------------------------------

	c.composeWallet()

	//--------------------------------------------------
	// Transaction
	//--------------------------------------------------

	c.BuildTransaction()

	//--------------------------------------------------
	// Auth
	//--------------------------------------------------

	c.BuildAuth()

	//--------------------------------------------------
	// Product
	//--------------------------------------------------
	c.composeProduct()
	//--------------------------------------------------
	// Order
	//--------------------------------------------------
	c.composeOrder()
	//--------------------------------------------------
	// Purchase
	//--------------------------------------------------
	c.composePurchase()
	c.composePurchaseHandler()
	//
	// These will be migrated in the next batches.
}

func (c *Container) composeWallet() {

	c.WalletRepository = walletpostgres.New(
		c.DB,
	)

	c.WalletService = walletservice.New(
		c.WalletRepository,
	)

	c.WalletHandler = wallethandler.New(
		c.WalletService,
	)
}

func (c *Container) composeProduct() {

	c.ProductRepository = productpostgres.New(
		c.DB,
	)

	c.ProductService = productservice.New(
		c.ProductRepository,
	)

	c.ProductHandler = producthandler.New(
		c.ProductService,
	)
}

func (c *Container) composeOrder() {

	c.OrderRepository = orderpostgres.New(
		c.DB,
	)

	c.OrderService = orderservice.New(
		c.OrderRepository,
	)

	c.OrderHandler = orderhandler.New(
		c.OrderService,
	)
}

func (c *Container) composePurchase() {

	c.PurchaseService = purchase.New(
		c.DB,

		c.ProductRepository,
		c.WalletRepository,
		c.OrderRepository,
		c.TransactionRepository,

		c.ProductService,
		c.WalletService,
		c.OrderService,
		c.TransactionService,
	)
}

func (c *Container) composePurchaseHandler() {

	c.PurchaseHandler = purchasehandler.New(
		c.PurchaseService,
	)
}
