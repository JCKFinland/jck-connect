package pi

import "time"

// Config contains all Pi Platform settings.
type Config struct {

	// BaseURL is the Pi API endpoint.
	BaseURL string

	// APIKey is the application's Pi Platform API key.
	APIKey string

	// Timeout controls request timeout.
	Timeout time.Duration
}

// DefaultConfig returns default Pi configuration.
//
// Sandbox URL will later be replaced by the
// production endpoint before deployment.
func DefaultConfig() Config {

	return Config{
		BaseURL: "https://api.minepi.com",
		APIKey:  "",
		Timeout: 10 * time.Second,
	}
}
