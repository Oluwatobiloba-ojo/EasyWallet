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

var userController = controllers.NewUserController()

func TestCreateAccountEndPoint(t *testing.T) {
	config.Load("../../.env")
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		router := gin.Default()
		router.POST("/api/v1/user", userController.CreateAccount)
		writer := httptest.NewRecorder()
		userRequest := request.NewCreateUserRequest("Olawale", "1234", "08032389557", "Tunji", "ooluwatobi825@gmail.com")
		data, _ := json.Marshal(userRequest)
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/user", bytes.NewReader(data))
		request.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(writer, request)
		log.Println(writer.Body.String())
		assert.Equal(t, http.StatusCreated, writer.Code)
		assert.NotNil(t, writer.Body)
	})
}
