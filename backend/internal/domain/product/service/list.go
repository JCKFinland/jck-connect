package service

import (
	"context"

	productentity "github.com/JCKFinland/jck-connect/backend/internal/domain/product/entity"
)

// List returns all products.
func (s *service) List(
	ctx context.Context,
) ([]*productentity.Product, error) {

	return s.repository.List(
		ctx,
	)
}
