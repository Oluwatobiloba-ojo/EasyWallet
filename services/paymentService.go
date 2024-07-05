package services

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/util/constant"
	"eazyWallet/util/message"
	"fmt"
)

type PaymentService interface {
	FundWallet(req map[string]any) (*response.InitiateTransactionResponse, error)
}

type PaymentServiceImpl struct {
}

func NewPaymentServiceImpl() *PaymentServiceImpl {
	return &PaymentServiceImpl{}
}

func (service *PaymentServiceImpl) InitiateTransaction(req *request.InitiateTransactionRequest) (*response.InitiateTransactionResponse, error) {
	if req.PaymentMeans != constant.MONNIFY && req.PaymentMeans != constant.Paystack {
		return nil, fmt.Errorf(message.PAYMENT_MEANS_ERROR)
	}
	if req.CurrencyChange != constant.NAIRA && req.CurrencyChange != constant.USA {
		return nil, fmt.Errorf(message.CURRENCY_MEANS_ERROR)
	}
	if req.PaymentMeans == constant.Paystack {
		paystackRequest := createPayStackRequest(req.Email, req.Amount, req.CurrencyChange)
		return NewPaystackService().FundWallet(paystackRequest)
	} else {
		monifyRequest := NewMonifyService().createMonnifyRequest(req.Email, req.Amount, req.CurrencyChange)
		return NewMonifyService().FundWallet(monifyRequest)
	}
}
