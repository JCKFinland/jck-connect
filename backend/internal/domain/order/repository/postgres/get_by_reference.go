package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

func (r *Repository) GetByReference(
	ctx context.Context,
	reference string,
) (*orderentity.Order, error) {

	const query = `
	SELECT
		id,
		user_id,
		product_id,
		reference,
		amount,
		currency,
		status,
		created_at,
		updated_at
	FROM orders
	WHERE reference = $1
	`

	var order orderentity.Order
	var amount string

	err := r.db.QueryRow(
		ctx,
		query,
		reference,
	).Scan(
		&order.ID,
		&order.UserID,
		&order.ProductID,
		&order.Reference,
		&amount,
		&order.Currency,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sharedErrors.OrderNotFound(
				sharedErrors.ErrNotFound,
			)
		}

		return nil, err
	}

	order.Amount = decimal.RequireFromString(amount)

	return &order, nil
}
