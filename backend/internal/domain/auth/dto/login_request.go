package dto

// LoginRequest represents a login request.
//
// During the initial implementation, authentication
// is based on Pi identity.
type LoginRequest struct {
	PiUID string `json:"pi_uid" binding:"required"`

	PiUsername string `json:"pi_username" binding:"required"`
}