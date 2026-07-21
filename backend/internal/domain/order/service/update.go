package service

import (
	"context"
	"time"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

// Update updates an existing order.
func (s *service) Update(
	ctx context.Context,
	order *orderentity.Order,
) error {

	if order == nil {
		return sharedErrors.BadRequest(nil)
	}

	//--------------------------------------------------
	// Refresh modification timestamp
	//--------------------------------------------------

	order.UpdatedAt = time.Now().UTC()

	return s.repository.Update(
		ctx,
		order,
	)
}
