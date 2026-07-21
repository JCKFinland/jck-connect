package errors

func InvalidCredentials(err error) *AppError {
	return New(
		CodeAuthInvalidCredentials,
		MsgInvalidCredentials,
		err,
	)
}

func InvalidToken(err error) *AppError {
	return New(
		CodeAuthInvalidToken,
		MsgInvalidToken,
		err,
	)
}

func TokenExpired(err error) *AppError {
	return New(
		CodeAuthTokenExpired,
		MsgTokenExpired,
		err,
	)
}

func UserNotFound(err error) *AppError {
	return New(
		CodeAuthUserNotFound,
		MsgUserNotFound,
		err,
	)
}

func UserAlreadyExists(err error) *AppError {
	return New(
		CodeAuthUserExists,
		MsgUserAlreadyExists,
		err,
	)
}

func PiUIDRequired(err error) *AppError {
	return New(
		CodeValidationFailed,
		MsgPiUIDRequired,
		err,
	)
}

func PiUsernameRequired(err error) *AppError {
	return New(
		CodeValidationFailed,
		MsgPiUsernameRequired,
		err,
	)
}
