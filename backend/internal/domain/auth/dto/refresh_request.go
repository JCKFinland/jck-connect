package dto

// RefreshRequest represents a refresh-token request.
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
