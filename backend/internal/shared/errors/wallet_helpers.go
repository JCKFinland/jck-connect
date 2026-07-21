package errors

func WalletNotFound(err error) *AppError {
	return New(
		CodeWalletNotFound,
		MsgWalletNotFound,
		err,
	)
}

func WalletAlreadyExists(err error) *AppError {
	return New(
		CodeWalletAlreadyExists,
		MsgWalletAlreadyExists,
		err,
	)
}

func InsufficientWalletBalance(err error) *AppError {
	return New(
		CodeWalletInsufficientBalance,
		MsgWalletInsufficientBalance,
		err,
	)
}

func InvalidWalletCurrency(err error) *AppError {
	return New(
		CodeWalletInvalidCurrency,
		MsgWalletInvalidCurrency,
		err,
	)
}
