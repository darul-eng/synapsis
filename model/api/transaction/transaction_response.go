package transaction

import (
	"tes-synapsis/model/domain"
	"time"
)

type TransactionResponse struct {
	Id         int                    `json:"id"`
	UserId     int                    `json:"user_id"`
	Amount     float64                `json:"amount"`
	Status     string                 `json:"status"`
	PaymentURL string                 `json:"payment_url"`
	User       domain.UserTransaction `json:"user"`
	CreatedAt  time.Time              `json:"created_at"`
}
