package seed

import "github.com/JCKFinland/jck-connect/backend/pkg/database"

// Run executes all SQL seed files.
func Run(db *database.Database) error {
	return RunSeeds(db)
}