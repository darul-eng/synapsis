package repository

import (
	"context"
	"database/sql"
	"tes-synapsis/model/domain"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	FindById(ctx context.Context, tx *sql.Tx, transactionId int) (domain.Transaction, error)
}
