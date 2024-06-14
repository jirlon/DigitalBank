package usecase

import (
	"github.com/jirlon/digitalbank/app/entities"
)

type MockAccountRepository struct {
	accounts map[string]entities.Account
	ErrSave  error
}

func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{
		accounts: make(map[string]entities.Account),
	}
}

func (m MockAccountRepository) SaveAccount(account entities.Account) error {
	//if there is an error it does not save and returns the error
	if m.ErrSave != nil {
		return m.ErrSave
	}
	//there is no error, so save the account and return nil
	m.accounts[account.GetID()] = account
	return nil
}
