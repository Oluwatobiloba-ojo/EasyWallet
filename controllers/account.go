package controllers

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	walletService services.WalletService
}

func NewAccountController() *AccountController {
	return &AccountController{walletService: services.NewWalletServiceImpl()}
}

func (c *AccountController) PerformTransaction(context *gin.Context) {
	var req = &request.PerformTransactionRequest{}
	err := context.BindJSON(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewApiResponse[string](err.Error(), false))
		return
	}
	res, err := c.walletService.PerformTransaction(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, response.NewApiResponse[string](err.Error(), false))
		return
	}
	context.JSON(http.StatusOK, response.NewApiResponse[*response.PerformTransactionResponse](res, true))
}
