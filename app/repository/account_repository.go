package repository

import "github.com/jirlon/digitalbank/app/entities"

// Defines the interface for account persistence operations.
type AccountRepository interface {
	SaveAccount(account *entities.Account) error
}
