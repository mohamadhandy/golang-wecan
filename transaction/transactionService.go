package transaction

import (
	"kitabisa/logger"
)

type transactionService struct {
	transactionRepository transactionRepository
}

type TransactionService interface {
	GetByCampaignID(GetCampaignTransactionsInput) ([]Transaction, error)
	GetByUserTransactions(int) ([]Transaction, error)
}

func NewTransactionService(tr transactionRepository) *transactionService {
	return &transactionService{transactionRepository: tr}
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
