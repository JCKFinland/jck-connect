package dto

// UpdateUserRequest represents an update to a user's profile.
type UpdateUserRequest struct {
	DisplayName string `json:"display_name"`

	Email string `json:"email"`

	PhoneNumber string `json:"phone_number"`
}