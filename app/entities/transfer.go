package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	transferID           string
	accountOriginID      string
	accountDestinationID string
	amount               int
	createdAt            time.Time
}

func (t Transfer) GetTransferID() string {
	return t.transferID
}

func (t Transfer) GetAccountOriginID() string {
	return t.accountOriginID
}

func (t Transfer) GetAccountDestinationID() string {
	return t.accountDestinationID
}

func (t Transfer) GetAmount() int {
	return t.amount
}

func (t Transfer) GetCreatedAt() time.Time {
	return t.createdAt
}

func NewTransfer(accountOriginID, accountDestinationID string, amount int) (Transfer, error) {
	if accountOriginID == "" {
		return Transfer{}, errors.New("empty originID")
	}
	if accountDestinationID == "" {
		return Transfer{}, errors.New("empty destinationID")
	}
	if amount == 0 {
		return Transfer{}, errors.New("empty amount")
	}

	transferID := uuid.New().String()

	transfer := Transfer{
		transferID:           transferID,
		accountOriginID:      accountOriginID,
		accountDestinationID: accountDestinationID,
		amount:               amount,
		createdAt:            time.Now(),
	}

	return transfer, nil
}
