package postgres

import (
	productrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/product/repository"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

// Repository implements the Product repository using PostgreSQL.
type Repository struct {
	db database.DBTX
}

// Compile-time interface check.
var _ productrepo.Repository = (*Repository)(nil)

// New creates a new PostgreSQL Product repository.
func New(
	db database.DBTX,
) productrepo.Repository {
	return &Repository{
		db: db,
	}
}
