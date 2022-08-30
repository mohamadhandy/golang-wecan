package campaign

import (
	"kitabisa/logger"

	"gorm.io/gorm"
)

type CampaignRepositoryDB interface {
	FindAllCampaign() ([]Campaign, error)
	FindCampaignById(int) (Campaign, error)
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

func (c *campaignRepositoryDB) FindCampaignById(campaignId int) (Campaign, error) {
	var err error
	var campaign Campaign
	if err = c.db.Where("campaign_id = ?", campaignId).Find(campaign).Error; err != nil {
		logger.Error("Error" + err.Error())
		return campaign, err
	}
	return campaign, nil
}
