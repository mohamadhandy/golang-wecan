package campaign

import "time"

type Campaign struct {
	ID               int       `json:"id" gorm:"campaign_id"`
	UserId           int       `json:"user_id" gorm:"user_id"`
	CampaignName     string    `json:"campaign_name" gorm:"campaign_name"`
	ShortDescription string    `json:"short_description" gorm:"short_description"`
	Description      string    `json:"description" gorm:"description"`
	GoalAmount       int       `json:"goal_amount" gorm:"goal_amount"`
	CurrentAmount    int       `json:"current_amount" gorm:"current_amount"`
	Perks            string    `json:"perks" gorm:"perks"`
	BackerCount      int       `json:"backer_count" gorm:"backer_count"`
	Slug             string    `json:"slug" gorm:"slug"`
	CreatedAt        time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"updated_at"`
}
