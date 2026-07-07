package errors

// Generic
const (
	CodeBadRequest          = "BAD_REQUEST"
	CodeUnauthorized        = "UNAUTHORIZED"
	CodeForbidden           = "FORBIDDEN"
	CodeNotFound            = "NOT_FOUND"
	CodeConflict            = "CONFLICT"
	CodeValidationFailed    = "VALIDATION_FAILED"
	CodeInternalServerError = "INTERNAL_SERVER_ERROR"
	CodeEndpointNotFound    = "ENDPOINT_NOT_FOUND"
)


// Authentication
const (
	CodeAuthInvalidCredentials = "AUTH_INVALID_CREDENTIALS"
	CodeAuthInvalidToken       = "AUTH_INVALID_TOKEN"
	CodeAuthTokenExpired       = "AUTH_TOKEN_EXPIRED"
	CodeAuthUserNotFound       = "AUTH_USER_NOT_FOUND"
	CodeAuthUserExists         = "AUTH_USER_ALREADY_EXISTS"
	CodeAuthRefreshExpired     = "AUTH_REFRESH_TOKEN_EXPIRED"
)


// Pi Network
const (
	CodePiLoginFailed          = "PI_LOGIN_FAILED"
	CodePiUserVerification     = "PI_USER_VERIFICATION_FAILED"
	CodePiPaymentVerification  = "PI_PAYMENT_VERIFICATION_FAILED"
	CodePiInsufficientBalance  = "PI_INSUFFICIENT_BALANCE"
)


// Products
const (
	CodeProductNotFound = "PRODUCT_NOT_FOUND"
	CodeProductInactive = "PRODUCT_INACTIVE"
)


// Orders
const (
	CodeOrderNotFound         = "ORDER_NOT_FOUND"
	CodeOrderAlreadyCompleted = "ORDER_ALREADY_COMPLETED"
	CodeOrderExpired          = "ORDER_EXPIRED"
)


// Transactions
const (
	CodeTransactionFailed    = "TRANSACTION_FAILED"
	CodeTransactionNotFound  = "TRANSACTION_NOT_FOUND"
	CodeTransactionDuplicate = "TRANSACTION_DUPLICATE"
)


// VTpass
const (
	CodeVTpassRequestFailed = "VTPASS_REQUEST_FAILED"
	CodeVTpassTimeout       = "VTPASS_TIMEOUT"
	CodeVTpassUnavailable   = "VTPASS_UNAVAILABLE"
)

// Wallet

const (
	CodeWalletNotFound            = "WALLET_NOT_FOUND"
	CodeWalletAlreadyExists       = "WALLET_ALREADY_EXISTS"
	CodeWalletInsufficientBalance = "WALLET_INSUFFICIENT_BALANCE"
	CodeWalletInvalidCurrency     = "WALLET_INVALID_CURRENCY"
)