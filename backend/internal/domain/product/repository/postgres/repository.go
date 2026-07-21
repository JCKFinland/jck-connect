package postgres

import (
	productrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/product/repository"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

type repository struct {
	db database.DBTX
}

// Compile-time interface check.
var _ productrepo.Repository = (*repository)(nil)

// New creates a new PostgreSQL Product repository.
func New(
	db *database.Database,
) productrepo.Repository {

	return &repository{
		db: db,
	}
}

// NewTx creates a PostgreSQL Product repository using an existing database transaction.
func NewTx(
	tx database.DBTX,
) productrepo.Repository {

	return &repository{
		db: tx,
	}
}
