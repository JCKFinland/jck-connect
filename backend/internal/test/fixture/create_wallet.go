package fixture

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
	walletrepo "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/repository"
)

func CreateWallet(
	t *testing.T,
	repo walletrepo.Repository,
	userID uuid.UUID,
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