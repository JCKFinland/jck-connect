package errors

func BadRequest(err error) *AppError {
	return New(
		CodeBadRequest,
		MsgBadRequest,
		err,
	)
}

func ValidationFailed(err error) *AppError {
	return New(
		CodeValidationFailed,
		MsgValidationFailed,
		err,
	)
}

func Unauthorized(err error) *AppError {
	return New(
		CodeUnauthorized,
		MsgUnauthorized,
		err,
	)
}

func Forbidden(err error) *AppError {
	return New(
		CodeForbidden,
		MsgForbidden,
		err,
	)
}

func Conflict(err error) *AppError {
	return New(
		CodeConflict,
		MsgConflict,
		err,
	)
}

func InternalServer(err error) *AppError {
	return New(
		CodeInternalServerError,
		MsgInternalServer,
		err,
	)
}

func NotFound(err error) *AppError {
	return New(
		CodeNotFound,
		MsgNotFound,
		err,
	)
}

func UserIDRequired(err error) *AppError {
	return New(
		CodeValidationFailed,
		"User ID is required.",
		err,
	)
}

func ProductIDRequired(err error) *AppError {
	return New(
		CodeValidationFailed,
		"Product ID is required.",
		err,
	)
}
