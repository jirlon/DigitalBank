package usecase

import (
	"context"

	"github.com/jirlon/digitalbank/app/entities"
)

type listAccountAccountRepository interface {
	FindAll(ctx context.Context) ([]entities.Account, error)
}

type ListAccountUC struct {
	repo listAccountAccountRepository
}

func (uc ListAccountUC) ListAccount() ([]entities.Account, error) {
	return uc.repo.FindAll(context.Background())
}

func NewListAccountUseCase(repo listAccountAccountRepository) ListAccountUC {
	return ListAccountUC{repo: repo}
}
