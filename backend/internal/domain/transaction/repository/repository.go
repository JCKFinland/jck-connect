package repository

import (
	"context"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// Repository defines transaction persistence behavior.
type Repository interface {

	// Create stores a new transaction.
	Create(
		ctx context.Context,
		transaction *transactionentity.Transaction,
	) error

	// GetByID returns a transaction by its ID.
	GetByID(
		ctx context.Context,
		id string,
	) (*transactionentity.Transaction, error)

	// GetByReference returns a transaction by its business reference.
	GetByReference(
		ctx context.Context,
		reference string,
	) (*transactionentity.Transaction, error)

	// ListByWallet returns all transactions belonging to a wallet.
	ListByWallet(
		ctx context.Context,
		WalletID string,
	) ([]*transactionentity.Transaction, error)
}
