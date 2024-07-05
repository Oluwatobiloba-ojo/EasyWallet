package request

type CreateWalletAccount struct {
	AccountNumber string
	Password      string
	UserId        uint64
}
