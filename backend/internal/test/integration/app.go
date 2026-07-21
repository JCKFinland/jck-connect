package integration

import (
	"net/http"
	"testing"

	"github.com/JCKFinland/jck-connect/backend/internal/app"
	"github.com/JCKFinland/jck-connect/backend/internal/config"
	"github.com/JCKFinland/jck-connect/backend/internal/container"
	"github.com/stretchr/testify/require"
)

// TestApp wraps the application used by integration tests.
type TestApp struct {
	App       *app.App
	Container *container.Container
	Engine    http.Handler
}

// NewTestApp creates a fully initialized application for integration tests.
func NewTestApp(
	t *testing.T,
) *TestApp {

	t.Helper()

	cfg := config.Load()

	//--------------------------------------------------
	// Always use the integration database
	//--------------------------------------------------

	cfg.DBName = "jck_connect_test"

	a, err := app.NewWithConfig(cfg)
	require.NoError(t, err)

	return &TestApp{
		App:       a,
		Container: a.Container(),
		Engine:    a.Engine(),
	}
}
