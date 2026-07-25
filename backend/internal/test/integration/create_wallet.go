package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"

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
	// Build wallet
	//--------------------------------------------------

	wallet := fixture.Wallet(
		t,
		userID,
	)

	//--------------------------------------------------
	// Persist
	//--------------------------------------------------

	err := app.Container.WalletRepository.Create(
		context.Background(),
		wallet,
	)

	require.NoError(t, err)

	return wallet
}
