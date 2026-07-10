package postgres

import (
	"context"

	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"
)

// Create inserts a new transaction.
func (r *repository) Create(
	ctx context.Context,
	transaction *transactionentity.Transaction,
) error {

	const query = `
INSERT INTO transactions
(
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
)
VALUES
(
	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
)
`

	_, err := r.db.Exec(
		ctx,
		query,
		transaction.ID,
		transaction.OrderID,
		transaction.WalletID,
		transaction.Type,
		transaction.Status,
		transaction.Amount.String(),
		transaction.Currency,
		transaction.BalanceBefore.String(),
		transaction.BalanceAfter.String(),
		transaction.Reference,
		transaction.Description,
		transaction.CreatedAt,
	)

	return err
}
