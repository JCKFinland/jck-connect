package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
)

func (r *repository) GetByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*entity.Wallet, error) {

	var wallet entity.Wallet

	var balance string

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			user_id,
			balance,
			currency,
			created_at,
			updated_at
		FROM wallets
		WHERE user_id=$1
		`,
		userID,
	).Scan(
		&wallet.ID,
		&wallet.UserID,
		&balance,
		&wallet.Currency,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	wallet.Balance = decimal.RequireFromString(balance)

	return &wallet, nil
}
