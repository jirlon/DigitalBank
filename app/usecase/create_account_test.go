package usecase

import (
	"errors"
	"testing"

	"github.com/jirlon/digitalbank/app/repository"
)

func TestCreateAccountUseCase(t *testing.T) {
	mockRepo := repository.NewMockAccountRepository()
	useCase := NewCreateAccountUseCase(mockRepo)

	// Test Case: Successful account creation.
	t.Run("Create new account successfully", func(t *testing.T) {
		t.Parallel()
		cpf := "123.456.789-00"
		name := "jirlon"
		secret := "chavesegura"
		account, err := useCase.CreateAccount(cpf, name, secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		// Check if the account was created correctly.
		if account.GetCPF() != cpf {
			t.Errorf("expected CPF %v, got %v", cpf, account.GetCPF())
		}
		if account.GetName() != name {
			t.Errorf("expected name %v, got %v", name, account.GetName())
		}
		if account.GetSecret() != secret {
			t.Errorf("expected secret %v, got %v", secret, account.GetSecret())
		}
		if account.GetBalance() != 0 {
			t.Errorf("expected balance 0, got %v", account.GetBalance())
		}
		if account.GetCreatedAt().IsZero() {
			t.Errorf("expected non-zero created_at")
		}
	})

	// Test case: Empty field.
	t.Run("Invalid input", func(t *testing.T) {
		t.Parallel()
		cpf := "123.456.789.00"
		name := ""
		secret := ""
		if EmptyField(cpf, name, secret) {
			t.Errorf("invalid input, some field is empty")
		}
	})

	// Test case: Invalid CPF.
	t.Run("Invalid CPF", func(t *testing.T) {
		t.Parallel()
		cpf := "123.456.789-00"
		name := "jirlon"
		secret := "chavesegura"
		_, err := useCase.CreateAccount(cpf, name, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err.Error() != "invalid CPF" {
			t.Errorf("expected error 'invalid CPF', got %v", err)
		}
	})

	// Test case: error saving account.
	t.Run("Error saving account", func(t *testing.T) {
		t.Parallel()
		cpf := "123.456.789-09"
		name := "jiw"
		secret := "chaveinsegura"
		mockRepo.ErrSave = errors.New("error saving account")
		_, err := useCase.CreateAccount(cpf, name, secret)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err.Error() == "error saving account" {
			t.Errorf("expected error 'error saving account', got %v", err)
		}
	})

}
