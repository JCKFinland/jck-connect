package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

// Create creates a new transaction.
func (s *service) Create(
	ctx context.Context,
	transaction *transactionentity.Transaction,
) error {

	if transaction == nil {
		return sharedErrors.BadRequest(nil)
	}

	now := time.Now().UTC()

	//--------------------------------------------------
	// Initialize immutable fields
	//--------------------------------------------------

	if transaction.ID == uuid.Nil {
		transaction.ID = uuid.New()
	}

	if transaction.Reference == "" {
		transaction.Reference = generateReference()
	}

	//--------------------------------------------------
	// Initialize timestamps
	//--------------------------------------------------

	transaction.CreatedAt = now

	//--------------------------------------------------
	// Initial status
	//--------------------------------------------------

	if transaction.Status == "" {
		transaction.Status = transactionentity.TransactionStatusPending
	}

	return s.repository.Create(
		ctx,
		transaction,
	)
}
