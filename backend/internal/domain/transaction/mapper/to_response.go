package mapper

import (
	transactiondto "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/dto"
	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// ToResponse converts a Transaction entity into a response DTO.
func ToResponse(
	transaction *transactionentity.Transaction,
) *transactiondto.TransactionResponse {

	if transaction == nil {
		return nil
	}

	return &transactiondto.TransactionResponse{
		ID:            transaction.ID.String(),
		OrderID:       transaction.OrderID.String(),
		WalletID:      transaction.WalletID,
		Type:          string(transaction.Type),
		Status:        string(transaction.Status),
		Amount:        transaction.Amount.String(),
		Currency:      transaction.Currency,
		BalanceBefore: transaction.BalanceBefore.String(),
		BalanceAfter:  transaction.BalanceAfter.String(),
		Reference:     transaction.Reference,
		Description:   transaction.Description,
		CreatedAt:     transaction.CreatedAt,
	}
}
