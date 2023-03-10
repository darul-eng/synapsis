package repository

import (
	"context"
	"database/sql"
	"errors"
	"tes-synapsis/helper"
	"tes-synapsis/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := `INSERT INTO "product"("name", "category_id", "price", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, product.Name, product.CategoryId, product.Price, product.CreatedAt, product.UpdatedAt)

	err := result.Scan(&lastInsertId)
	helper.PanicIfError(err)

	product.Id = lastInsertId

	return product
}

func (repository *ProductRepositoryImpl) FindByCategoryId(ctx context.Context, tx *sql.Tx, categoryId int) ([]domain.Product, error) {
	SQL := `SELECT "id", "name", "price", "category_id" FROM "product" WHERE "category_id" = $1`
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	if len(products) == 0 {
		return products, errors.New("product is not found")
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := `SELECT "id", "name", "price", "category_id" FROM "product" WHERE "id" = $1`
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId)
		helper.PanicIfError(err)

		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}
