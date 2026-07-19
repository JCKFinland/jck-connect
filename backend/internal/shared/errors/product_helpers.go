package errors

func ProductNotFound(err error) *AppError {
	return New(
		CodeProductNotFound,
		MsgProductNotFound,
		err,
	)
}

func ProductInactive(err error) *AppError {
	return New(
		CodeProductInactive,
		"Product is inactive.",
		err,
	)
}