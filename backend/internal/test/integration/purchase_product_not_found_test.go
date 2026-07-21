package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestPurchase_ProductNotFound(t *testing.T) {

	//--------------------------------------------------
	// Test application
	//--------------------------------------------------

	app := NewTestApp(t)

	ResetDatabase(
		t,
		app.App.DB(),
	)

	//--------------------------------------------------
	// Create user
	//--------------------------------------------------

	user := CreateUser(
		t,
		app,
	)

	CreateWallet(
		t,
		app,
		user.ID,
	)

	CreditWallet(
		t,
		app,
		user.ID,
		decimal.NewFromInt(100),
	)

	//--------------------------------------------------
	// Authentication
	//--------------------------------------------------

	token := AuthToken(
		t,
		app,
		user,
	)

	//--------------------------------------------------
	// Non-existent product
	//--------------------------------------------------

	productID := uuid.New()

	body := []byte(
		`{
			"product_id":"` + productID.String() + `"
		}`,
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/purchases",
		bytes.NewBuffer(body),
	)

	req.Header.Set(
		"Authorization",
		token,
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
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
		http.StatusNotFound,
		w.Code,
	)
}
