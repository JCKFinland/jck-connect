package errors

import (
    "errors"
    "net/http"
)

// HTTPStatus converts an AppError into an HTTP status code.
func HTTPStatus(err error) int {

    var appErr *AppError

    if !errors.As(err, &appErr) {
        return http.StatusInternalServerError
    }

    switch appErr.Code {

    case CodeBadRequest,
        CodeValidationFailed:
        return http.StatusBadRequest

    case CodeUnauthorized,
        CodeAuthInvalidCredentials,
        CodeAuthInvalidToken,
        CodeAuthTokenExpired:
        return http.StatusUnauthorized

    case CodeForbidden:
        return http.StatusForbidden

    case CodeNotFound,
        CodeProductNotFound,
        CodeWalletNotFound,
        CodeOrderNotFound,
        CodeTransactionNotFound,
        CodeAuthUserNotFound:
        return http.StatusNotFound

    case CodeConflict,
        CodeWalletAlreadyExists,
        CodeAuthUserExists,
        CodeTransactionDuplicate:
        return http.StatusConflict

    default:
        return http.StatusInternalServerError
    }
}