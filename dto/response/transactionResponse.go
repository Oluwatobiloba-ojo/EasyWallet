package response

type TransactionResponse struct {
	Amount        float64
	AccountId     uint64
	Description   string
	RecipientName string
	Status        string
}
