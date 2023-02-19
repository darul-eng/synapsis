package repository

import (
	"context"
	"database/sql"
	"tes-synapsis/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	FindByCategoryId(ctx context.Context, tx *sql.Tx, categoryId int) ([]domain.Product, error)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error)
}
