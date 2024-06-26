package repositories

import (
	"context"

	"github.com/jirlon/digitalbank/app/entities"
)

func (r AccountRepository) SaveTransfer(ctx context.Context, transfer entities.Transfer, originAccount, destinationAccount entities.Account) error {
	tx, err := r.q.Begin(ctx)
	if err != nil {
		return err
	}

	updateBalanceQuery := "UPDATE accounts SET balance = $1 WHERE id = $2"

	_, err = tx.Exec(ctx, updateBalanceQuery, originAccount.GetBalance(), originAccount.GetID())
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, updateBalanceQuery, destinationAccount.GetBalance(), destinationAccount.GetID())
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
