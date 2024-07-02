package response

type InitiateTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Authorization_Url string `json:"authorization_url"`
	Access_Code       string `json:"access_code"`
	Refrence          string `json:"reference"`
}
