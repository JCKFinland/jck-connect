package service

import (
	"context"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
)

// GetByReference returns an order by its public reference.
func (s *service) GetByReference(
	ctx context.Context,
	reference string,
) (*orderentity.Order, error) {

	return s.repository.GetByReference(
		ctx,
		reference,
	)
}
