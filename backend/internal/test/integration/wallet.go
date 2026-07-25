package integration

import (
	"context"
	"testing"
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

	err := app.Container.WalletService.Credit(
		context.Background(),
		userID,
		amount,
	)

	require.NoError(t, err)
}
