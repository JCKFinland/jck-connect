package purchase

import (
	"context"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

// TransactionManager executes work inside a database transaction.
//
// It is intentionally small so it can be implemented by
// *database.Database in production and by a fake in tests.
type TransactionManager interface {
	WithTransaction(
		ctx context.Context,
		fn func(database.DBTX) error,
	) error
}
