package dto

// UpdateRequest represents a profile update request.
type UpdateRequest struct {
	DisplayName string `json:"display_name"`

	Email string `json:"email"`

	PhoneNumber string `json:"phone_number"`
}