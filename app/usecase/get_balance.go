package usecase

import "context"

type getBalanceAccountRepository interface {
	GetByAccountID(ctx context.Context, accountID string) (int, error)
}

type GetBalanceUC struct {
	balanceRepo getBalanceAccountRepository
}

func (uc GetBalanceUC) GetBalance(accountID string) (int, error) {

	balance, err := uc.balanceRepo.GetByAccountID(context.Background(), accountID)
	return balance, err
}

func NewGetBalanceUseCase(repo getBalanceAccountRepository) GetBalanceUC {
	return GetBalanceUC{balanceRepo: repo}
}
