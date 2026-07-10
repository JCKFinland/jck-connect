package service

import (
	"context"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// GetByReference returns a transaction by its business reference.
func (s *service) GetByReference(
	ctx context.Context,
	reference string,
) (*transactionentity.Transaction, error) {

	return s.repository.GetByReference(
		ctx,
		reference,
	)
}