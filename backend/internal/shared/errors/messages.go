package errors

const (

	// Generic

	MsgSuccess            = "Request completed successfully."
	MsgBadRequest         = "The request is invalid."
	MsgUnauthorized       = "Authentication required."
	MsgForbidden          = "You are not authorized to perform this action."
	MsgNotFound           = "The requested resource was not found."
	MsgConflict           = "The requested operation cannot be completed."
	MsgValidationFailed   = "Validation failed."
	MsgInternalServer     = "An internal server error occurred."
	MsgEndpointNotFound   = "The requested endpoint does not exist."
	MsgPiUIDRequired      = "Pi UID is required."
	MsgPiUsernameRequired = "Pi Username is required."

	// Authentication

	MsgInvalidCredentials = "Invalid username or password."
	MsgInvalidToken       = "The authentication token is invalid."
	MsgTokenExpired       = "The authentication token has expired."
	MsgUserNotFound       = "User account not found."
	MsgUserAlreadyExists  = "User already exists."

	// Pi

	MsgPiLoginFailed      = "Unable to authenticate with Pi."
	MsgPiPaymentFailed    = "Pi payment verification failed."
	MsgPiBalanceLow       = "Insufficient Pi balance."

	// Product

	MsgProductNotFound    = "Requested product not found."

	// Order

	MsgOrderExpired       = "Order has expired."

	// Transaction

	MsgTransactionFailed  = "Transaction failed."

	// VTpass

	MsgVTpassFailed       = "Unable to process utility purchase."
)

// Wallet

const (
	MsgWalletNotFound            = "Wallet not found."
	MsgWalletAlreadyExists       = "Wallet already exists."
	MsgWalletInsufficientBalance = "Insufficient wallet balance."
	MsgWalletInvalidCurrency     = "Invalid wallet currency."
)