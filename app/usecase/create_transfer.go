package usecase

import (
	"context"
	"errors"

	"github.com/jirlon/digitalbank/app/entities"
)

type createTransferAccountRepository interface {
	GetByAccountID(ctx context.Context, accountID string) (entities.Account, error)
	SaveAccount(account entities.Account) error
}

type createTransferTransferRepository interface {
	SaveTransfer(ctx context.Context, transfer entities.Transfer, originAccount, destinationAccount entities.Account) error
}

type CreateTransferUC struct {
	transferRepo createTransferTransferRepository
	accountRepo  createTransferAccountRepository
}

func (uc CreateTransferUC) CreateTransfer(ctx context.Context, accountOriginID, accountDestinationID string, amount int) (entities.Transfer, error) {

	originAccount, err := uc.accountRepo.GetByAccountID(ctx, accountOriginID)
	if err != nil {
		return entities.Transfer{}, errors.New("origin account not found")
	}

	destinationAccount, err := uc.accountRepo.GetByAccountID(ctx, accountDestinationID)
	if err != nil {
		return entities.Transfer{}, errors.New("destination account not found")
	}

	if originAccount.GetBalance() < amount {
		return entities.Transfer{}, errors.New("insufficient founds")
	}

	transfer, err := entities.NewTransfer(accountOriginID, accountDestinationID, amount)
	if err != nil {
		return entities.Transfer{}, err
	}

	originAccount.SubtractBalance(amount)
	destinationAccount.SetBalance(amount)

	if err := uc.transferRepo.SaveTransfer(ctx, transfer, originAccount, destinationAccount); err != nil {
		return entities.Transfer{}, err
	}

	return transfer, nil
}

func NewCreateTransferUseCase(tr createTransferTransferRepository, ar createTransferAccountRepository) CreateTransferUC {
	return CreateTransferUC{
		transferRepo: tr,
		accountRepo:  ar,
	}
}
