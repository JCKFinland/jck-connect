package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
	"github.com/JCKFinland/jck-connect/backend/internal/test/fixture"
)

func CreateWallet(
	t *testing.T,
	app *TestApp,
	userID string,
) *walletentity.Wallet {

	t.Helper()

	//--------------------------------------------------
	// Convert User ID
	//--------------------------------------------------

	userUUID, err := uuid.Parse(userID)
	require.NoError(t, err)

	//--------------------------------------------------
	// Build wallet
	//--------------------------------------------------

	wallet := fixture.Wallet(
		t,
		userUUID,
	)

	//--------------------------------------------------
	// Persist
	//--------------------------------------------------

	err = app.Container.WalletRepository.Create(
		context.Background(),
		wallet,
	)

	require.NoError(t, err)

	return wallet
}