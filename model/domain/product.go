package domain

import "time"

type Product struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CategoryId int       `json:"category_id"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
