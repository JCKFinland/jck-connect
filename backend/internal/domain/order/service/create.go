package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

// Create creates a new order.
func (s *service) Create(
	ctx context.Context,
	order *orderentity.Order,
) error {

	if order == nil {
		return sharedErrors.New(
			sharedErrors.CodeBadRequest,
			sharedErrors.MsgBadRequest,
			nil,
		)
	}

	now := time.Now().UTC()

	//--------------------------------------------------
	// Initialize immutable fields
	//--------------------------------------------------

	if order.ID == uuid.Nil {
		order.ID = uuid.New()
	}

	if order.Reference == "" {
		order.Reference = generateReference()
	}

	//--------------------------------------------------
	// Initialize timestamps
	//--------------------------------------------------

	order.CreatedAt = now
	order.UpdatedAt = now

	//--------------------------------------------------
	// Initial status
	//--------------------------------------------------

	if order.Status == "" {
		order.Status = orderentity.StatusPending
	}

	return s.repository.Create(
		ctx,
		order,
	)
}
