package transaction

import (
	"github.com/gin-gonic/gin"

	transactionhandler "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/handler"
)

// RegisterRoutes registers transaction routes.
func RegisterRoutes(
	router gin.IRoutes,
	handler *transactionhandler.Handler,
) {

	router.GET(
		"/transactions/:id",
		handler.GetByID,
	)

	router.GET(
		"/transactions/reference/:reference",
		handler.GetByReference,
	)

	router.GET(
		"/wallets/:walletId/transactions",
		handler.ListByWallet,
	)
}
