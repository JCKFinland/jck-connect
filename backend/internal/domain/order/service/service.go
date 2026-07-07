package service

import (
	"context"

	"github.com/google/uuid"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
	orderrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/order/repository"
)

// Service defines the business operations supported by the Order domain.
type Service interface {

	// Create creates a new order.
	Create(
		ctx context.Context,
		order *orderentity.Order,
	) error

	// GetByID returns an order by its internal ID.
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

	// Update updates an existing order.
	Update(
		ctx context.Context,
		order *orderentity.Order,
	) error
}

type service struct {
	repository orderrepo.Repository
}

// New creates a new Order service.
func New(
	repository orderrepo.Repository,
) Service {
	return &service{
		repository: repository,
	}
}
