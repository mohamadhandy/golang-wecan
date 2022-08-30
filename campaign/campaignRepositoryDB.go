package campaign

import (
	"kitabisa/logger"

	"gorm.io/gorm"
)

type CampaignRepositoryDB interface {
	FindAllCampaign() ([]Campaign, error)
}

type campaignRepositoryDB struct {
	db *gorm.DB
}

func NewCampaignRepositoryDB(db *gorm.DB) *campaignRepositoryDB {
	return &campaignRepositoryDB{db}
}

func (c *campaignRepositoryDB) FindAllCampaign() ([]Campaign, error) {
	var err error
	var campaigns []Campaign
	if err = c.db.Find(&campaigns).Error; err != nil {
		logger.Error("Error" + err.Error())
		return campaigns, err
	}
	return campaigns, nil
}
