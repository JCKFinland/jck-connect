package entity

import (
	"time"

	"github.com/google/uuid"
)

// New creates a valid User entity.
func New(
	piUID string,
	piUsername string,
	displayName string,
	email string,
	phoneNumber string,
) *User {

	now := time.Now().UTC()

	return &User{
		ID: uuid.NewString(),

		PiUID:       piUID,
		PiUsername:  piUsername,
		DisplayName: displayName,

		Email:       email,
		PhoneNumber: phoneNumber,

		Role: RoleUser,

		Status: StatusActive,

		CreatedAt: now,
		UpdatedAt: now,
	}
}
