package transaction

import "time"

type TransactionResponse struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
