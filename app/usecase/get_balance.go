package usecase

import (
	"context"

	"github.com/jirlon/digitalbank/app/entities"
)

type getBalanceAccountRepository interface {
	GetByAccountID(ctx context.Context, accountID string) (entities.Account, error)
}

type GetBalanceUC struct {
	balanceRepo getBalanceAccountRepository
}

func (uc GetBalanceUC) GetBalance(accountID string) (entities.Account, error) {

	account, err := uc.balanceRepo.GetByAccountID(context.Background(), accountID)
	return account, err
}

func NewGetBalanceUseCase(repo getBalanceAccountRepository) GetBalanceUC {
	return GetBalanceUC{balanceRepo: repo}
}
