package domain

type Transaction struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	Amount int `json:"amount"`
}
