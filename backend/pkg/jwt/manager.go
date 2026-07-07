package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	secretKey      []byte
	accessDuration time.Duration
	refreshDuration time.Duration
}

// New creates a new JWT manager.
func New(
	secret string,
	accessDuration time.Duration,
	refreshDuration time.Duration,
) *Manager {
	return &Manager{
		secretKey:       []byte(secret),
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
	}
}



// AccessTokenDuration returns the configured access token lifetime.
func (m *Manager) AccessTokenDuration() time.Duration {
	return m.accessDuration
}

// RefreshTokenDuration returns the configured refresh token lifetime.
func (m *Manager) RefreshTokenDuration() time.Duration {
	return m.refreshDuration
}

// GenerateAccessToken generates a signed JWT access token.
func (m *Manager) GenerateAccessToken(
	userID string,
	piUID string,
	username string,
	role string,
) (string, error) {

	now := time.Now()

	claims := Claims{
		UserID:   userID,
		PiUID:    piUID,
		Username: username,
		Role:     role,

		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.accessDuration)),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(m.secretKey)
}

// GenerateRefreshToken generates a signed refresh token.
func (m *Manager) GenerateRefreshToken(
	userID string,
) (string, error) {

	now := time.Now()

	claims := jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(m.refreshDuration)),
		Subject:   userID,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(m.secretKey)
}
