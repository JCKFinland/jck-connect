package dto

import (
	"time"

	userdto "github.com/JCKFinland/jck-connect/backend/internal/domain/user/dto"
)

// LoginResponse represents a successful authentication.
type LoginResponse struct {
	AccessToken string `json:"access_token"`

	RefreshToken string `json:"refresh_token"`

	TokenType string `json:"token_type"`

	ExpiresAt time.Time `json:"expires_at"`

	User userdto.UserResponse `json:"user"`
}
