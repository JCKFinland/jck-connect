package service

import (
	"context"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
	transactionrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/repository"
)

// Service defines the Transaction business logic.
type Service interface {
	Create(
		ctx context.Context,
		transaction *transactionentity.Transaction,
	) error

	GetByID(
		ctx context.Context,
		id string,
	) (*transactionentity.Transaction, error)

	GetByReference(
		ctx context.Context,
		reference string,
	) (*transactionentity.Transaction, error)

	ListByWallet(
		ctx context.Context,
		walletID string,
	) ([]*transactionentity.Transaction, error)
}

type service struct {
	repository transactionrepo.Repository
}

// New creates a new Transaction service.
func New(
	repository transactionrepo.Repository,
) Service {

	return &service{
		repository: repository,
	}
}
