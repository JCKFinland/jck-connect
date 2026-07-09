package service

import (
	"context"

	"github.com/google/uuid"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
	productrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/product/repository"
)

// Service defines the Product business logic.
type Service interface {
	Create(
		ctx context.Context,
		product *productentity.Product,
	) error

	Update(
		ctx context.Context,
		product *productentity.Product,
	) error

	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*productentity.Product, error)

	List(
		ctx context.Context,
	) ([]*productentity.Product, error)
}

type service struct {
	repository productrepo.Repository
}

// New creates a new Product service.
func New(
	repository productrepo.Repository,
) Service {
	return &service{
		repository: repository,
	}
}
