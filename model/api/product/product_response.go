package product

type ProductResponse struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryId int     `json:"category_id"`
}
