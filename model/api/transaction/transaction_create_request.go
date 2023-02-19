package transaction

type TransactionCreateRequest struct {
	UserId int     `validate:"required" json:"user_id"`
	Amount float64 `validate:"required" json:"amount"`
}
