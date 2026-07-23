package dto

// LoginRequest represents a Pi authentication request.
//
// The frontend authenticates the user using the official
// Pi SDK and sends only the access token to the backend.
//
// The backend never trusts client-supplied PiUID or
// PiUsername. Those values are retrieved directly from
// the Pi Platform after token verification.
type LoginRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
}