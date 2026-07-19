package errors

func OrderNotFound(err error) *AppError {
	return New(
		CodeOrderNotFound,
		MsgNotFound,
		err,
	)
}

func OrderExpired(err error) *AppError {
	return New(
		CodeOrderExpired,
		MsgOrderExpired,
		err,
	)
}