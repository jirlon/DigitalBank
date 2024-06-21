package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jirlon/digitalbank/app/entities"
)

type AccountRepository struct {
	q *pgxpool.Pool
}

func New(q *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{
		q: q,
	}
}

func (a AccountRepository) FindAll(ctx context.Context) ([]entities.Account, error) {
	rows, err := a.q.Query(ctx, "SELECT id, cpf, name, balance, created_at FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []entities.Account
	for rows.Next() {

		var (
			id        string
			cpf       string
			name      string
			balance   int
			createdAt time.Time
		)

		if err := rows.Scan(&id, &cpf, &name, &balance, &createdAt); err != nil {
			return nil, err
		}
		account, err := entities.ParseAccount(id, cpf, name, "", balance, createdAt)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
