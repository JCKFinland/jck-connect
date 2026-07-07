package entity

import (
	"time"

	"github.com/shopspring/decimal"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

// Debit removes funds from the wallet.
func (w *Wallet) Debit(
	amount decimal.Decimal,
) error {

	if w.Balance.LessThan(amount) {
		return sharedErrors.InsufficientWalletBalance(nil)
	}

	w.Balance = w.Balance.Sub(amount)
	w.UpdatedAt = time.Now().UTC()

	return nil
}