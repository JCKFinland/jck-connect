package integration

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func CreditWallet(
	t *testing.T,
	app *TestApp,
	userID string,
	amount decimal.Decimal,
) {
	t.Helper()

	userUUID, err := uuid.Parse(userID)
	require.NoError(t, err)

	err = app.Container.WalletService.Credit(
		context.Background(),
		userUUID,
		amount,
	)

	require.NoError(t, err)
}