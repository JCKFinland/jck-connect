package purchase

import (
	"context"

	"github.com/google/uuid"

	orderrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/order/repository"
	orderservice "github.com/JCKFinland/jck-connect/backend/internal/domain/order/service"

	productrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/product/repository"
	productservice "github.com/JCKFinland/jck-connect/backend/internal/domain/product/service"

	transactionrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/repository"
	transactionservice "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/service"

	walletrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository"
	walletservice "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/service"
)

// PurchaseRequest contains everything required to perform a purchase.
type PurchaseRequest struct {
	UserID    uuid.UUID
	ProductID uuid.UUID
}

// Service defines the purchase use case.
type Service interface {
	Purchase(
		ctx context.Context,
		request PurchaseRequest,
	) error
}

type service struct {
	txManager TransactionManager

	//--------------------------------------------------
	// Repository dependencies
	//--------------------------------------------------

	productRepository     productrepo.Repository
	walletRepository      walletrepo.Repository
	orderRepository       orderrepo.Repository
	transactionRepository transactionrepo.Repository

	//--------------------------------------------------
	// Domain services
	//--------------------------------------------------

	productService     productservice.Service
	walletService      walletservice.Service
	orderService       orderservice.Service
	transactionService transactionservice.Service
}

// New creates the Purchase use case.
func New(
	txManager TransactionManager,

	productRepository productrepo.Repository,
	walletRepository walletrepo.Repository,
	orderRepository orderrepo.Repository,
	transactionRepository transactionrepo.Repository,

	productService productservice.Service,
	walletService walletservice.Service,
	orderService orderservice.Service,
	transactionService transactionservice.Service,
) Service {

	return &service{
		txManager: txManager,

		productRepository:     productRepository,
		walletRepository:      walletRepository,
		orderRepository:       orderRepository,
		transactionRepository: transactionRepository,

		productService:     productService,
		walletService:      walletService,
		orderService:       orderService,
		transactionService: transactionService,
	}
}
