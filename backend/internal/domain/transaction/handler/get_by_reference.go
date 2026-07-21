package handler

import (
	"github.com/gin-gonic/gin"

	transactionmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/mapper"

	response "github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

// GetByReference returns a transaction by business reference.
func (h *Handler) GetByReference(
	c *gin.Context,
) {

	transaction, err := h.service.GetByReference(
		c.Request.Context(),
		c.Param("reference"),
	)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.Success(
		c,
		response.MsgTransactionRetrieved,
		transactionmapper.ToResponse(transaction),
	)
}
