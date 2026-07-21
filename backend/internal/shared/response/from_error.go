package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	sharedErrors "github.com/JCKFinland/jck-connect/backend/internal/shared/errors"
)

// FromError converts an application error into a standardized HTTP response.
func FromError(
	c *gin.Context,
	err error,
) {
	if err == nil {
		return
	}

	appErr, ok := err.(*sharedErrors.AppError)
	if !ok {
		InternalServerError(
			c,
			sharedErrors.CodeInternalServerError,
			sharedErrors.MsgInternalServer,
			err.Error(),
		)
		return
	}

	status := http.StatusInternalServerError

	switch appErr.Code {

	case sharedErrors.CodeBadRequest:
		status = http.StatusBadRequest

	case sharedErrors.CodeValidationFailed:
		status = http.StatusUnprocessableEntity

	case sharedErrors.CodeUnauthorized,
		sharedErrors.CodeAuthInvalidCredentials,
		sharedErrors.CodeAuthInvalidToken,
		sharedErrors.CodeAuthTokenExpired:
		status = http.StatusUnauthorized

	case sharedErrors.CodeForbidden:
		status = http.StatusForbidden

	case sharedErrors.CodeNotFound,
		sharedErrors.CodeAuthUserNotFound,
		sharedErrors.CodeProductNotFound,
		sharedErrors.CodeOrderNotFound,
		sharedErrors.CodeTransactionNotFound,
		sharedErrors.CodeWalletNotFound:
		status = http.StatusNotFound

	case sharedErrors.CodeWalletInsufficientBalance:
		status = http.StatusUnprocessableEntity

	case sharedErrors.CodeConflict,
		sharedErrors.CodeAuthUserExists,
		sharedErrors.CodeTransactionDuplicate,
		sharedErrors.CodeOrderAlreadyCompleted,
		sharedErrors.CodeWalletAlreadyExists:
		status = http.StatusConflict
	}

	errorText := ""
	if appErr.Err != nil {
		errorText = appErr.Err.Error()
	}

	Fail(
		c,
		status,
		appErr.Code,
		appErr.Message,
		errorText,
	)
}
