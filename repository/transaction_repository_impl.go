package repository

import (
	"context"
	"database/sql"
	"errors"
	"tes-synapsis/helper"
	"tes-synapsis/model/domain"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := `INSERT INTO "transaction"("user_id", "amount", "status", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, transaction.UserId, transaction.Amount, transaction.Status, transaction.CreatedAt, transaction.UpdatedAt)

	err := result.Scan(&lastInsertId)
	helper.PanicIfError(err)

	transaction.Id = lastInsertId

	return transaction
}

func (repository *TransactionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := `UPDATE "transaction" SET "payment_url"=$1, "status"=$2, "updated_at"=$3 WHERE "id"=$4`
	_, err := tx.ExecContext(ctx, SQL, transaction.PaymentURL, transaction.Status, transaction.UpdatedAt, transaction.Id)
	helper.PanicIfError(err)

	return transaction
}

func (repository *TransactionRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, transactionId int) (domain.Transaction, error) {
	SQL := `SELECT "id", "payment_url", "status", "updated_at" FROM "transaction" WHERE "id"=$1`
	rows, err := tx.QueryContext(ctx, SQL, transactionId)
	helper.PanicIfError(err)
	defer rows.Close()

	transactions := domain.Transaction{}
	if rows.Next() {
		err := rows.Scan(&transactions.Id, &transactions.PaymentURL, &transactions.Status, &transactions.UpdatedAt)
		helper.PanicIfError(err)

		return transactions, nil
	} else {
		return transactions, errors.New("transaction is not found")
	}
}
