package errorMessage

import "errors"

var (
	PAYMENT_MEANS_ERROR        = "Invalid payment means provided"
	CURRENCY_MEANS_ERROR       = "Invalid currency exchange"
	PAYMENT_TRANSACTION_FAILED = "Transaction failed try again later"
)

func PaymentTransactionFailed() error {
	return errors.New(PAYMENT_TRANSACTION_FAILED)
}
