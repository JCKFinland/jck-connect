package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	transactionmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/transaction/mapper"
)

// ListByWallet returns wallet transaction history.
func (h *Handler) ListByWallet(
	c *gin.Context,
) {

	walletID, err := uuid.Parse(
		c.Param("walletId"),
	)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid wallet id",
			},
		)
		return
	}

	transactions, err := h.service.ListByWallet(
		c.Request.Context(),
		walletID,
	)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	response := make([]any, 0, len(transactions))

	for _, transaction := range transactions {
		response = append(
			response,
			transactionmapper.ToResponse(transaction),
		)
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}