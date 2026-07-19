package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
	walletrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository"
)

type Service interface {

	GetByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*walletentity.Wallet, error)

	Credit(
		ctx context.Context,
		userID uuid.UUID,
		amount decimal.Decimal,
	) error

	Debit(
		ctx context.Context,
		userID uuid.UUID,
		amount decimal.Decimal,
	) (*walletentity.Wallet, error)

	HasSufficientBalance(
		ctx context.Context,
		userID uuid.UUID,
		amount decimal.Decimal,
	) (bool, error)
}

type service struct {
	repository walletrepo.Repository
}

func New(
	repository walletrepo.Repository,
) Service {
	return &service{
		repository: repository,
	}
}