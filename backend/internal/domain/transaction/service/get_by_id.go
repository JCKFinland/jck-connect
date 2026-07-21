package service

import (
	"context"

	"github.com/google/uuid"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// GetByID returns a transaction by its ID.
func (s *service) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*transactionentity.Transaction, error) {

	return s.repository.GetByID(
		ctx,
		id,
	)
}
