package repositories

import (
	"context"
	"database/sql"
	"errors"
)

func (a AccountRepository) GetByAccountID(ctx context.Context, accountID string) (int, error) {
	row := a.q.QueryRow(ctx, "SELECT balance FROM accounts where id = $1", accountID)

	var balance int

	if err := row.Scan(&balance); err != nil {
		if err == sql.ErrNoRows {
			return balance, errors.New("account not found")
		}
		return balance, err
	}
	return balance, nil
}
