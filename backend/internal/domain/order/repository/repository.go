package repository

import (
	"context"

	"github.com/google/uuid"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
)

// Repository defines the persistence contract for orders.
type Repository interface {

	// Create persists a new order.
	Create(
		ctx context.Context,
		order *orderentity.Order,
	) error

	// GetByID returns an order by its ID.
	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*orderentity.Order, error)

	// GetByReference returns an order by its public reference.
	GetByReference(
		ctx context.Context,
		reference string,
	) (*orderentity.Order, error)

	// ListByUserID returns all orders belonging to a user.
	ListByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) ([]*orderentity.Order, error)

	// Update persists changes to an existing order.
	Update(
		ctx context.Context,
		order *orderentity.Order,
	) error
}
