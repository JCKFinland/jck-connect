package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	transactionmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/mapper"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	response "github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

// ListByWallet returns wallet transaction history.
func (h *Handler) ListByWallet(
	c *gin.Context,
) {

	walletID, err := uuid.Parse(
		c.Param("walletId"),
	)
	if err != nil {
		response.BadRequest(
			c,
			sharedErrors.CodeBadRequest,
			sharedErrors.MsgBadRequest,
			"invalid wallet id",
		)
		return
	}

	transactions, err := h.service.ListByWallet(
		c.Request.Context(),
		walletID,
	)
	if err != nil {
		response.FromError(c, err)
		return
	}

	items := make([]any, 0, len(transactions))

	for _, transaction := range transactions {
		items = append(
			items,
			transactionmapper.ToResponse(transaction),
		)
	}

	response.Success(
		c,
		response.MsgTransactionsRetrieved,
		items,
	)
}
