package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jirlon/digitalbank/app/entities"
)

func (a AccountRepository) GetByAccountID(ctx context.Context, accountID string) (entities.Account, error) {
	row := a.q.QueryRow(ctx, "SELECT id, cpf, name, balance, created_at FROM accounts where id = $1", accountID)

	var (
		id        string
		cpf       string
		name      string
		balance   int
		createdAt time.Time
	)
	if err := row.Scan(&id, &cpf, &name, &balance, &createdAt); err != nil {
		if err == sql.ErrNoRows {
			return entities.Account{}, errors.New("account not found")
		}
	}
	account, err := entities.ParseAccount(id, cpf, name, "", balance, createdAt)
	if err != nil {
		return entities.Account{}, err
	}

	return account, nil
}
