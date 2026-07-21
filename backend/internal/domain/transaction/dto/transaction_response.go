package dto

import "time"

// TransactionResponse represents a transaction response.
type TransactionResponse struct {
	ID string `json:"id"`

	OrderID string `json:"order_id"`

	WalletID string `json:"wallet_id"`

	Type string `json:"type"`

	Status string `json:"status"`

	Amount string `json:"amount"`

	Currency string `json:"currency"`

	BalanceBefore string `json:"balance_before"`

	BalanceAfter string `json:"balance_after"`

	Reference string `json:"reference"`

	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
}
