package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	transactionmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/mapper"
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