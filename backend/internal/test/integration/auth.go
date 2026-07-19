package integration

import (
	"testing"

	"github.com/stretchr/testify/require"

	userentity "github.com/JCKFinland/jck-connect/backend/internal/domain/user/entity"
)

// AuthToken returns a valid Bearer token for the supplied user.
func AuthToken(
	t *testing.T,
	app *TestApp,
	user *userentity.User,
) string {

	t.Helper()

	token, err := app.Container.JWTManager.GenerateAccessToken(
		user.ID,
		user.PiUID,
		user.PiUsername,
		string(user.Role),
	)

	require.NoError(t, err)

	return "Bearer " + token
}