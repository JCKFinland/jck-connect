package dto

// UpdateRequest represents a profile update request.
type UpdateRequest struct {
	DisplayName string `json:"display_name" binding:"omitempty,min=2,max=100"`

	Email string `json:"email" binding:"omitempty,email,max=255"`

	PhoneNumber string `json:"phone_number" binding:"omitempty,max=32"`
}
