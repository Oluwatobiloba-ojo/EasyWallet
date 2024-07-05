package service

import (
	"eazyWallet/dto/request"
	"eazyWallet/services"
	"eazyWallet/util/config"
	"testing"
)

func TestThatWhenATransactionInitiatedUntoAWalletATransactionIsCreatedButPending(t *testing.T) {
	config.Load("../../.env")
	var req *request.PerformTransactionRequest
	var walletservice services.WalletService
	req = request.NewPerformTransactionRequest("08129810794", "School fees", "paystack", 1000, "ola", "NGN")
	walletservice = services.NewWalletServiceImpl()
	_, err := walletservice.PerformTransaction(req)
	if err != nil {
		t.Errorf("Actual %v\n  Expected %v", err, nil)
	}
	transactions, err := walletservice.GetTransactionBelongingTo(req.AccountNumber)
	if err != nil {
		t.Errorf("Actual %v\n  Expected %v", err, nil)
	}
	if len(transactions) == 0 {
		t.Errorf("Actual %v\n  Expected %v", len(transactions), 1)
	}
}
