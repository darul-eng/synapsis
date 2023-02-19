package domain

import "time"

type Cart struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	Price     float64   `json:"price"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
