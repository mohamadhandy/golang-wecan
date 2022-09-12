package transaction

import "kitabisa/user"

type GetCampaignTransactionsInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	User       user.User
}
