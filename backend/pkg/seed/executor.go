package seed

import (
	"context"
	"fmt"
	"log"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

// RunSeeds executes every SQL seed file inside a single transaction.
func RunSeeds(db *database.Database) error {

	files, err := LoadFiles()
	if err != nil {
		return err
	}

	if len(files) == 0 {
		log.Println("No seed files found.")
		return nil
	}

	ctx := context.Background()

	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if tx != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	log.Println("----------------------------------------")
	log.Println("Running database seeds...")
	log.Println("----------------------------------------")

	for _, file := range files {

		log.Printf("Running %s", file)

		sqlBytes, err := ReadFile(file)
		if err != nil {
			return fmt.Errorf("read %s: %w", file, err)
		}

		_, err = tx.Exec(
			ctx,
			string(sqlBytes),
		)
		if err != nil {
			return fmt.Errorf("execute %s: %w", file, err)
		}

		log.Printf("✓ %s", file)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	tx = nil

	log.Println("----------------------------------------")
	log.Printf("Seed completed successfully.")
	log.Printf("Files executed: %d", len(files))
	log.Println("----------------------------------------")

	return nil
}
