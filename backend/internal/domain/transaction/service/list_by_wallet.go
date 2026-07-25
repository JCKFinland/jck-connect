package service

import (
	"context"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// ListByWallet returns all wallet transactions.
func (s *service) ListByWallet(
	ctx context.Context,
	walletID string,
) ([]*transactionentity.Transaction, error) {

	return s.repository.ListByWallet(
		ctx,
		walletID,
	)
}
