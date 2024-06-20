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
	var ctx context.Context
	return uc.repo.FindAll(ctx)
}

func NewListAccountUseCase(repo listAccountAccountRepository) ListAccountUC {
	return ListAccountUC{repo: repo}
}
