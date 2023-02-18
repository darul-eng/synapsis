package repository

import (
	"context"
	"database/sql"
	"errors"
	"tes-synapsis/helper"
	"tes-synapsis/model/domain"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `INSERT INTO "user"("username", "email", "password") VALUES ($1, $2, $3) RETURNING id`
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, user.Username, user.Email, user.Password)

	err := result.Scan(&lastInsertId)
	helper.PanicIfError(err)

	user.Id = lastInsertId

	return user
}

func (repository *UserRepositoryImpl) Find(ctx context.Context, tx *sql.Tx, username string, password string) (domain.User, error) {
	SQL := `SELECT "id", "username" FROM "user" WHERE "username" = $1 AND "password" = $2`
	rows, err := tx.QueryContext(ctx, SQL, username, password)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}
