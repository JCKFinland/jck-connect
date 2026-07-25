package service

import (
	"context"

	"github.com/shopspring/decimal"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
	walletrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository"
)

type Service interface {
	GetByUserID(
		ctx context.Context,
		userID string,
	) (*walletentity.Wallet, error)

	Credit(
		ctx context.Context,
		userID string,
		amount decimal.Decimal,
	) error

	Debit(
		ctx context.Context,
		userID string,
		amount decimal.Decimal,
	) (*walletentity.Wallet, error)

	HasSufficientBalance(
		ctx context.Context,
		userID string,
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
