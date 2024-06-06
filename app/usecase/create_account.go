package usecase

//packages imports.
import (
	"errors"

	"github.com/google/uuid"
	"github.com/jirlon/digitalbank/app/entities"
	"github.com/jirlon/digitalbank/app/repository"
)

// Defines the interface for the account creation use case.
type CreateAccountUseCase interface {
	CreateAccount(id, cpf, name, secret string) (*entities.Account, error)
}

// Implementation of the CreateAccountUseCase.
type createAccountUseCase struct {
	repo repository.AccountRepository
}

// Creates a new instance of the CreateAccountUseCase.
func NewCreateAccountUseCase(repo repository.AccountRepository) *createAccountUseCase {
	return &createAccountUseCase{repo: repo}
}

// Create a new account with the data provided.
func (uc *createAccountUseCase) CreateAccount(cpf, name, secret string) (*entities.Account, error) {

	// Checks if any field is empty.
	if entities.EmptyField(cpf, name, secret) {
		return nil, errors.New("invalid input, some field is empty")
	}

	// Checks CPF validation.
	if !entities.ValidaCPF(cpf) {
		return nil, errors.New("invalid CPF")
	}

	// Generates a UUID.
	id := uuid.New().String()

	// Calls the constructor to validate the data.
	account, err := entities.NewAccount(id, cpf, name, secret)

	// If there was an error, return nil and the error.
	if err != nil {
		return nil, err
	}

	// There was no error, calls the method to save the account, which has a boolean return.
	err = uc.repo.SaveAccount(account)

	// If there was an error in the SaveAccount method, then return nil and the error.
	if err != nil {
		return nil, err
	}

	return account, err
}
