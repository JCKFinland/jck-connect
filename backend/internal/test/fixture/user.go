package fixture

import (
	"testing"

	userentity "github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
)

func User(
	t *testing.T,
) *userentity.User {

	t.Helper()

	return userentity.New(
		"pi-test-user",
		"testuser",
		"Integration Test User",
		"",
		"",
	)
}