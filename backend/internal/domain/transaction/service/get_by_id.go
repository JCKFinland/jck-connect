package service

import (
	"context"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// GetByID returns a transaction by its ID.
func (s *service) GetByID(
	ctx context.Context,
	id string,
) (*transactionentity.Transaction, error) {

	return s.repository.GetByID(
		ctx,
		id,
	)
}
