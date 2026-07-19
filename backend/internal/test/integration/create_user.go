package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/JCKFinland/jck-connect/backend/internal/test/fixture"
	userentity "github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
)

func CreateUser(
	t *testing.T,
	app *TestApp,
) *userentity.User {

	t.Helper()

	user := fixture.User(t)

	err := app.Container.UserRepository.Create(
		context.Background(),
		user,
	)

	require.NoError(t, err)

	return user
}