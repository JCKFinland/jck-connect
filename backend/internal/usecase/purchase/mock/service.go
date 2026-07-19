package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	purchase "github.com/JCKFinland/jck-connect/backend/internal/usecase/purchase"
)

type Service struct {
	mock.Mock
}

func (m *Service) Purchase(
	ctx context.Context,
	request purchase.PurchaseRequest,
) error {

	args := m.Called(
		ctx,
		request,
	)

	return args.Error(0)
}