package fixture

import (
	"testing"

	"github.com/google/uuid"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
)

func Wallet(
	t *testing.T,
	userID uuid.UUID,
) *walletentity.Wallet {

	t.Helper()

	return walletentity.New(
		userID,
	)
}
