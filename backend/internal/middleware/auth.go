package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
	"github.com/JCKFinland/jck-connect/backend/internal/shared/response"
	"github.com/JCKFinland/jck-connect/backend/pkg/jwt"
)

const (
	ContextUserID   = "user_id"
	ContextPiUID    = "pi_uid"
	ContextUsername = "username"
	ContextRole     = "role"
)

// Auth validates JWT access tokens.
func Auth(
	jwtManager *jwt.Manager,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(
				c,
				sharedErrors.CodeUnauthorized,
				sharedErrors.MsgUnauthorized,
				"authorization header is required",
			)
			c.Abort()
			return
		}

		const bearer = "Bearer "

		if !strings.HasPrefix(authHeader, bearer) {
			response.Unauthorized(
				c,
				sharedErrors.CodeUnauthorized,
				sharedErrors.MsgUnauthorized,
				"invalid authorization header",
			)
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, bearer)

		claims, err := jwtManager.ValidateToken(token)
		if err != nil {
			response.Unauthorized(
				c,
				sharedErrors.CodeAuthInvalidToken,
				sharedErrors.MsgInvalidToken,
				err.Error(),
			)
			c.Abort()
			return
		}

		c.Set(ContextUserID, claims.UserID)
		c.Set(ContextPiUID, claims.PiUID)
		c.Set(ContextUsername, claims.Username)
		c.Set(ContextRole, claims.Role)

		c.Next()
	}
}

// UserID returns the authenticated user's ID.
func UserID(c *gin.Context) string {
	value, ok := c.Get(ContextUserID)
	if !ok {
		return ""
	}

	userID, ok := value.(string)
	if !ok {
		return ""
	}

	return userID
}

// PiUID returns the authenticated user's Pi UID.
func PiUID(c *gin.Context) string {
	value, ok := c.Get(ContextPiUID)
	if !ok {
		return ""
	}

	piUID, ok := value.(string)
	if !ok {
		return ""
	}

	return piUID
}

// Username returns the authenticated user's Pi username.
func Username(c *gin.Context) string {
	value, ok := c.Get(ContextUsername)
	if !ok {
		return ""
	}

	username, ok := value.(string)
	if !ok {
		return ""
	}

	return username
}

// Role returns the authenticated user's role.
func Role(c *gin.Context) string {
	value, ok := c.Get(ContextRole)
	if !ok {
		return ""
	}

	role, ok := value.(string)
	if !ok {
		return ""
	}

	return role
}
