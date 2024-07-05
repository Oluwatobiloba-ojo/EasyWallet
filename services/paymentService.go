package services

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/util/constant"
	"eazyWallet/util/message"
	"fmt"
	"strings"
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
	req.PaymentMeans = strings.ToUpper(req.PaymentMeans)
	req.CurrencyChange = strings.ToUpper(req.CurrencyChange)
	if req.PaymentMeans != constant.MONNIFY && req.PaymentMeans != constant.Paystack {
		return nil, fmt.Errorf(message.PAYMENT_MEANS_ERROR)
	}
	if strings.ToUpper(req.CurrencyChange) != constant.NAIRA && req.CurrencyChange != constant.USA {
		return nil, fmt.Errorf(message.CURRENCY_MEANS_ERROR)
	}
	if req.PaymentMeans == constant.Paystack {
		paystackRequest := createPayStackRequest(req)
		return NewPaystackService().FundWallet(paystackRequest)
	} else {
		monifyRequest := NewMonifyService().createMonnifyRequest(req)
		return NewMonifyService().FundWallet(monifyRequest)
	}
}
