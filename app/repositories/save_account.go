package repositories

import (
	"context"

	"github.com/jirlon/digitalbank/app/entities"
)

func (r AccountRepository) SaveAccount(account entities.Account) error {
	const (
		query = `INSERT INTO accounts (
					id, 
					cpf,
					name,
					secret,
					balance,
					created_at)
				VALUES ($1, $2, $3, $4, $5, $6)`
	)

	_, err := r.q.Exec(
		context.Background(),
		query,
		account.GetID(),
		account.GetCPF(),
		account.GetName(),
		account.GetSecret(),
		account.GetBalance(),
		account.GetCreatedAt())

	if err != nil {
		return err
	}
	return nil
}
