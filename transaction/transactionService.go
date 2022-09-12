package transaction

import (
	"kitabisa/logger"
)

type transactionService struct {
	transactionRepository transactionRepository
	midtransService       MidtransService
}

type TransactionService interface {
	GetByCampaignID(GetCampaignTransactionsInput) ([]Transaction, error)
	GetByUserTransactions(int) ([]Transaction, error)
	CreateTransaction(CreateTransactionInput) (Transaction, error)
}

func NewTransactionService(tr transactionRepository, midtransService MidtransService) *transactionService {
	return &transactionService{tr, midtransService}
}

func (t *transactionService) GetByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	transactions, err := t.transactionRepository.GetByCampaignID(input.ID)
	if err != nil {
		logger.Error("err" + err.Error())
		return transactions, err
	}
	return transactions, nil
}

func (t *transactionService) GetByUserTransactions(userID int) ([]Transaction, error) {
	transactions, err := t.transactionRepository.GetUserTransactions(userID)
	if err != nil {
		logger.Error("Error" + err.Error())
		return transactions, err
	}
	return transactions, nil
}

func (t *transactionService) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignId = input.CampaignId
	transaction.Amount = input.Amount
	transaction.UserId = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := t.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		logger.Error("error" + err.Error())
		return newTransaction, err
	}

	paymentURL, err := t.midtransService.GetPaymentURL(newTransaction, input.User)
	if err != nil {
		logger.Error("error" + err.Error())
		return newTransaction, err
	}
	newTransaction.PaymentURL = paymentURL

	newTransaction, err = t.transactionRepository.UpdateTransaction(newTransaction)
	if err != nil {
		logger.Error("error" + err.Error())
		return newTransaction, err
	}

	return newTransaction, nil
}
