package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

// ResetDatabase removes all application data from the test database.
func ResetDatabase(
	t *testing.T,
	db *database.Database,
) {
	t.Helper()

	_, err := db.Exec(
		context.Background(),
		`
		TRUNCATE TABLE
			transactions,
			orders,
			wallets,
			products,
			users
		RESTART IDENTITY
		CASCADE;
		`,
	)

	require.NoError(t, err)
}
