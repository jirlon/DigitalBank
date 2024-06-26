package usecase

import (
	"context"
	"errors"

	"github.com/jirlon/digitalbank/app/entities"
)

type MockCreateTransferAccountRepository struct {
	accounts map[string]entities.Account
	ErrSave  error
	ErrGet   error
}

type MockCreateTransferTransferRepository struct {
	transfers []entities.Transfer
	SaveErr   error
}

func (m *MockCreateTransferTransferRepository) SaveTransfer(ctx context.Context, transfer entities.Transfer, originAccount, destinationAccount entities.Account) error {
	if m.SaveErr != nil {
		return m.SaveErr
	}
	m.transfers = append(m.transfers, transfer)
	return nil
}

func (m MockCreateTransferAccountRepository) GetByAccountID(ctx context.Context, accountID string) (entities.Account, error) {
	if m.ErrGet != nil {
		return entities.Account{}, m.ErrGet
	}
	account, err := m.accounts[accountID]
	if !err {
		return entities.Account{}, errors.New("account not found")
	}
	return account, nil
}

func (m MockCreateTransferAccountRepository) SaveAccount(account entities.Account) error {
	if m.ErrSave != nil {
		return m.ErrSave
	}
	m.accounts[account.GetID()] = account
	return nil
}

func NewMockCreateTransferRepository() (*MockCreateTransferAccountRepository, *MockCreateTransferTransferRepository) {
	return &MockCreateTransferAccountRepository{
			accounts: make(map[string]entities.Account),
		}, &MockCreateTransferTransferRepository{
			transfers: make([]entities.Transfer, 0),
		}
}
