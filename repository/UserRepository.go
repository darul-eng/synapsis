package repository

import (
	"context"
	"database/sql"
	"tes-synapsis/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Find(ctx context.Context, tx *sql.Tx, username string, password string) (domain.User, error)
}
