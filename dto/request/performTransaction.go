package request

type PerformTransactionRequest struct {
	AccountNumber  string  `json:"account_number"`
	Description    string  `json:"description"`
	PaymentMeans   string  `json:"payment_means"`
	Amount         float64 `json:"amount"`
	CurrencyChange string  `json:"currency_change"`
	Recipient_Name string  `json:"recipient_name"`
}

func NewPerformTransactionRequest(accountNumber string, description string, paymentMeans string, amount float64, recipient_Name string, currency_change string) *PerformTransactionRequest {
	return &PerformTransactionRequest{AccountNumber: accountNumber, Description: description, PaymentMeans: paymentMeans, Amount: amount, Recipient_Name: recipient_Name, CurrencyChange: currency_change}
}
