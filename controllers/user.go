package controllers

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController() *UserController {
	walletService := services.NewWalletServiceImpl()
	return &UserController{userService: services.NewUserService(walletService)}
}

func (c *UserController) CreateAccount(context *gin.Context) {
	createAccountRequest := &request.CreateUserRequest{}
	err := context.BindJSON(createAccountRequest)
	log.Println("Account Request ", createAccountRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewApiResponse[string](err.Error(), false))
		return
	}
	createAccountResponse, err := c.userService.CreateAccount(createAccountRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewApiResponse[string](err.Error(), false))
		return
	}
	log.Println("user account successfully created: ", createAccountResponse)
	context.JSON(http.StatusCreated, response.NewApiResponse[*response.CreateUserResponse](createAccountResponse, true))
}
