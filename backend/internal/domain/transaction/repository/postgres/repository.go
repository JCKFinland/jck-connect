package postgres

import (
	transactionrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/repository"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

type repository struct {
	db database.DBTX
}

// Compile-time interface check.
var _ transactionrepo.Repository = (*repository)(nil)

// New creates a PostgreSQL transaction repository.
func New(
	db *database.Database,
) transactionrepo.Repository {

	return &repository{
		db: db,
	}
}

// NewTx creates a PostgreSQL transaction repository using an existing database transaction.
func NewTx(
	tx database.DBTX,
) transactionrepo.Repository {

	return &repository{
		db: tx,
	}
}
