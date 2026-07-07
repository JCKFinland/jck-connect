package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// DBTX represents something that can execute SQL.
//
// Both *pgxpool.Pool and pgx.Tx satisfy this interface.
type DBTX interface {

	Exec(
		ctx context.Context,
		sql string,
		arguments ...any,
	) (pgconn.CommandTag, error)

	Query(
		ctx context.Context,
		sql string,
		args ...any,
	) (pgx.Rows, error)

	QueryRow(
		ctx context.Context,
		sql string,
		args ...any,
	) pgx.Row
}

func (db *Database) Exec(
	ctx context.Context,
	sql string,
	args ...any,
) (pgconn.CommandTag, error) {
	return db.Pool.Exec(ctx, sql, args...)
}

func (db *Database) Query(
	ctx context.Context,
	sql string,
	args ...any,
) (pgx.Rows, error) {
	return db.Pool.Query(ctx, sql, args...)
}

func (db *Database) QueryRow(
	ctx context.Context,
	sql string,
	args ...any,
) pgx.Row {
	return db.Pool.QueryRow(ctx, sql, args...)
}

