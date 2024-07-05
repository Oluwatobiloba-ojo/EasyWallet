package services

import (
	"eazyWallet/dto/request"
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
	if err != nil {
		return nil, err
	}
	return mapPaystackToResponse(payStackResponse)
}

func mapPaystackToResponse(stackResponse *response.PaystackTransactionResponse) (*response.InitiateTransactionResponse, error) {
	return response.NewInitiateTransactionResponse(stackResponse.Data.Authorization_Url, stackResponse.Data.Refrence), nil
}

func createPayStackRequest(request *request.InitiateTransactionRequest) map[string]any {
	if request.CurrencyChange == constant.NAIRA {
		request.Amount *= 100
	}
	return map[string]any{
		"email":     request.Email,
		"amount":    request.Amount,
		"currency":  request.CurrencyChange,
		"reference": request.RefrenceCode,
	}
}
