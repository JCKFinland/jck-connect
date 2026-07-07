package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims represents the JWT claims for jck-connect.
type Claims struct {
	UserID   string `json:"user_id"`
	PiUID    string `json:"pi_uid,omitempty"`
	Username string `json:"username"`
	Role     string `json:"role"`

	jwt.RegisteredClaims
}