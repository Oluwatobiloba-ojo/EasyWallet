package response

type MonnifyTransactionResponse struct {
	RequestSuccessful bool   `json:"requestSuccessful"`
	ResponseMessage   string `json:"responseMessage"`
	ResponseCode      string `json:"responseCode"`
	ResponseBody      struct {
		TransactionReference string   `json:"transactionReference"`
		PaymentReference     string   `json:"paymentReference"`
		MerchantName         string   `json:"merchantName"`
		ApiKey               string   `json:"apiKey"`
		EnabledPaymentMethod []string `json:"enabledPaymentMethod"`
		CheckoutUrl          string   `json:"checkoutUrl"`
	} `json:"responseBody"`
}
