package repository

import (
	"context"
	"database/sql"
	"errors"
	"tes-synapsis/helper"
	"tes-synapsis/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := `INSERT INTO "category"("name", "created_at", "updated_at") VALUES ($1, $2, $3) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, category.Name, category.CreatedAt, category.UpdatedAt)

	err := result.Scan(&lastInsertId)
	helper.PanicIfError(err)

	category.Id = lastInsertId

	return category
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := `SELECT "id", "name" FROM "category" WHERE "id" = $1`
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}
