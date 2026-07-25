package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID        string
	UserID    string
	Balance   decimal.Decimal
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
