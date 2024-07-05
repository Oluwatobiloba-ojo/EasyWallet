package controller

import (
	"bytes"
	"eazyWallet/controllers"
	"eazyWallet/dto/request"
	"eazyWallet/util/config"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var walletController = controllers.NewAccountController()

func TestThatUserCanInitializeATransaction(t *testing.T) {
	config.Load("../../.env")
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		router := gin.Default()
		router.POST("/api/v1/init_transaction", walletController.PerformTransaction)
		writer := httptest.NewRecorder()
		performTransactionRequest := request.NewPerformTransactionRequest("08032389557", "Delivery goods", "Paystack", 10000, "ojot630@gmail.com", "NGN")
		data, _ := json.Marshal(performTransactionRequest)
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/init_transaction", bytes.NewReader(data))
		request.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(writer, request)
		log.Println(writer.Body.String())
		assert.Equal(t, http.StatusOK, writer.Code)
		assert.NotNil(t, writer.Body)
	})
}
