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
	var responses response.PaystackTransactionResponse
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	key := "Bearer " + config.PaystackSecretKey
	payStackResponse, err := util.MakePostRequest[response.PaystackTransactionResponse](key, jsonData, responses, config.PaystackTransactionUrl)
	return mapPaystackToResponse(payStackResponse)
}

func mapPaystackToResponse(stackResponse *response.PaystackTransactionResponse) (*response.InitiateTransactionResponse, error) {
	return response.NewInitiateTransactionResponse(stackResponse.Data.Authorization_Url, stackResponse.Data.Refrence), nil
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
