package service

import (
	"context"

	"github.com/google/uuid"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
)

// ListByUserID returns all orders belonging to a user.
func (s *service) ListByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]*orderentity.Order, error) {

	return s.repository.ListByUserID(
		ctx,
		userID,
	)
}
