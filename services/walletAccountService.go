package services

import (
	"eazyWallet/data/models"
	"eazyWallet/data/repositories"
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/logger"
	"eazyWallet/util/constant"
	"eazyWallet/util/message"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type WalletService interface {
	CreateWalletAccount(*request.CreateWalletAccount) (*response.CreateWalletAccountResponse, error)
	GetWalletAccountById(id uint64) (*models.Account, error)
	PerformTransaction(req *request.PerformTransactionRequest) (*response.PerformTransactionResponse, error)
	GetTransactionBelongingTo(accountNumber string) ([]response.TransactionResponse, error)
	FundWallet(transactionId string, status string)
}

type WalletServiceImpl struct {
	repository         repositories.WalletRepository
	paymentService     *PaymentServiceImpl
	transactionService TransactionService
	userService        UserService
}

func NewWalletServiceImpl() WalletService {
	walletService := &WalletServiceImpl{
		repository:     repositories.NewWalletRepository(),
		paymentService: NewPaymentServiceImpl(),
	}
	walletService.transactionService = NewTransactionServiceImpl(walletService)
	walletService.userService = NewUserService(walletService)
	return walletService
}

func (receiver *WalletServiceImpl) CreateWalletAccount(account *request.CreateWalletAccount) (*response.CreateWalletAccountResponse, error) {
	var wallet *models.Account
	wallet = MapRequestToWallet(account)
	wallet, err := receiver.repository.Save(wallet)
	if err != nil {
		return nil, err
	}
	return MapCreateWalletResponse(wallet), nil
}

func (receiver *WalletServiceImpl) GetWalletAccountById(id uint64) (*models.Account, error) {
	var wallet *models.Account
	wallet, err := receiver.repository.FindById(&id)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (receiver *WalletServiceImpl) PerformTransaction(transactionRequest *request.PerformTransactionRequest) (transactionResponse *response.PerformTransactionResponse, err error) {
	wallet, err := receiver.getWalletAccountByAccountNumber(transactionRequest.AccountNumber)
	if err != nil {
		return nil, err
	}
	user, err := receiver.userService.GetUserById(wallet.UserId)
	if err != nil {
		return nil, err
	}
	transaction, err := receiver.transactionService.CreateTransaction(mapToCreateTransactionRequest(transactionRequest, wallet.ID))
	if err != nil {
		return nil, err
	}
	initiateTransaction, err := receiver.paymentService.InitiateTransaction(mapToInitiateTransaction(transactionRequest, transaction.ID, user.Email))
	if err != nil {
		return nil, err
	}
	res := &response.PerformTransactionResponse{
		Url:     initiateTransaction.Url,
		Message: message.PERFORM_TRANSACTION}
	return res, nil
}

func (receiver *WalletServiceImpl) FundWallet(transactionId string, status string) {
	fmt.Println("It came into fund wallet ", status, "Transaction ", transactionId)
	if status == constant.PAYSTACK_SUCCESS || status == constant.MONNIFY_SUCCESS {
		transaction, err := receiver.transactionService.UpdateTransaction(transactionId, constant.SUCCESS)
		if err != nil {
			logger.ErrorLogger(err)
		}
		log.Println("Transaction ", transaction)
		wallet, err := receiver.GetWalletAccountById(transaction.AccountId)
		if err != nil {
			logger.ErrorLogger(err)
		}
		log.Println("Wallet ", wallet)
		wallet.AccountBalance += transaction.Amount
		wallet, err = receiver.repository.Save(wallet)
		if err != nil {
			logger.ErrorLogger(err)
		}
		return
	}
	transaction, err := receiver.transactionService.UpdateTransaction(transactionId, constant.FAILED)
	if err != nil {
		logger.ErrorLogger(err)
	}
	log.Println(transaction)
}

func mapToInitiateTransaction(transactionRequest *request.PerformTransactionRequest, refrenceCode uuid.UUID, email string) *request.InitiateTransactionRequest {
	return request.NewInitiateTransactionRequest(email, transactionRequest.Amount, transactionRequest.PaymentMeans, transactionRequest.CurrencyChange, refrenceCode)
}

func mapToCreateTransactionRequest(transactionRequest *request.PerformTransactionRequest, walletId uint64) *request.CreateTransactionRequest {
	return request.NewCreateTransactionRequest(transactionRequest.Amount, walletId, transactionRequest.Description, transactionRequest.Recipient_Name)
}

func (receiver *WalletServiceImpl) GetTransactionBelongingTo(accountNumber string) ([]response.TransactionResponse, error) {
	wallet, err := receiver.getWalletAccountByAccountNumber(accountNumber)
	if err != nil {
		return nil, err
	}
	transactions, err := receiver.transactionService.GetTransactionsByAccountId(wallet.ID)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (receiver *WalletServiceImpl) getWalletAccountByAccountNumber(number string) (*models.Account, error) {
	wallet, err := receiver.repository.FindWalletByAccountNumber(number)
	if err != nil {
		return nil, message.WalletDoesNotExist()
	}
	return wallet, nil
}

func MapCreateWalletResponse(wallet *models.Account) *response.CreateWalletAccountResponse {
	return &response.CreateWalletAccountResponse{
		WalletId:      wallet.ID,
		AccountNumber: wallet.AccountNumber,
	}
}

func MapRequestToWallet(account *request.CreateWalletAccount) *models.Account {
	return &models.Account{
		AccountNumber: account.AccountNumber,
		Password:      account.Password,
		UserId:        account.UserId,
	}
}

func MapWalletAccountRequest(user *models.User, accountNumber string, password string) *request.CreateWalletAccount {
	return &request.CreateWalletAccount{
		AccountNumber: accountNumber,
		UserId:        user.ID,
		Password:      password,
	}
}
