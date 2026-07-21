package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Product represents a purchasable product or service.
type Product struct {
	ID uuid.UUID

	// Business identity
	Code     string
	Name     string
	Category string
	Provider string

	// Pricing
	Price    decimal.Decimal
	Currency string

	// Availability
	Active bool

	// Audit
	CreatedAt time.Time
	UpdatedAt time.Time
}
