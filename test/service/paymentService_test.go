package service

import (
	"eazyWallet/dto/request"
	"eazyWallet/services"
	"fmt"
	"testing"
)

func TestThatPaymentTransactionWithOutTheCorrectPaymentMeansReturnError(t *testing.T) {
	var req *request.InitiateTransactionRequest
	paymentServiceImpl := services.NewPaymentServiceImpl()
	req = request.NewInitiateTransactionRequest("olawale@gmail.com", 1000, "WrongPaymentMeans", "NGN")
	_, err := paymentServiceImpl.InitiateTransaction(req)
	mockerror := fmt.Errorf("Invalid payment means provided")
	if err == nil {
		t.Errorf("Actual %v\n Expected %v\n", mockerror, err)
	}
}

func TestThatPaymentTransactionFailsWhenTheCurrencyChangeIsWrong(t *testing.T) {
	var req *request.InitiateTransactionRequest
	paymentServiceImpl := services.NewPaymentServiceImpl()
	req = request.NewInitiateTransactionRequest("olawale@gmail.com", 1000, "PAYSTACK", "WrongCurrencyChange")
	_, err := paymentServiceImpl.InitiateTransaction(req)
	mockerror := fmt.Errorf("Invalid currency change")
	if err == nil {
		t.Errorf("Actual %v\n Expected %v\n", mockerror, err)
	}
}
