package service

import (
	"context"

	"github.com/google/uuid"

	walletentity "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/entity"
	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

func (s *service) GetByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*walletentity.Wallet, error) {

	wallet, err := s.repository.GetByUserID(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	if wallet == nil {
		return nil, sharedErrors.WalletNotFound(nil)
	}

	return wallet, nil
}