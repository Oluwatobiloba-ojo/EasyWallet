package request

type CreateTransactionRequest struct {
	Amount        float64 `json:"amount"`
	WalletId      uint64  `json:"wallet_id"`
	Description   string  `json:"description"`
	RecipientName string  `json:"recipient_name"`
}

func NewCreateTransactionRequest(amount float64, walletId uint64, description string, recipientName string) *CreateTransactionRequest {
	return &CreateTransactionRequest{Amount: amount, WalletId: walletId, Description: description, RecipientName: recipientName}
}
