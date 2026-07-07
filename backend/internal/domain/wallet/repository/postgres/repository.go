package postgres

import (
	walletrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

type Repository struct {
	db database.DBTX
}

// Compile-time interface check.
var _ walletrepo.Repository = (*Repository)(nil)

// New creates a PostgreSQL wallet repository.
func New(
	db database.DBTX,
) walletrepo.Repository {
	return &Repository{
		db: db,
	}
}