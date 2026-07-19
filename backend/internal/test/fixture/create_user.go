package fixture

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	userentity "github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
	userrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/user/repository"
)

func CreateUser(
	t *testing.T,
	repo userrepo.Repository,
) *userentity.User {

	t.Helper()

	user := userentity.New(
		"pi-test-user",
		"testuser",
		"Integration Test User",
		"",
		"",
	)

	err := repo.Create(
		context.Background(),
		user,
	)

	require.NoError(t, err)

	return user
}