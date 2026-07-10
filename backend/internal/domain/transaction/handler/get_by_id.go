package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	transactionmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/mapper"
)

// GetByID returns a transaction by ID.
func (h *Handler) GetByID(
	c *gin.Context,
) {

	id, err := uuid.Parse(
		c.Param("id"),
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid transaction id",
			},
		)
		return
	}

	transaction, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "transaction not found",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		transactionmapper.ToResponse(transaction),
	)
}