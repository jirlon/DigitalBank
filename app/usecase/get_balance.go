package usecase

import "context"

type getBalanceAccountRepository interface {
	GetByAccountID(ctx context.Context, accountID string) (int, error)
}

type GetBalanceUC struct {
	balanceRepo getBalanceAccountRepository
}

func (uc GetBalanceUC) GetBalance(accountID string) (int, error) {
	var ctx context.Context
	balance, err := uc.balanceRepo.GetByAccountID(ctx, accountID)
	return balance, err
}

func NewGetBalanceUseCase(repo getBalanceAccountRepository) GetBalanceUC {
	return GetBalanceUC{balanceRepo: repo}
}
