package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

func (s *service) Debit(
	ctx context.Context,
	userID uuid.UUID,
	amount decimal.Decimal,
) error {

	if amount.LessThanOrEqual(decimal.Zero) {
		return sharedErrors.New(
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
		return err
	}

	if err := wallet.Debit(amount); err != nil {
		return err
	}

	return s.repository.Update(
		ctx,
		wallet,
	)
}