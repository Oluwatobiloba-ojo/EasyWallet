package service

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/services"
	"eazyWallet/util/config"
	"eazyWallet/util/message"
	"log"
	"testing"
)

func TestTransactionCanBeCreatedWithACorrectWalletId(t *testing.T) {
	config.Load("../../.env")
	var req *request.CreateTransactionRequest
	var service services.TransactionService
	req = request.NewCreateTransactionRequest(1500, 1, "Shopping and marketing", "ope")
	service = services.NewTransactionServiceImpl()
	res, err := service.CreateTransaction(req)
	mockResponse := response.CreateTransactionResponse{}
	if err != nil {
		t.Errorf("Actual %v\n Expected %v", err, nil)
	}
	if res == nil {
		t.Errorf("Actuals %v\n Expected %v ", res, mockResponse)
	}
	log.Println(res)
}
func TestTransactionWillNotBeCreatedWithAIncorrectWalletId(t *testing.T) {
	config.Load("../../.env")
	var req *request.CreateTransactionRequest
	var service services.TransactionService
	req = request.NewCreateTransactionRequest(1000, 10, "Shopping", "ope")
	service = services.NewTransactionServiceImpl()
	_, err := service.CreateTransaction(req)
	if err == nil {
		t.Errorf("Actual %v\n Expected %v", err, message.WalletDoesNotExist())
	}
	log.Println(err)
}
