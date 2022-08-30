package campaign

import "kitabisa/logger"

type CampaignService interface {
	FindAllCampaign() ([]Campaign, error)
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
