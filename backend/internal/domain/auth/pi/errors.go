package pi

import "errors"

var (

	// ErrInvalidAccessToken is returned when
	// Pi rejects the supplied access token.
	ErrInvalidAccessToken = errors.New("invalid pi access token")

	// ErrExpiredAccessToken indicates that
	// the supplied access token has expired.
	ErrExpiredAccessToken = errors.New("expired pi access token")

	// ErrUnauthorized indicates that the
	// request is not authorized by Pi.
	ErrUnauthorized = errors.New("pi authentication failed")

	// ErrNetworkFailure indicates communication
	// failure with the Pi Platform.
	ErrNetworkFailure = errors.New("unable to communicate with Pi Platform")
)
