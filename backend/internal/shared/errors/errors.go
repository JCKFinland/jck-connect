package errors

import "errors"

// Common reusable application errors.
var (
	ErrNotFound           = errors.New("resource not found")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrConflict           = errors.New("resource already exists")
	ErrValidationFailed   = errors.New("validation failed")
	ErrInternalServer     = errors.New("internal server error")
)

// AppError represents a reusable application error.
type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return e.Message
}

// Unwrap allows errors.Is() and errors.As() to work.
func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates a new application error.
func New(
	code string,
	message string,
	err error,
) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}