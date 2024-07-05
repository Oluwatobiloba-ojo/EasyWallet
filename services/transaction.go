package services

import (
	"eazyWallet/data/models"
	"eazyWallet/data/repositories"
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/util/message"
)

type TransactionService interface {
	CreateTransaction(request *request.CreateTransactionRequest) (*response.CreateTransactionResponse, error)
}

type TransactionServiceImpl struct {
	repository    repositories.TransactionRepository
	walletService WalletService
}

func NewTransactionServiceImpl() TransactionService {
	return &TransactionServiceImpl{
		repository:    repositories.NewTransactionRepositoryImpl(),
		walletService: NewWalletServiceImpl(),
	}
}

func (receiver *TransactionServiceImpl) CreateTransaction(transactionRequest *request.CreateTransactionRequest) (*response.CreateTransactionResponse, error) {
	_, err := receiver.walletService.GetWalletAccountById(transactionRequest.WalletId)
	if err != nil {
		return nil, message.WalletDoesNotExist()
	}
	newTransaction := mapRequestToTransaction(transactionRequest)
	newTransaction, err = receiver.repository.Save(newTransaction)
	if err != nil {
		return nil, err
	}
	return mapToCreateTransactionResponse(newTransaction), nil
}

func mapToCreateTransactionResponse(transaction *models.Transaction) *response.CreateTransactionResponse {
	return &response.CreateTransactionResponse{
		TransactionId: transaction.ID,
		Message:       message.TRANSACTION_CREATED,
	}
}

func mapRequestToTransaction(transactionRequest *request.CreateTransactionRequest) *models.Transaction {
	return &models.Transaction{
		RecipientName: transactionRequest.RecipientName,
		Description:   transactionRequest.Description,
		Amount:        transactionRequest.Amount,
		AccountId:     transactionRequest.WalletId,
	}
}
