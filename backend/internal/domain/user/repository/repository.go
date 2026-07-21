package repository

import (
	"context"

	"github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
)

// Repository defines the user persistence contract.
type Repository interface {
	Create(ctx context.Context, user *entity.User) error

	Update(ctx context.Context, user *entity.User) error

	FindByID(ctx context.Context, id string) (*entity.User, error)

	FindByPiUID(ctx context.Context, piUID string) (*entity.User, error)

	FindByPiUsername(ctx context.Context, username string) (*entity.User, error)

	List(ctx context.Context) ([]*entity.User, error)

	Delete(ctx context.Context, id string) error
}
