package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	ordermapper "github.com/JCKFinland/jck-connect/backend/internal/domain/order/mapper"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	"github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

// GetOrder returns an order by its ID.
//
// Route:
//
//	GET /orders/:id
func (h *Handler) GetOrder(
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
			"",
		)
		return
	}

	order, err := h.service.GetByID(
		c.Request.Context(),
		id,
	)
	if err != nil {
		response.FromError(
			c,
			err,
		)
		return
	}

	response.Success(
		c,
		"Order retrieved successfully.",
		ordermapper.ToOrderResponse(order),
	)
}
