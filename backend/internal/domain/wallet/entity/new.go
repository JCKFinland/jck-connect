package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// New creates a valid Wallet entity.
func New(
	userID uuid.UUID,
) *Wallet {

	now := time.Now().UTC()

	return &Wallet{
		ID: uuid.New(),

		UserID: userID,

		Balance: decimal.Zero,

		Currency: "PI",

		CreatedAt: now,
		UpdatedAt: now,
	}
}
