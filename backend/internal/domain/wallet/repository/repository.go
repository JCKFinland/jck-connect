package repository

import (
	"context"

	"github.com/google/uuid"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
)

type Repository interface {

	// Returns a wallet by user ID.
	GetByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*walletentity.Wallet, error)

	// Returns and locks a wallet row inside a transaction.
	GetByUserIDForUpdate(
		ctx context.Context,
		userID uuid.UUID,
	) (*walletentity.Wallet, error)

	// Persists wallet changes.
	Update(
		ctx context.Context,
		wallet *walletentity.Wallet,
	) error

	// Creates a wallet.
	Create(
		ctx context.Context,
		wallet *walletentity.Wallet,
	) error
}