package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
	"github.com/JCKFinland/jck-connect/backend/internal/domain/user/repository"
	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

type Service interface {
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetByPiUID(ctx context.Context, piUID string) (*entity.User, error)
	GetByPiUsername(ctx context.Context, username string) (*entity.User, error)
	List(ctx context.Context) ([]*entity.User, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repo repository.Repository
}

// New creates a new user service.
func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

// Create creates a new user.
func (s *service) Create(
	ctx context.Context,
	user *entity.User,
) error {

	if user == nil {
		return sharedErrors.BadRequest(
			sharedErrors.BadRequest(nil),
		)
	}

	// Normalize input.
	user.PiUID = strings.TrimSpace(user.PiUID)
	user.PiUsername = strings.TrimSpace(user.PiUsername)
	user.DisplayName = strings.TrimSpace(user.DisplayName)
	user.Email = strings.TrimSpace(user.Email)
	user.PhoneNumber = strings.TrimSpace(user.PhoneNumber)

	// Validate required fields.
	if user.PiUID == "" {
		return sharedErrors.PiUIDRequired(nil)
	}

	if user.PiUsername == "" {
		return sharedErrors.PiUsernameRequired(nil)
	}

	now := time.Now().UTC()

	// Generate a UUID only if one has not already been assigned.
	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	// Apply default values only when not already provided.
	if user.Role == "" {
		user.Role = entity.RoleUser
	}

	if user.Status == "" {
		user.Status = entity.StatusActive
	}

	// Preserve an existing creation timestamp if present.
	if user.CreatedAt.IsZero() {
		user.CreatedAt = now
	}

	// Always update the modification timestamp.
	user.UpdatedAt = now

	return s.repo.Create(ctx, user)
}

// Update updates an existing user.
func (s *service) Update(
	ctx context.Context,
	user *entity.User,
) error {

	if user == nil {
		return sharedErrors.BadRequest(
			errors.New("user cannot be nil"),
		)
	}

	// Load the current record from the database.
	current, err := s.repo.FindByID(ctx, strings.TrimSpace(user.ID))
	if err != nil {
		return err
	}

	// Update only the fields that a user is allowed to change.
	current.DisplayName = strings.TrimSpace(user.DisplayName)
	current.Email = strings.TrimSpace(user.Email)
	current.PhoneNumber = strings.TrimSpace(user.PhoneNumber)

	current.UpdatedAt = time.Now().UTC()

	return s.repo.Update(ctx, current)
}

// GetByID returns a user by ID.
func (s *service) GetByID(
	ctx context.Context,
	id string,
) (*entity.User, error) {
	return s.repo.FindByID(ctx, strings.TrimSpace(id))
}

// GetByPiUID returns a user by Pi UID.
func (s *service) GetByPiUID(
	ctx context.Context,
	piUID string,
) (*entity.User, error) {
	return s.repo.FindByPiUID(ctx, strings.TrimSpace(piUID))
}

// GetByPiUsername returns a user by Pi username.
func (s *service) GetByPiUsername(
	ctx context.Context,
	username string,
) (*entity.User, error) {
	return s.repo.FindByPiUsername(ctx, strings.TrimSpace(username))
}

// List returns all users.
func (s *service) List(
	ctx context.Context,
) ([]*entity.User, error) {
	return s.repo.List(ctx)
}

// Delete deactivates or removes a user.
//
// NOTE:
// Permanent deletion is intentionally deferred.
// The repository currently returns "not supported".
// This will later become a soft-delete operation.
func (s *service) Delete(
	ctx context.Context,
	id string,
) error {
	return s.repo.Delete(ctx, strings.TrimSpace(id))
}
