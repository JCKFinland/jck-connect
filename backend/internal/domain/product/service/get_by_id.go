package service

import (
	"context"

	"github.com/google/uuid"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// GetByID returns a product by its ID.
func (s *service) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*productentity.Product, error) {

	return s.repository.GetByID(
		ctx,
		id,
	)
}
