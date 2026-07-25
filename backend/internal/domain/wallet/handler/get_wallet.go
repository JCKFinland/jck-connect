package handler

import (
	"github.com/gin-gonic/gin"

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

	wallet, err := h.service.GetByUserID(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.Success(
		c,
		response.MsgWalletRetrieved,
		walletmapper.ToWalletResponse(wallet),
	)
}
