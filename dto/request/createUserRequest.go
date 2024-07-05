package request

type CreateUserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"lastName"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func NewCreateUserRequest(firstName string, password string, phoneNumber string, last_name string, email string) *CreateUserRequest {
	return &CreateUserRequest{FirstName: firstName, Password: password, PhoneNumber: phoneNumber, LastName: last_name, Email: email}
}
