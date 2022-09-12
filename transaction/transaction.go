package transaction

import (
	"kitabisa/user"
	"time"
)

type Transaction struct {
	ID         int    `gorm:"column:transaction_id"`
	UserId     int    `gorm:"column:user_id"`
	CampaignId int    `gorm:"column:campaign_id"`
	Amount     int    `gorm:"column:amount"`
	Status     string `gorm:"column:status"`
	Code       string `gorm:"column:code"`
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
