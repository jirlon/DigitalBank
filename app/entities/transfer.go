package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	id                   string
	accountOriginID      string
	accountDestinationID string
	amount               int
	createdAt            time.Time
}

func (t Transfer) GetTransferID() string {
	return t.id
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

func NewTransferHelper(accountOriginID, accountDestinationID string, amount int) Transfer {
	transfer, err := NewTransfer(accountOriginID, accountDestinationID, amount)
	if err != nil {
		return Transfer{}
	}
	return transfer
}

func NewTransfer(accountOriginID, accountDestinationID string, amount int) (Transfer, error) {
	if accountOriginID == "" {
		return Transfer{}, errors.New("empty originID")
	}
	if accountDestinationID == "" {
		return Transfer{}, errors.New("empty destinationID")
	}
	if accountOriginID == accountDestinationID {
		return Transfer{}, errors.New("origin and destination account are the same")
	}
	if amount <= 0 {
		return Transfer{}, errors.New("amount should be greater than zero")
	}

	transferID := uuid.New().String()

	transfer := Transfer{
		id:                   transferID,
		accountOriginID:      accountOriginID,
		accountDestinationID: accountDestinationID,
		amount:               amount,
		createdAt:            time.Now(),
	}

	return transfer, nil
}
