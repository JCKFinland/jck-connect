package database

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Migrate executes all *.up.sql files in the migrations directory.
func Migrate(db *Database, migrationsDir string) error {
	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.up.sql"))
	if err != nil {
		return fmt.Errorf("scan migrations: %w", err)
	}

	sort.Strings(files)

	ctx := context.Background()

	for _, file := range files {
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("read %s: %w", file, err)
		}

		sql := strings.TrimSpace(string(sqlBytes))
		if sql == "" {
			continue
		}

		if _, err := db.Exec(ctx, sql); err != nil {
			return fmt.Errorf("execute %s: %w", filepath.Base(file), err)
		}
	}

	return nil
}