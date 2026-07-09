package service

import (
	"context"
	"time"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// Update updates an existing product.
func (s *service) Update(
	ctx context.Context,
	product *productentity.Product,
) error {

	product.UpdatedAt = time.Now().UTC()

	return s.repository.Update(
		ctx,
		product,
	)
}