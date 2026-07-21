package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Credit adds funds to the wallet.
func (w *Wallet) Credit(
	amount decimal.Decimal,
) {
	w.Balance = w.Balance.Add(amount)
	w.UpdatedAt = time.Now().UTC()
}
