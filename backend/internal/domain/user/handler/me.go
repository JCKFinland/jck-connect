package handler

import (
	"github.com/gin-gonic/gin"

	usermapper "github.com/JCKFinland/jck-connect/backend/internal/domain/user/mapper"
	"github.com/JCKFinland/jck-connect/backend/internal/middleware"
	response "github.com/JCKFinland/jck-connect/backend/internal/shared/response"
	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

// Me returns the authenticated user's profile.
func (h *Handler) Me(c *gin.Context) {

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

	u, err := h.service.GetByID(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		response.FromError(c, err)
		return
	}

	userResponse := usermapper.ToUserResponse(u)

	response.Success(
		c,
		"User profile retrieved successfully.",
		userResponse,
	)
}