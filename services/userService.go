package services

import (
	"eazyWallet/data/models"
	"eazyWallet/data/repositories"
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/util/message"
)

type UserService interface {
	CreateAccount(request *request.CreateUserRequest) (*response.CreateUserResponse, error)
	GetUserById(id uint64) (*models.User, error)
}

type UserServiceImpl struct {
	repositories   repositories.UserRepository
	accountService WalletService
}

func NewUserService(walletService WalletService) UserService {
	return &UserServiceImpl{
		repositories:   repositories.NewUserRepository(),
		accountService: walletService,
	}
}

func (serviceImpl *UserServiceImpl) CreateAccount(userRequest *request.CreateUserRequest) (*response.CreateUserResponse, error) {
	var wallet *response.CreateWalletAccountResponse
	if userRequest == nil {
		return nil, message.InvalidRequestObject()
	}
	user, err := serviceImpl.repositories.GetByEmail(userRequest.Email)
	if err == nil {
		return nil, message.UserAlreadyExist()
	}
	user, err = serviceImpl.repositories.GetByPhoneNumber(userRequest.PhoneNumber)
	if err == nil {
		return nil, message.UserAlreadyExist()
	}
	user = mapToUser(userRequest)
	user, err = serviceImpl.repositories.Save(user)
	if err != nil {
		return nil, err
	}
	wallet, err = serviceImpl.accountService.CreateWalletAccount(MapWalletAccountRequest(user, userRequest.PhoneNumber, userRequest.Password))
	if err != nil {
		return nil, err
	}
	return mapCreateAccountResponse(wallet), nil
}

func (serviceImpl *UserServiceImpl) GetUserById(id uint64) (*models.User, error) {
	user, err := serviceImpl.repositories.FindById(&id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func mapCreateAccountResponse(wallet *response.CreateWalletAccountResponse) *response.CreateUserResponse {
	return &response.CreateUserResponse{
		Account_number: wallet.AccountNumber,
		Message:        message.ACCOUNT_CREATED,
	}
}

func mapToUser(userRequest *request.CreateUserRequest) *models.User {
	return &models.User{
		FirstName:   userRequest.FirstName,
		LastName:    userRequest.LastName,
		PhoneNumber: userRequest.PhoneNumber,
		Email:       userRequest.Email,
	}
}
