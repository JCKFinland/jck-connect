package postgres

import (
	walletrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

// Repository implements the PostgreSQL wallet repository.
type repository struct {
	db database.DBTX
}

// Compile-time interface check.
var _ walletrepo.Repository = (*repository)(nil)

// New creates a PostgreSQL wallet repository.
func New(
	db *database.Database,
) walletrepo.Repository {

	return &repository{
		db: db,
	}
}

// NewTx creates a PostgreSQL wallet repository using an existing transaction.
func NewTx(
	tx database.DBTX,
) walletrepo.Repository {

	return &repository{
		db: tx,
	}
}