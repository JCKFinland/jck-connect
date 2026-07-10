package service

import (
	"context"

	"github.com/google/uuid"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// ListByWallet returns all wallet transactions.
func (s *service) ListByWallet(
	ctx context.Context,
	walletID uuid.UUID,
) ([]*transactionentity.Transaction, error) {

	return s.repository.ListByWallet(
		ctx,
		walletID,
	)
}