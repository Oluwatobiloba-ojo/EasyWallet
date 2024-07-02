package request

type InitiateTransactionRequest struct {
	Email          string
	Amount         float64
	PaymentMeans   string
	CurrencyChange string
}

func NewInitiateTransactionRequest(email string, amount float64, paymentMeans string, currencyChange string) *InitiateTransactionRequest {
	return &InitiateTransactionRequest{Email: email, Amount: amount, PaymentMeans: paymentMeans, CurrencyChange: currencyChange}
}
