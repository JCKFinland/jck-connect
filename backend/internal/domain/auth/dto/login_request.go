package dto

// LoginRequest represents a Pi authentication request.
type LoginRequest struct {
	PiUID string `json:"pi_uid" binding:"required,min=3,max=128"`

	PiUsername string `json:"pi_username" binding:"required,min=2,max=64"`
}
