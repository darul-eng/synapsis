package repository

import (
	"context"
	"database/sql"
	"tes-synapsis/helper"
	"tes-synapsis/model/domain"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := `INSERT INTO "transaction"("user_id", "amount") VALUES ($1, $2)`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, transaction.UserId, transaction.Amount)

	err := result.Scan(&lastInsertId)
	helper.PanicIfError(err)

	transaction.Id = lastInsertId

	return transaction
}
