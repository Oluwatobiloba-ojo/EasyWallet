package services

import (
	"eazyWallet/data/models"
	"eazyWallet/data/repositories"
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/util/message"
	"github.com/google/uuid"
	"log"
)

type TransactionService interface {
	CreateTransaction(request *request.CreateTransactionRequest) (*models.Transaction, error)
	GetTransactionsByAccountId(id uint64) ([]response.TransactionResponse, error)
	UpdateTransaction(id string, status string) (*models.Transaction, error)
}

type TransactionServiceImpl struct {
	repository    repositories.TransactionRepository
	walletService WalletService
}

func NewTransactionServiceImpl(wallet WalletService) TransactionService {
	return &TransactionServiceImpl{
		repository:    repositories.NewTransactionRepositoryImpl(),
		walletService: wallet,
	}
}

func (t *TransactionServiceImpl) UpdateTransaction(id string, status string) (*models.Transaction, error) {
	transactionId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	transaction, err := t.getTransactionById(transactionId)
	if err != nil {
		return nil, err
	}
	log.Println("Update ", transaction)
	transaction.Status = status
	transaction, err = t.repository.Save(transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (receiver *TransactionServiceImpl) CreateTransaction(transactionRequest *request.CreateTransactionRequest) (*models.Transaction, error) {
	_, err := receiver.walletService.GetWalletAccountById(transactionRequest.WalletId)
	if err != nil {
		return nil, message.WalletDoesNotExist()
	}
	newTransaction := mapRequestToTransaction(transactionRequest)
	newTransaction, err = receiver.repository.Save(newTransaction)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return newTransaction, nil
}

func (receiver *TransactionServiceImpl) getTransactionById(id uuid.UUID) (*models.Transaction, error) {
	transaction, err := receiver.repository.FindById(&id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (receiver *TransactionServiceImpl) GetTransactionsByAccountId(id uint64) ([]response.TransactionResponse, error) {
	transactions, err := receiver.repository.FindAllTransactionByAccountId(id)
	if err != nil {
		return nil, err
	}
	return mapToTransactionResponse(transactions), nil
}

func mapToTransactionResponse(transactions []*models.Transaction) []response.TransactionResponse {
	var responses []response.TransactionResponse
	for _, transaction := range transactions {
		transaction := response.TransactionResponse{
			Amount: transaction.Amount, AccountId: transaction.AccountId,
			Description: transaction.Description, RecipientName: transaction.RecipientName,
			Status: transaction.Status}
		responses = append(responses, transaction)
	}
	return responses
}

func mapRequestToTransaction(transactionRequest *request.CreateTransactionRequest) *models.Transaction {
	return &models.Transaction{
		RecipientName: transactionRequest.RecipientName,
		Description:   transactionRequest.Description,
		Amount:        transactionRequest.Amount,
		AccountId:     transactionRequest.WalletId,
	}
}
