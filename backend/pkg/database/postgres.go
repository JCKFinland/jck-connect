package database

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"

    "github.com/JCKFinland/jck-connect/backend/internal/config"
)

type Database struct {
	Pool *pgxpool.Pool
}

// New creates a PostgreSQL connection pool.
func New(cfg *config.Config) (*Database, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	log.Printf(
		"Connected to PostgreSQL (%s:%s/%s)",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	return &Database{
		Pool: pool,
	}, nil
}

// Close gracefully closes the PostgreSQL connection pool.
func (db *Database) Close() {
	if db == nil || db.Pool == nil {
		return
	}

	db.Pool.Close()
}
