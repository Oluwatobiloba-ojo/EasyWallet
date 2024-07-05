package service

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/services"
	"eazyWallet/util/config"
	"fmt"
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestThatPaymentTransactionWithOutTheCorrectPaymentMeansReturnError(t *testing.T) {
	var req *request.InitiateTransactionRequest
	paymentServiceImpl := services.NewPaymentServiceImpl()
	req = request.NewInitiateTransactionRequest("olawale@gmail.com", 1000, "WrongPaymentMeans", "NGN", uuid.New())
	_, err := paymentServiceImpl.InitiateTransaction(req)
	mockerror := fmt.Errorf("Invalid payment means provided")
	if err == nil {
		t.Errorf("Actual %v\n Expected %v\n", mockerror, err)
	}
}

func TestThatPaymentTransactionFailsWhenTheCurrencyChangeIsWrong(t *testing.T) {
	var req *request.InitiateTransactionRequest
	paymentServiceImpl := services.NewPaymentServiceImpl()
	req = request.NewInitiateTransactionRequest("olawale@gmail.com", 1000, "PAYSTACK", "WrongCurrencyChange", uuid.New())
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
	req = request.NewInitiateTransactionRequest("olawale@gmail.com", 1000, "PAYSTACK", "NGN", uuid.New())
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

func TestThatPaymentTransactionOfMonnifyServiceReturnAnObjectAfterSuccessfullyInitiateTransaction(t *testing.T) {
	config.Load("../../.env")
	var req *request.InitiateTransactionRequest
	paymentServiceImpl := services.NewPaymentServiceImpl()
	req = request.NewInitiateTransactionRequest("ojot630@gmail.com", 1000, "MONNIFY", "NGN", uuid.New())
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
