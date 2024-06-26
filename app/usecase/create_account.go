package usecase

import (
	"context"

	"github.com/jirlon/digitalbank/app/entities"
)

type createAccountAccountRepository interface {
	SaveAccount(account entities.Account) error
}

type CreateAccountUseCase struct {
	repo createAccountAccountRepository
}

// Create a new account with the data provided.
func (uc CreateAccountUseCase) CreateAccount(ctx context.Context, cpf, name, secret string, balance int) (entities.Account, error) {

	// Calls the constructor to validate the data.
	account, err := entities.NewAccount(cpf, name, secret, balance)

	if err != nil {
		return entities.Account{}, err
	}

	err = uc.repo.SaveAccount(account)

	// If there was an error in the SaveAccount method, then return nil and the error.
	if err != nil {
		return entities.Account{}, err
	}

	return account, err
}

func NewCreateAccountUseCase(repo createAccountAccountRepository) CreateAccountUseCase {
	return CreateAccountUseCase{repo: repo}
}
