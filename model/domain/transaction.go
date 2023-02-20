package domain

import "time"

type Transaction struct {
	Id         int       `json:"id"`
	UserId     int       `json:"user_id"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	PaymentURL string    `json:"payment_url"`
	User       User      `json:"user"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
