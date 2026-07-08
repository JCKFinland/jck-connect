package handler
import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	ordermapper "github.com/JCKFinland/jck-connect/backend/internal/domain/order/mapper"

	"github.com/JCKFinland/jck-connect/backend/internal/middleware"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	"github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

// ListOrders returns all orders belonging to the authenticated user.
//
// Route:
//
//	GET /orders
func (h *Handler) ListOrders(
	c *gin.Context,
) {

	userID := middleware.UserID(c)

	if userID == "" {
		response.Unauthorized(
			c,
			sharedErrors.CodeUnauthorized,
			sharedErrors.MsgUnauthorized,
			"",
		)
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		response.BadRequest(
			c,
			sharedErrors.CodeBadRequest,
			sharedErrors.MsgBadRequest,
			"",
		)
		return
	}

	orders, err := h.service.ListByUserID(
		c.Request.Context(),
		userUUID,
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
		"Orders retrieved successfully.",
		ordermapper.ToOrderResponseList(orders),
	)
}
