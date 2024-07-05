package response

import "github.com/google/uuid"

type CreateTransactionResponse struct {
	Message       string    `json:"message"`
	TransactionId uuid.UUID `json:"transaction_Id"`
}
