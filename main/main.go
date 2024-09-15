package main

import (
	"eazyWallet/controllers"
	"eazyWallet/dataBase"
	"eazyWallet/util/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.Load(".env")
	db := dataBase.DBConnection()
	log.Print("connected", db)
	router := gin.Default()
	userController := controllers.NewUserController()
	accountController := controllers.NewAccountController()
	router.POST("/api/v1/user", userController.CreateAccount)
	router.POST("/api/v1/init_transaction", accountController.PerformTransaction)
	router.POST("/api/v1/paystack_webhook", accountController.WebHookPaystackEndPoint)
	router.POST("/api/v1/monnify_webhook", accountController.WebHookMonnifyEndPoint)
	router.GET("/api/v1/transactions", accountController.GetAllTransaction)
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
