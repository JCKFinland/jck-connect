package errors

func PiLoginFailed(err error) *AppError {
	return New(
		CodePiLoginFailed,
		MsgPiLoginFailed,
		err,
	)
}

func PiPaymentFailed(err error) *AppError {
	return New(
		CodePiPaymentVerification,
		MsgPiPaymentFailed,
		err,
	)
}
