package service

import (
	"context"
	"fmt"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// Create creates a new transaction.
func (s *service) Create(
	ctx context.Context,
	transaction *transactionentity.Transaction,
) error {

	if transaction == nil {
		return fmt.Errorf("transaction is required")
	}

	return s.repository.Create(
		ctx,
		transaction,
	)
}