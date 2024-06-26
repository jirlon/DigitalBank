package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jirlon/digitalbank/app/entities"
)

type TransferRepository struct {
	q *pgxpool.Pool
}

func NewTransfer(q *pgxpool.Pool) *TransferRepository {
	return &TransferRepository{
		q: q,
	}
}

func (r TransferRepository) SaveTransfer(ctx context.Context, transfer entities.Transfer, originAccount, destinationAccount entities.Account) error {
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

	const (
		query = `INSERT INTO transfers (
					id,
					account_origin_id,
					account_destination_id,
					amount,
					created_at)
				VALUES ($1, $2, $3, $4, $5)`
	)

	_, err = tx.Exec(
		ctx,
		query,
		transfer.GetTransferID(),
		transfer.GetAccountOriginID(),
		transfer.GetAccountDestinationID(),
		transfer.GetAmount(),
		transfer.GetCreatedAt())
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
