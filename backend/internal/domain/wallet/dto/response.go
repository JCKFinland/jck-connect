package dto

// WalletResponse is returned to clients.
type WalletResponse struct {
	Balance  string `json:"balance"`
	Currency string `json:"currency"`
}