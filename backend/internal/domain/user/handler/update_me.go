package handler

import (
	"github.com/gin-gonic/gin"

	userdto "github.com/JCKFinland/jck-connect/backend/internal/domain/user/dto"
	usermapper "github.com/JCKFinland/jck-connect/backend/internal/domain/user/mapper"
	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"

	"github.com/JCKFinland/jck-connect/backend/internal/middleware"

	response "github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

// UpdateMe updates the authenticated user's profile.
func (h *Handler) UpdateMe(c *gin.Context) {

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

	var request userdto.UpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.BadRequest(
			c,
			sharedErrors.CodeBadRequest,
			sharedErrors.MsgBadRequest,
			err.Error(),
		)
		return
	}

	user, err := h.service.GetByID(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		response.FromError(c, err)
		return
	}

	user.DisplayName = request.DisplayName
	user.Email = request.Email
	user.PhoneNumber = request.PhoneNumber

	if err := h.service.Update(
		c.Request.Context(),
		user,
	); err != nil {
		response.FromError(c, err)
		return
	}

	response.Success(
		c,
		response.MsgProfileUpdated,
		usermapper.ToUserResponse(user),
	)
}
