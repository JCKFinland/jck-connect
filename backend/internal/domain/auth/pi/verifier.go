package pi

import (
	"context"
	"fmt"
)

// Verifier verifies Pi Authentication tokens.
type Verifier interface {

	// Verify validates a Pi access token and returns
	// the verified Pi user.
	Verify(
		ctx context.Context,
		accessToken string,
	) (*VerifiedUser, error)
}

type verifier struct {
	client Client
}

// NewVerifier creates a new Pi verifier.
func NewVerifier(client Client) Verifier {

	return &verifier{
		client: client,
	}
}

// Verify validates an access token using the Pi Platform.
func (v *verifier) Verify(
	ctx context.Context,
	accessToken string,
) (*VerifiedUser, error) {

	user, err := v.client.VerifyAccessToken(
		ctx,
		accessToken,
	)
	if err != nil {

		switch err {

		case ErrUnauthorized:
			return nil, ErrUnauthorized

		case ErrInvalidAccessToken:
			return nil, ErrInvalidAccessToken

		case ErrExpiredAccessToken:
			return nil, ErrExpiredAccessToken

		case ErrNetworkFailure:
			return nil, ErrNetworkFailure

		default:
			return nil, fmt.Errorf(
				"verify Pi token: %w",
				err,
			)
		}
	}

	if user == nil {
		return nil, ErrUnauthorized
	}

	if user.UID == "" {
		return nil, ErrUnauthorized
	}

	return user, nil
}
