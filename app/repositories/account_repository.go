package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepository struct {
	q *pgxpool.Pool
}

func New(q *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{
		q: q,
	}
}
