package usecase

import "github.com/jirlon/digitalbank/app/entities"

type listAccountAccountRepository interface {
	FindAll() ([]entities.Account, error)
}

type ListAccountUC struct {
	repo listAccountAccountRepository
}

func (uc ListAccountUC) ListAccount() ([]entities.Account, error) {
	return uc.repo.FindAll()
}

func NewListAccountUseCase(repo listAccountAccountRepository) ListAccountUC {
	return ListAccountUC{repo: repo}
}
