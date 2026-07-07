package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// WithTransaction executes fn inside a database transaction.
func (db *Database) WithTransaction(
	ctx context.Context,
	fn func(tx pgx.Tx) error,
) error {

	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}