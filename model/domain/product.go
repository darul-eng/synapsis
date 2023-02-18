package domain

type Product struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
	Price      int    `json:"price"`
}
