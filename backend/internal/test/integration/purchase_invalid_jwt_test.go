package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPurchase_InvalidJWT(t *testing.T) {

	//--------------------------------------------------
	// Test application
	//--------------------------------------------------

	app := NewTestApp(t)

	ResetDatabase(
		t,
		app.App.DB(),
	)

	//--------------------------------------------------
	// Build request
	//--------------------------------------------------

	body := []byte(
		`{
			"product_id":"` + uuid.New().String() + `"
		}`,
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/purchases",
		bytes.NewBuffer(body),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	req.Header.Set(
		"Authorization",
		"Bearer this-is-not-a-valid-jwt",
	)

	//--------------------------------------------------
	// Execute
	//--------------------------------------------------

	w := httptest.NewRecorder()

	app.Engine.ServeHTTP(
		w,
		req,
	)

	//--------------------------------------------------
	// Assert
	//--------------------------------------------------

	require.Equal(
		t,
		http.StatusUnauthorized,
		w.Code,
	)
}
