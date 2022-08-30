package campaign

import (
	"errors"
	"kitabisa/logger"
)

type CampaignService interface {
	FindAllCampaign() ([]Campaign, error)
	FindByIdCampaign(int) (Campaign, error)
}

type campaignService struct {
	campaignRepositoryDB CampaignRepositoryDB
}

func NewCampaignService(campaignRepo CampaignRepositoryDB) *campaignService {
	return &campaignService{campaignRepositoryDB: campaignRepo}
}

func (c *campaignService) FindAllCampaign() ([]Campaign, error) {
	campaigns, err := c.campaignRepositoryDB.FindAllCampaign()
	if err != nil {
		logger.Error("Error service" + err.Error())
		return campaigns, err
	}
	return campaigns, nil
}

func (c *campaignService) FindByIdCampaign(id int) (Campaign, error) {
	campaign, err := c.campaignRepositoryDB.FindCampaignById(id)
	if err != nil {
		return campaign, err
	}
	if campaign.ID == 0 {
		logger.Error("campaign not found")
		return campaign, errors.New("campaign not found")
	}
	return campaign, nil
}
