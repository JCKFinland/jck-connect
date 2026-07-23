package pi

// VerifiedUser represents a verified Pi user returned
// after successful access token validation.
type VerifiedUser struct {
	UID      string
	Username string
}