package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestPurchase_Success(t *testing.T) {

	//--------------------------------------------------
	// Test application
	//--------------------------------------------------

	app := NewTestApp(t)

	require.NotNil(t, app)
	require.NotNil(t, app.App)
	require.NotNil(t, app.Engine)
	require.NotNil(t, app.Container)

	//--------------------------------------------------
	// Reset database
	//--------------------------------------------------

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

	//--------------------------------------------------
	// Create wallet
	//--------------------------------------------------

	wallet := CreateWallet(
		t,
		app,
		user.ID,
	)

	//--------------------------------------------------
	// Fund wallet
	//--------------------------------------------------

	CreditWallet(
		t,
		app,
		user.ID,
		decimal.NewFromInt(100),
	)

	//--------------------------------------------------
	// Create product
	//--------------------------------------------------

	product := CreateProduct(
		t,
		app,
	)

	//--------------------------------------------------
	// Authentication
	//--------------------------------------------------

	token := AuthToken(
		t,
		app,
		user,
	)

	body := []byte(
		`{
		"product_id":"` + product.ID.String() + `"
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

	w := httptest.NewRecorder()

	app.Engine.ServeHTTP(
		w,
		req,
	)

	require.Equal(
		t,
		http.StatusCreated,
		w.Code,
	)

	//--------------------------------------------------
	// Verify
	//--------------------------------------------------

	require.NotNil(t, user)
	require.NotNil(t, wallet)
	require.NotNil(t, product)
}
