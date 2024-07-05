package services

import (
	"eazyWallet/data/models"
	"eazyWallet/data/repositories"
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
)

type WalletService interface {
	CreateWalletAccount(request.CreateWalletAccount) (*response.CreateWalletAccountResponse, error)
	GetWalletAccountById(id uint64) (*models.Account, error)
}

type WalletServiceImpl struct {
	repository repositories.WalletRepository
}

func NewWalletServiceImpl() WalletService {
	return &WalletServiceImpl{repository: repositories.NewWalletRepository()}
}

func (receiver *WalletServiceImpl) CreateWalletAccount(account request.CreateWalletAccount) (*response.CreateWalletAccountResponse, error) {
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

func MapCreateWalletResponse(wallet *models.Account) *response.CreateWalletAccountResponse {
	return &response.CreateWalletAccountResponse{
		WalletId:      wallet.ID,
		AccountNumber: wallet.AccountNumber,
	}
}

func MapRequestToWallet(account request.CreateWalletAccount) *models.Account {
	return &models.Account{
		AccountNumber: account.AccountNumber,
		Password:      account.Password,
		UserId:        account.UserId,
	}
}

func MapWalletAccountRequest(user *models.User, accountNumber string, password string) request.CreateWalletAccount {
	return request.CreateWalletAccount{
		AccountNumber: accountNumber,
		UserId:        user.ID,
		Password:      password,
	}
}
