package request

import "github.com/google/uuid"

type InitiateTransactionRequest struct {
	Email          string
	Amount         float64
	PaymentMeans   string
	CurrencyChange string
	RefrenceCode   uuid.UUID
}

func NewInitiateTransactionRequest(email string, amount float64, paymentMeans string, currencyChange string, refrenceCode uuid.UUID) *InitiateTransactionRequest {
	return &InitiateTransactionRequest{Email: email, Amount: amount, PaymentMeans: paymentMeans, CurrencyChange: currencyChange, RefrenceCode: refrenceCode}
}
