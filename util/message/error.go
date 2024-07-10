package message

import "errors"

var (
	PAYMENT_MEANS_ERROR        = "Invalid payment means provided"
	CURRENCY_MEANS_ERROR       = "Invalid currency exchange"
	PAYMENT_TRANSACTION_FAILED = "Transaction failed try again later"
	INVALID_REQUEST            = "Invalid request body"
	USER_ALREADY_EXIST         = "User already exist"
	WALLET_DOESNT_EXIST        = "Wallet does not exist"
	AUTHORIZATION_WRONG        = "Authorization went wrong, credential invalid"
)

func PaymentTransactionFailed() error {
	return errors.New(PAYMENT_TRANSACTION_FAILED)
}

func InvalidRequestObject() error {
	return errors.New(INVALID_REQUEST)
}

func UserAlreadyExist() error {
	return errors.New(USER_ALREADY_EXIST)
}

func WalletDoesNotExist() error {
	return errors.New(WALLET_DOESNT_EXIST)
}
