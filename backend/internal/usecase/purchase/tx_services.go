package purchase

import (
	"github.com/JCKFinland/jck-connect/backend/pkg/database"

	//--------------------------------------------------
	// Product
	//--------------------------------------------------

	productpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/product/repository/postgres"
	productservice "github.com/JCKFinland/jck-connect/backend/internal/domain/product/service"

	//--------------------------------------------------
	// Wallet
	//--------------------------------------------------

	walletpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository/postgres"
	walletservice "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/service"

	//--------------------------------------------------
	// Order
	//--------------------------------------------------

	orderpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/order/repository/postgres"
	orderservice "github.com/JCKFinland/jck-connect/backend/internal/domain/order/service"

	//--------------------------------------------------
	// Transaction
	//--------------------------------------------------

	transactionpostgres "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/repository/postgres"
	transactionservice "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/service"
)

// transactionServices contains services that all participate
// in the same database transaction.
type transactionServices struct {
	productService     productservice.Service
	walletService      walletservice.Service
	orderService       orderservice.Service
	transactionService transactionservice.Service
}

// newTransactionServices creates transaction-scoped repositories
// and wraps them with their corresponding domain services.
func (s *service) newTransactionServices(
	tx database.DBTX,
) *transactionServices {

	//--------------------------------------------------
	// Product
	//--------------------------------------------------

	productRepository := productpostgres.NewTx(tx)

	//--------------------------------------------------
	// Wallet
	//--------------------------------------------------

	walletRepository := walletpostgres.NewTx(tx)

	//--------------------------------------------------
	// Order
	//--------------------------------------------------

	orderRepository := orderpostgres.NewTx(tx)

	//--------------------------------------------------
	// Transaction
	//--------------------------------------------------

	transactionRepository := transactionpostgres.NewTx(tx)

	return &transactionServices{
		productService: productservice.New(
			productRepository,
		),

		walletService: walletservice.New(
			walletRepository,
		),

		orderService: orderservice.New(
			orderRepository,
		),

		transactionService: transactionservice.New(
			transactionRepository,
		),
	}
}