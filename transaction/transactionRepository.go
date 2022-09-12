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
