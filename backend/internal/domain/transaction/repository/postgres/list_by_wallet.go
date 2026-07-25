package postgres

import (
	"context"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// ListByWallet returns all transactions belonging to a wallet.
func (r *repository) ListByWallet(
	ctx context.Context,
	walletID string,
) ([]*transactionentity.Transaction, error) {

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
WHERE wallet_id = $1
ORDER BY created_at DESC
`

	rows, err := r.db.Query(
		ctx,
		query,
		walletID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]*transactionentity.Transaction, 0)

	for rows.Next() {

		transaction := &transactionentity.Transaction{}

		err := rows.Scan(
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

		transactions = append(
			transactions,
			transaction,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
