package dto

import "time"

// RefreshResponse represents newly issued tokens.
type RefreshResponse struct {
	AccessToken string `json:"access_token"`

	RefreshToken string `json:"refresh_token"`

	TokenType string `json:"token_type"`

	ExpiresAt time.Time `json:"expires_at"`
}
