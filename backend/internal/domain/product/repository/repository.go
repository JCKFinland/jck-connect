// Package repository defines persistence contracts
// for the Product domain.
package repository

import (
	"context"

	"github.com/google/uuid"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// Repository defines the Product persistence contract.
type Repository interface {
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
