package service

import (
	"context"

	"github.com/shopspring/decimal"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

func (s *service) HasSufficientBalance(
	ctx context.Context,
	userID string,
	amount decimal.Decimal,
) (bool, error) {

	if amount.LessThanOrEqual(decimal.Zero) {
		return false, sharedErrors.BadRequest(nil)
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
