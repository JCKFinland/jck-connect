package pi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client defines communication with the Pi Platform.
type Client interface {

	// VerifyAccessToken sends the access token to the Pi Platform
	// and returns the verified Pi user.
	VerifyAccessToken(
		ctx context.Context,
		accessToken string,
	) (*VerifiedUser, error)
}

type client struct {
	httpClient *http.Client
	config     Config
}

// NewClient creates a new Pi API client.
func NewClient(cfg Config) Client {

	return &client{
		httpClient: &http.Client{
			Timeout: cfg.Timeout,
		},
		config: cfg,
	}
}

// VerifyAccessToken validates an access token with the Pi Platform.
//
// NOTE:
// The endpoint below is a placeholder until we integrate the
// official Pi Authentication API.
//
// During local development this function simply demonstrates
// the HTTP communication layer.
func (c *client) VerifyAccessToken(
	ctx context.Context,
	accessToken string,
) (*VerifiedUser, error) {

	if accessToken == "" {
		return nil, ErrInvalidAccessToken
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.config.BaseURL,
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+accessToken,
	)

	req.Header.Set(
		"X-API-Key",
		c.config.APIKey,
	)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, ErrNetworkFailure
	}

	defer resp.Body.Close()

	switch resp.StatusCode {

	case http.StatusUnauthorized:
		return nil, ErrUnauthorized

	case http.StatusForbidden:
		return nil, ErrInvalidAccessToken

	case http.StatusOK:

		var user VerifiedUser

		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			return nil, fmt.Errorf(
				"decode pi response: %w",
				err,
			)
		}

		return &user, nil

	default:
		return nil, fmt.Errorf(
			"unexpected Pi response: %d",
			resp.StatusCode,
		)
	}
}
