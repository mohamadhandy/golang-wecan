package transaction

import (
	"kitabisa/logger"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

type TransactionRepositoryDB interface {
	GetByCampaignID(int) ([]Transaction, error)
	GetUserTransactions(int) ([]Transaction, error)
	CreateTransaction(Transaction) ([]Transaction, error)
}

func NewTransactionRepositoryDB(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (t *transactionRepository) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var err error
	var transactions []Transaction
	if err = t.db.Where("campaign_id = ?", campaignID).Find(&transactions).Error; err != nil {
		logger.Error("error " + err.Error())
		return transactions, err
	}
	return transactions, nil
}

func (t *transactionRepository) GetUserTransactions(userID int) ([]Transaction, error) {
	var err error
	var transactions []Transaction
	if err = t.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		logger.Error("error" + err.Error())
		return transactions, err
	}
	return transactions, nil
}

func (t *transactionRepository) CreateTransaction(transaction Transaction) (Transaction, error) {
	err := t.db.Create(&transaction).Error
	if err != nil {
		logger.Error("Error create transaction" + err.Error())
		return transaction, err
	}
	return transaction, nil
}
