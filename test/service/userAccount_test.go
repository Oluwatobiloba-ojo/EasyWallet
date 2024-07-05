package service

import (
	"eazyWallet/dto/request"
	response2 "eazyWallet/dto/response"
	"eazyWallet/services"
	"eazyWallet/util/config"
	"eazyWallet/util/message"
	"errors"
	"testing"
)

func TestThatUserCanNotCreateAnAccountWithAnInvalidPhoneNumber(t *testing.T) {
	var user request.CreateUserRequest
	var user_service services.UserService
	user_service = services.NewUserService()
	user = *request.NewCreateUserRequest("firstName", "password", "08129810794", "lastName", "ojot630@gmail.com")
	_, err := user_service.CreateAccount(&user)
	if err != nil {
		t.Errorf("Actuals %v\n  Expected %v\n", errors.New("Invalid phone number"), err)
	}
}

func TestThatUserCanCreateAnAccountWithAnValidPhoneNumber(t *testing.T) {
	config.Load("../../.env")
	var user request.CreateUserRequest
	var user_service services.UserService
	user_service = services.NewUserService()
	user = *request.NewCreateUserRequest("firstName", "password", "08129810794", "lastName", "ojot630@gmail.com")
	response, _ := user_service.CreateAccount(&user)
	var mock_response response2.CreateUserResponse
	if response == nil {
		t.Errorf("Actual %v\n   Expected %v\n", nil, mock_response)
	}
}

func TestThatUserCanNotCreateAnAccountWhichAlreadyExist(t *testing.T) {
	config.Load("../../.env")
	var user request.CreateUserRequest
	var user_service services.UserService
	user_service = services.NewUserService()
	user = *request.NewCreateUserRequest("firstName", "password", "08129810594", "lastName", "ojot630@gmail.com")
	response, _ := user_service.CreateAccount(&user)
	var mock_response response2.CreateUserResponse
	if response == nil {
		t.Errorf("Actual %v\n   Expected %v\n", nil, mock_response)
	}
	response, err := user_service.CreateAccount(&user)
	if err == nil {
		t.Errorf("Actual %v\n Expected %v", message.UserAlreadyExist(), err)
	}
}
