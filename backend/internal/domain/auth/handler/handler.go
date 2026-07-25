package handler

import (
	"github.com/gin-gonic/gin"

	authdto "github.com/JCKFinland/jck-connect/backend/internal/domain/auth/dto"
	authservice "github.com/JCKFinland/jck-connect/backend/internal/domain/auth/service"
	"github.com/JCKFinland/jck-connect/backend/internal/shared/response"
)

type Handler struct {
	authService authservice.Service
}

// New creates a new authentication handler.
func New(
	authService authservice.Service,
) *Handler {
	return &Handler{
		authService: authService,
	}
}

// Login authenticates a Pi user using a Pi access token.
//
// POST /api/v1/auth/pi-login

func (h *Handler) PiLogin(c *gin.Context) {
	var req authdto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(
			c,
			"BAD_REQUEST",
			"Invalid request payload.",
			err.Error(),
		)
		return
	}

	res, err := h.authService.Login(
		c.Request.Context(),
		&req,
	)
	if err != nil {
		response.FromError(c, err)
		return
	}

	response.Success(
		c,
		response.MsgLoginSuccessful,
		res,
	)
}
