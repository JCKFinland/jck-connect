package postgres

import (
	"context"

	"github.com/google/uuid"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// GetByID returns a transaction by its ID.
func (r *repository) GetByID(
	ctx context.Context,
	id uuid.UUID,
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
WHERE id = $1
`

	transaction := &transactionentity.Transaction{}

	err := r.db.QueryRow(
		ctx,
		query,
		id,
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
