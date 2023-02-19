package cart

type CartResponse struct {
	Id        int     `json:"id"`
	ProductId int     `json:"product_id"`
	Name      string  `json:"name"`
	Amount    int     `json:"amount"`
	Price     float64 `json:"price"`
	Category  string  `json:"category"`
}
