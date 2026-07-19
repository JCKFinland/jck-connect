package errors

import "errors"

// Code extracts the application error code.
func Code(err error) string {
	var appErr *AppError

	if errors.As(err, &appErr) {
		return appErr.Code
	}

	return CodeInternalServerError
}

// Message extracts the application error message.
func Message(err error) string {
	var appErr *AppError

	if errors.As(err, &appErr) {
		return appErr.Message
	}

	return MsgInternalServer
}