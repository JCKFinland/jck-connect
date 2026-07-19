package errors

func TransactionNotFound(err error) *AppError {
	return New(
		CodeTransactionNotFound,
		MsgNotFound,
		err,
	)
}

func TransactionFailed(err error) *AppError {
	return New(
		CodeTransactionFailed,
		MsgTransactionFailed,
		err,
	)
}