package services

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/util"
	"eazyWallet/util/config"
	"encoding/json"
)

type MonifyService struct {
}

func NewMonifyService() *MonifyService {
	return &MonifyService{}
}

func (service *MonifyService) FundWallet(req map[string]any) (*response.InitiateTransactionResponse, error) {
	var responses response.MonnifyTransactionResponse
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	base := util.NewEncoding().EncodeTo([]byte(config.MonnifyApiKey + ":" + config.MonnifySecretKey))
	key := "Basic " + base
	Newresponses, err := util.MakePostRequest[response.MonnifyTransactionResponse](key, jsonData, responses, config.MonnifyInitUrl)
	if err != nil {
		return nil, err
	}
	return mapMonnifyToRequest(Newresponses)
}

func mapMonnifyToRequest(newresponses *response.MonnifyTransactionResponse) (*response.InitiateTransactionResponse, error) {
	return response.NewInitiateTransactionResponse(newresponses.ResponseBody.CheckoutUrl, newresponses.ResponseBody.PaymentReference), nil
}

func (service *MonifyService) createMonnifyRequest(request *request.InitiateTransactionRequest) map[string]any {
	return map[string]any{
		"amount":             request.Amount,
		"customerName":       request.Email,
		"customerEmail":      request.Email,
		"paymentReference":   request.RefrenceCode,
		"paymentDescription": "CREDIT",
		"currencyCode":       request.CurrencyChange,
		"contractCode":       config.MonnifyContractCode,
	}
}
