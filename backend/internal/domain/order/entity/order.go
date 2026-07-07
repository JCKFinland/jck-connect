package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Order status values.
const (
	StatusPending    = "PENDING"
	StatusProcessing = "PROCESSING"
	StatusCompleted  = "COMPLETED"
	StatusFailed     = "FAILED"
)

// Order represents a customer purchase.
//
// An Order is the central business record for every purchase made
// through JCK Connect.
//
// The charged Amount is stored on the order so historical purchases
// remain accurate even if product pricing changes later.
type Order struct {
	ID uuid.UUID

	UserID    uuid.UUID
	ProductID uuid.UUID

	// Human-readable order reference.
	Reference string

	// Amount charged for this order.
	Amount decimal.Decimal

	// Currency used for the purchase.
	Currency string

	// Current order status.
	Status string

	CreatedAt time.Time
	UpdatedAt time.Time
}
