package mapper

import (
	walletdto "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/dto"
	"github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
)

// ToWalletResponse converts a Wallet entity into a response DTO.
func ToWalletResponse(
	wallet *entity.Wallet,
) walletdto.WalletResponse {

	if wallet == nil {
		return walletdto.WalletResponse{}
	}

	return walletdto.WalletResponse{
		Balance:  wallet.Balance.String(),
		Currency: wallet.Currency,
	}
}
