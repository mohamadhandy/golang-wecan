package campaign

import (
	"errors"
	"fmt"
	"kitabisa/logger"

	"github.com/gosimple/slug"
)

type CampaignService interface {
	FindAllCampaign() ([]Campaign, error)
	FindByIdCampaign(int) (Campaign, error)
	CreateCampaign(CreateCampaignInput) (Campaign, error)
	UpdateCampaign(GetCampaignDetailInput, CreateCampaignInput) (Campaign, error)
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

func (c *campaignService) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.CampaignName = input.Name
	campaign.Description = input.Description
	campaign.ShortDescription = input.ShortDescription
	campaign.Perks = input.Perks
	campaign.UserId = input.UserId
	campaign.GoalAmount = input.GoalAmount

	slugString := fmt.Sprintf("%s %d", input.Name, input.UserId)
	campaign.Slug = slug.Make(slugString)

	newCampaign, err := c.campaignRepositoryDB.CreateCampaign(campaign)
	if err != nil {
		logger.Error("Error new campaign" + err.Error())
		return newCampaign, err
	}
	return newCampaign, nil
}

func (c *campaignService) UpdateCampaign(inputGet GetCampaignDetailInput, inputCreate CreateCampaignInput) (Campaign, error) {
	campaign, err := c.campaignRepositoryDB.FindCampaignById(inputGet.ID)
	if err != nil {
		return campaign, err
	}
	if campaign.UserId != inputCreate.UserId {
		return campaign, errors.New("not an owner of the campaign")
	}
	campaign.CampaignName = inputCreate.Name
	campaign.Description = inputCreate.Description
	campaign.ShortDescription = inputCreate.ShortDescription
	campaign.Perks = inputCreate.Perks
	campaign.GoalAmount = inputCreate.GoalAmount

	updatedCampaign, err := c.campaignRepositoryDB.UpdateCampaign(campaign)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}
