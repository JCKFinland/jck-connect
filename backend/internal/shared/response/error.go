package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fail sends a standardized error response.
func Fail(
	c *gin.Context,
	status int,
	code string,
	message string,
	err string,
) {
	c.JSON(
		status,
		Response{
			Success: false,
			Code:    code,
			Message: message,
			Error:   err,
		},
	)
}

// BadRequest returns HTTP 400.
func BadRequest(
	c *gin.Context,
	code string,
	message string,
	err string,
) {
	Fail(c, http.StatusBadRequest, code, message, err)
}

// Unauthorized returns HTTP 401.
func Unauthorized(
	c *gin.Context,
	code string,
	message string,
	err string,
) {
	Fail(c, http.StatusUnauthorized, code, message, err)
}

// Forbidden returns HTTP 403.
func Forbidden(
	c *gin.Context,
	code string,
	message string,
	err string,
) {
	Fail(c, http.StatusForbidden, code, message, err)
}

// NotFound returns HTTP 404.
func NotFound(
	c *gin.Context,
	code string,
	message string,
	err string,
) {
	Fail(c, http.StatusNotFound, code, message, err)
}

// Conflict returns HTTP 409.
func Conflict(
	c *gin.Context,
	code string,
	message string,
	err string,
) {
	Fail(c, http.StatusConflict, code, message, err)
}

// UnprocessableEntity returns HTTP 422.
func UnprocessableEntity(
	c *gin.Context,
	code string,
	message string,
	err string,
) {
	Fail(c, http.StatusUnprocessableEntity, code, message, err)
}

// InternalServerError returns HTTP 500.
func InternalServerError(
	c *gin.Context,
	code string,
	message string,
	err string,
) {
	Fail(c, http.StatusInternalServerError, code, message, err)
}
