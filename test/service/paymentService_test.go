package service

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/services"
	"eazyWallet/util/config"
	"fmt"
	"log"
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

func TestThatPaymentTransactionReturnAnObjectAfterSuccessfullyInitiateTransaction(t *testing.T) {
	config.Load("../../.env")
	var req *request.InitiateTransactionRequest
	paymentServiceImpl := services.NewPaymentServiceImpl()
	req = request.NewInitiateTransactionRequest("olawale@gmail.com", 1000, "PAYSTACK", "NGN")
	transaction, err := paymentServiceImpl.InitiateTransaction(req)
	if err != nil {
		t.Errorf("Actual %v\n Expected %v\n", nil, err)
	}
	actual := response.InitiateTransactionResponse{}
	if transaction == nil {
		t.Errorf("Actual %v\n Expected %v\n", actual, transaction)
	}
	log.Println(transaction)
}
