package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

func (s *service) HasSufficientBalance(
	ctx context.Context,
	userID uuid.UUID,
	amount decimal.Decimal,
) (bool, error) {

	if amount.LessThanOrEqual(decimal.Zero) {
		return false, sharedErrors.New(
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
		return false, err
	}

	return wallet.Balance.GreaterThanOrEqual(amount), nil
}