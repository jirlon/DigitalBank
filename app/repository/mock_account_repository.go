package repository

import (
	"github.com/jirlon/digitalbank/app/entities"
)

// Mock implementation of AccountRepository.
type MockAccountRepository struct {
	accounts map[string]*entities.Account
	ErrSave  error
}

// Creates a new MockAccountRepository instance.
func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{
		accounts: make(map[string]*entities.Account),
	}
}

// Saves an account to the mock repository.
func (r *MockAccountRepository) SaveAccount(account *entities.Account) error {
	//if there is an error it does not save and returns the error
	if r.ErrSave != nil {
		return r.ErrSave
	}
	//there is no error, so save the account and return nil
	r.accounts[account.GetID()] = account
	return nil
}
