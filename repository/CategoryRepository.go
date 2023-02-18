package repository

import (
	"context"
	"database/sql"
	"tes-synapsis/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
}
