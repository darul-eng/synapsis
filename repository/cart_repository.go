package repository

import (
	"context"
	"database/sql"
	"tes-synapsis/model/domain"
)

type CartRepository interface {
	Save(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart
	Delete(ctx context.Context, tx *sql.Tx, Id int)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Cart
}
