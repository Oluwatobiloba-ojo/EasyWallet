package services

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/util/constant"
	"eazyWallet/util/errorMessage"
	"fmt"
)

type PaymentService interface {
}

type PaymentServiceImpl struct {
}

func NewPaymentServiceImpl() *PaymentServiceImpl {
	return &PaymentServiceImpl{}
}

func (service *PaymentServiceImpl) InitiateTransaction(req *request.InitiateTransactionRequest) (*response.InitiateTransactionResponse, error) {
	if req.PaymentMeans != constant.MONNIFY && req.PaymentMeans != constant.Paystack {
		return nil, fmt.Errorf(errorMessage.PAYMENT_MEANS_ERROR)
	}
	if req.CurrencyChange != constant.NAIRA && req.CurrencyChange != constant.USA {
		return nil, fmt.Errorf(errorMessage.CURRENCY_MEANS_ERROR)
	}
	return nil, nil
}
