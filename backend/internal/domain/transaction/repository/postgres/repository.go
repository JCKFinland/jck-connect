package postgres

import (
	transactionrepository "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/repository"
	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

type repository struct {
	db *database.Database
}

// New creates a new PostgreSQL transaction repository.
func New(
	db *database.Database,
) transactionrepository.Repository {

	return &repository{
		db: db,
	}
}
