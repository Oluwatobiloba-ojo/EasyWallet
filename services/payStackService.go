package services

import (
	"eazyWallet/dto/response"
	"eazyWallet/util"
	"eazyWallet/util/config"
	"eazyWallet/util/constant"
	"encoding/json"
)

type PaystackService struct {
}

func NewPaystackService() *PaystackService {
	return &PaystackService{}
}

func (service *PaystackService) FundWallet(req map[string]any) (*response.InitiateTransactionResponse, error) {
	var responses response.InitiateTransactionResponse
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	url := "Bearer " + config.PaystackSecretKey
	return util.MakePostRequest[response.InitiateTransactionResponse](url, jsonData, responses)
}

func createPayStackRequest(email string, amount float64, change string) map[string]any {
	if change == constant.NAIRA {
		amount *= 100
	}
	return map[string]any{
		"email":    email,
		"amount":   amount,
		"currency": change,
	}
}
