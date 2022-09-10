package transaction

import "time"

type TransactionDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt time.Time
}
