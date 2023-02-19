package repository

import (
	"context"
	"database/sql"
	"tes-synapsis/helper"
	"tes-synapsis/model/domain"
)

type CartRepositoryImpl struct {
}

func NewCartRepository() CartRepository {
	return &CartRepositoryImpl{}
}

func (repository *CartRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := `INSERT INTO "cart"("product_id", "amount", "created_at", "updated_at") VALUES ($1, $2, $3, $4) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, cart.ProductId, cart.Amount, cart.CreatedAt, cart.UpdatedAt)

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
	SQL := `SELECT t1.id, t1.product_id, t2.name, t1.amount, t2.price, t3.name FROM cart AS t1 JOIN product AS t2 ON t1.product_id = t2.id JOIN category AS t3 ON t2.category_id = t3.id`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(&cart.Id, &cart.ProductId, &cart.Name, &cart.Amount, &cart.Price, &cart.Category)
		helper.PanicIfError(err)

		carts = append(carts, cart)
	}

	return carts
}

func (repository CartRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, cartId int) domain.Cart {
	SQL := `SELECT t1.id, t1.product_id, t2.name, t1.amount, t2.price FROM cart AS t1 JOIN product AS t2 ON t1.product_id = t2.id JOIN category AS t3 ON t2.category_id = t3.id WHERE t1.id = $1`
	rows, err := tx.QueryContext(ctx, SQL, cartId)
	helper.PanicIfError(err)
	defer rows.Close()

	cart := domain.Cart{}
	if rows.Next() {
		err := rows.Scan(&cart.Id, &cart.ProductId, &cart.Name, &cart.Amount, &cart.Price, &cart.Category)
		helper.PanicIfError(err)
	}
	return cart
}
