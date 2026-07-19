package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

func (s *service) Debit(
	ctx context.Context,
	userID uuid.UUID,
	amount decimal.Decimal,
) (*walletentity.Wallet, error) {

	if amount.LessThanOrEqual(decimal.Zero) {
		return nil, sharedErrors.New(
			sharedErrors.CodeBadRequest,
			sharedErrors.MsgBadRequest,
			nil,
		)
	}

	wallet, err := s.GetByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if err := wallet.Debit(amount); err != nil {
		return nil, err
	}

	if err := s.repository.Update(
		ctx,
		wallet,
	); err != nil {
		return nil, err
	}

	return wallet, nil
}