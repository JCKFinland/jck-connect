package postgres

import (
orderrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/order/repository"
"github.com/JCKFinland/jck-connect/backend/pkg/database"

)

// Repository implements the PostgreSQL order repository.
type Repository struct {
db database.DBTX
}

// Compile-time interface check.
var _ orderrepo.Repository = (*Repository)(nil)

// New creates a new PostgreSQL order repository.
//
// The repository depends on database.DBTX instead of a concrete database
// connection so it can operate with either:
//
// - *database.Database
// - pgx.Tx
//
// This allows the same repository implementation to participate in
// transactional workflows without any code changes.
func New(
    db *database.Database,
) orderrepo.Repository {

    return &Repository{
        db: db,
    }
}

func NewTx(
tx database.DBTX,
) orderrepo.Repository {

return &Repository{
	db: tx,
}


}