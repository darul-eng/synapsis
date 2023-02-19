package service

import (
	"context"
	"tes-synapsis/model/api/transaction"
)

type TransactionService interface {
	Create(ctx context.Context, request transaction.TransactionCreateRequest) transaction.TransactionResponse
}
