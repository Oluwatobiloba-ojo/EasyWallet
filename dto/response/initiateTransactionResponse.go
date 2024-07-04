package response

type InitiateTransactionResponse struct {
	Url      string `json:"url"`
	Refrence string `json:"refrence"`
}

func NewInitiateTransactionResponse(url string, refrence string) *InitiateTransactionResponse {
	return &InitiateTransactionResponse{Url: url, Refrence: refrence}
}
