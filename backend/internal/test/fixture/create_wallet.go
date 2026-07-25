package fixture

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
	walletrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository"
)

func CreateWallet(
	t *testing.T,
	repo walletrepo.Repository,
	userID string,
) *walletentity.Wallet {

	t.Helper()

	wallet := walletentity.New(
		userID,
	)

	err := repo.Create(
		context.Background(),
		wallet,
	)

	require.NoError(t, err)

	return wallet
}
