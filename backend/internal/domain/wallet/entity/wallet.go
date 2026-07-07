package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Balance   decimal.Decimal
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}