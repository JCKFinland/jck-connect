package entity

import "time"

// UserRole defines the role assigned to a user.
type UserRole string

const (
	RoleUser  UserRole = "USER"
	RoleAdmin UserRole = "ADMIN"
)

// UserStatus defines the current status of a user account.
type UserStatus string

const (
	StatusActive   UserStatus = "ACTIVE"
	StatusInactive UserStatus = "INACTIVE"
	StatusBlocked  UserStatus = "BLOCKED"
)

// User represents a jck-connect user.
type User struct {
	ID string `db:"id" json:"id"`

	PiUID      string `db:"pi_uid" json:"pi_uid"`
	PiUsername string `db:"pi_username" json:"pi_username"`

	DisplayName string `db:"display_name" json:"display_name"`

	Email string `db:"email" json:"email,omitempty"`

	PhoneNumber string `db:"phone_number" json:"phone_number,omitempty"`

	Role UserRole `db:"role" json:"role"`

	Status UserStatus `db:"status" json:"status"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`

	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
