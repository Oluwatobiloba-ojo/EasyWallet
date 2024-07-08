package controllers

import (
	"eazyWallet/dto/request"
	"eazyWallet/dto/response"
	"eazyWallet/logger"
	"eazyWallet/services"
	"github.com/gin-gonic/gin"
	"log"
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

func (c *AccountController) WebHookPaystackEndPoint(context *gin.Context) {
	req := &response.PaystackWebhookResponse{}
	err := context.ShouldBindJSON(req)
	if err != nil {
		return
	}
	log.Println(req)
	go func() {
		c.walletService.FundWallet(req.Data.Reference, req.Event)
	}()
	context.JSON(http.StatusOK, response.NewApiResponse[string]("update successfully", true))
}

func (c *AccountController) WebHookMonnifyEndPoint(ctx *gin.Context) {
	req := &response.MonnifyWebhookResponse{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		logger.ErrorLogger(err)
		return
	}
	go func() {
		c.walletService.FundWallet(req.EventData.Reference, req.Event)
	}()
	ctx.JSON(http.StatusOK, response.NewApiResponse[string]("successfully", true))
}
