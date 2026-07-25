package handler

import (
	"github.com/gin-gonic/gin"

	transactionmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/mapper"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	response "github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

// GetByID returns a transaction by ID.
func (h *Handler) GetByID(
	c *gin.Context,
) {

	id := c.Param("id")
	if id == "" {
		response.BadRequest(
			c,
			sharedErrors.CodeBadRequest,
			sharedErrors.MsgBadRequest,
			"transaction id is required",
		)
		return
	}

	transaction, err := h.service.GetByID(
		c.Request.Context(),
		id,
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
