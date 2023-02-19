package product

type ProductCreateRequest struct {
	Name       string  `validate:"required" json:"name"`
	CategoryId int     `validate:"required" json:"category_id"`
	Price      float64 `validate:"required" json:"price"`
}
