package campaign

import "time"

type Campaign struct {
	ID               int       `json:"id" gorm:"column:campaign_id"`
	UserId           int       `json:"user_id" gorm:"column:user_id"`
	CampaignName     string    `json:"campaign_name" gorm:"column:campaign_name"`
	ShortDescription string    `json:"short_description" gorm:"column:short_description"`
	Description      string    `json:"description" gorm:"column:description"`
	GoalAmount       int       `json:"goal_amount" gorm:"column:goal_amount"`
	CurrentAmount    int       `json:"current_amount" gorm:"column:current_amount"`
	Perks            string    `json:"perks" gorm:"column:perks"`
	BackerCount      int       `json:"backer_count" gorm:"column:backer_count"`
	Slug             string    `json:"slug" gorm:"column:slug"`
	CreatedAt        time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
