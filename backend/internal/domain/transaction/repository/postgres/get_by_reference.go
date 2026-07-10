package postgres

import (
	"context"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// GetByReference returns a transaction by its business reference.
func (r *repository) GetByReference(
	ctx context.Context,
	reference string,
) (*transactionentity.Transaction, error) {

	const query = `
SELECT
	id,
	order_id,
	wallet_id,
	type,
	status,
	amount,
	currency,
	balance_before,
	balance_after,
	reference,
	description,
	created_at
FROM transactions
WHERE reference = $1
`

	transaction := &transactionentity.Transaction{}

	err := r.db.QueryRow(
		ctx,
		query,
		reference,
	).Scan(
		&transaction.ID,
		&transaction.OrderID,
		&transaction.WalletID,
		&transaction.Type,
		&transaction.Status,
		&transaction.Amount,
		&transaction.Currency,
		&transaction.BalanceBefore,
		&transaction.BalanceAfter,
		&transaction.Reference,
		&transaction.Description,
		&transaction.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return transaction, nil
}