package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ProductResponse represents a product returned by the API.
type ProductResponse struct {
	ID uuid.UUID `json:"id"`

	Code     string `json:"code"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Provider string `json:"provider"`

	Price    decimal.Decimal `json:"price"`
	Currency string          `json:"currency"`

	Active bool `json:"active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}