package dto

// CreateUserRequest represents a new user registration request.
type CreateUserRequest struct {
	PiUID string `json:"pi_uid" binding:"required"`

	PiUsername string `json:"pi_username" binding:"required"`

	DisplayName string `json:"display_name"`

	Email string `json:"email"`

	PhoneNumber string `json:"phone_number"`
}
