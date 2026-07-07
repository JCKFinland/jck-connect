package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	walletmapper "github.com/JCKFinland/jck-connect/backend/internal/domain/wallet/mapper"

	"github.com/JCKFinland/jck-connect/backend/internal/middleware"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	response "github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

func (h *Handler) GetWallet(
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

	wallet, err := h.service.GetByUserID(
		c.Request.Context(),
		userUUID,
	)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.Success(
		c,
		"Wallet retrieved successfully.",
		walletmapper.ToWalletResponse(wallet),
	)
}