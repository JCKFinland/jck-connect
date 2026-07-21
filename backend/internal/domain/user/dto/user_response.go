package dto

import "time"

// UserResponse is the public representation of a user.
type UserResponse struct {
	ID string `json:"id"`

	PiUID string `json:"pi_uid"`

	PiUsername string `json:"pi_username"`

	DisplayName string `json:"display_name"`

	Email string `json:"email,omitempty"`

	PhoneNumber string `json:"phone_number,omitempty"`

	Role string `json:"role"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
}
