package fixture

import (
	"testing"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
)

func Wallet(
	t *testing.T,
	userID string,
) *walletentity.Wallet {

	t.Helper()

	return walletentity.New(
		userID,
	)
}
