package repository

import (
	"context"
	"database/sql"
	"tes-synapsis/helper"
	"tes-synapsis/model/domain"
)

type CartRepositoryImpl struct {
}

func (repository *CartRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := `INSERT INTO "cart"("product_id", "amount") VALUES ($1, $2) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, cart.ProductId, cart.Amount)

	err := result.Scan(&lastInsertId)
	helper.PanicIfError(err)

	cart.Id = lastInsertId

	return cart
}

func (repository *CartRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, Id int) {
	SQL := `DELETE FROM "cart" WHERE "id" = $1`
	_, err := tx.ExecContext(ctx, SQL, Id)
	helper.PanicIfError(err)
}

func (repository *CartRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Cart {
	SQL := `SELECT "id", "product_id", "amount" FROM "cart"`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(&cart.Id, &cart.ProductId, &cart.Amount)
		helper.PanicIfError(err)

		carts = append(carts, cart)
	}

	return carts
}

