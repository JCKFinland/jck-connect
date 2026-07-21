package service

import (
	"context"

	"github.com/google/uuid"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
)

// GetByID returns an order by its internal ID.
func (s *service) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*orderentity.Order, error) {

	return s.repository.GetByID(
		ctx,
		id,
	)
}
