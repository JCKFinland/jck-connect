package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// Create creates a new product.
func (s *service) Create(
	ctx context.Context,
	product *productentity.Product,
) error {

	now := time.Now().UTC()

	if product.ID == uuid.Nil {
		product.ID = uuid.New()
	}

	product.CreatedAt = now
	product.UpdatedAt = now

	// New products are active by default.
	product.Active = true

	return s.repository.Create(
		ctx,
		product,
	)
}
