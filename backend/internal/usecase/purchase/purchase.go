package purchase

import (
	"context"

	"github.com/google/uuid"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"

	orderentity "github.com/JCKFinland/jck-connect/backend/internal/domain/order/entity"
	transactionentity "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/entity"

	"github.com/JCKFinland/jck-connect/backend/pkg/database"
)

// Purchase executes the purchase workflow.
func (s *service) Purchase(
	ctx context.Context,
	request PurchaseRequest,
) error {

	//--------------------------------------------------
	// Validate request
	//--------------------------------------------------

	if request.UserID == uuid.Nil {
		return sharedErrors.UserIDRequired(nil)
	}

	if request.ProductID == uuid.Nil {
		return sharedErrors.ProductIDRequired(nil)
	}

	//--------------------------------------------------
	// Execute transaction
	//--------------------------------------------------

	return s.txManager.WithTransaction(
		ctx,
		func(tx database.DBTX) error {

			//--------------------------------------------------
			// Transaction-scoped services
			//--------------------------------------------------

			txServices := s.newTransactionServices(tx)

			//--------------------------------------------------
			// Load Product
			//--------------------------------------------------

			product, err := txServices.productService.GetByID(
				ctx,
				request.ProductID,
			)
			if err != nil {
				return err
			}

			if product == nil {
				return sharedErrors.ProductNotFound(
					sharedErrors.ErrNotFound,
				)
			}

			//--------------------------------------------------
			// Load Wallet
			//--------------------------------------------------

			wallet, err := txServices.walletService.GetByUserID(
				ctx,
				request.UserID.String(),
			)
			if err != nil {
				return err
			}

			if wallet == nil {
				return sharedErrors.InsufficientWalletBalance(nil)
			}

			//--------------------------------------------------
			// Verify Balance
			//--------------------------------------------------

			hasBalance, err := txServices.walletService.HasSufficientBalance(
				ctx,
				request.UserID.String(),
				product.Price,
			)
			if err != nil {
				return err
			}

			if !hasBalance {
				return sharedErrors.InsufficientWalletBalance(nil)
			}

			//--------------------------------------------------
			// Create Order
			//--------------------------------------------------

			order := &orderentity.Order{
				UserID:    request.UserID,
				ProductID: request.ProductID,
				Amount:    product.Price,
				Currency:  product.Currency,
			}

			if err := txServices.orderService.Create(
				ctx,
				order,
			); err != nil {
				return err
			}

			//--------------------------------------------------
			// Debit Wallet
			//--------------------------------------------------

			balanceBefore := wallet.Balance

			wallet, err = txServices.walletService.Debit(
				ctx,
				request.UserID.String(),
				product.Price,
			)
			if err != nil {
				return err
			}

			balanceAfter := wallet.Balance

			//--------------------------------------------------
			// Create Ledger Transaction
			//--------------------------------------------------

			transaction := &transactionentity.Transaction{
				WalletID: wallet.ID,
				OrderID:  order.ID,

				BalanceBefore: balanceBefore,
				BalanceAfter:  balanceAfter,

				Type: transactionentity.TransactionTypeDebit,

				Amount: product.Price,

				Currency: product.Currency,

				Description: product.Name,

				Status: transactionentity.TransactionStatusCompleted,
			}

			if err := txServices.transactionService.Create(
				ctx,
				transaction,
			); err != nil {
				return err
			}

			return nil
		},
	)
}
