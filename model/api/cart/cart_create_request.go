package cart

type CartCreateRequest struct {
	ProductId int `validate:"required" json:"product_id"`
}
