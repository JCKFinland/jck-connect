package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	transactionmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/mapper"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	response "github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

// GetByID returns a transaction by ID.
func (h *Handler) GetByID(
	c *gin.Context,
) {

	id, err := uuid.Parse(
		c.Param("id"),
	)
	if err != nil {
		response.BadRequest(
			c,
			sharedErrors.CodeBadRequest,
			sharedErrors.MsgBadRequest,
			"invalid transaction id",
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
