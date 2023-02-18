package domain

type Cart struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	Amount    int `json:"amount"`
}
