package repository

import (
	"errors"
	"testing"

	"github.com/jirlon/digitalbank/app/entities"
)

func TestMockAccountRepository_Save(t *testing.T) {
	repo := NewMockAccountRepository()

	t.Run("save account sucessfully", func(t *testing.T) {
		t.Parallel()
		id := "1"
		cpf := "123.456.789-09"
		name := "maria"
		secret := "onemore"
		account, err := entities.NewAccount(id, cpf, name, secret)
		if err != nil {
			t.Fatalf("unexpected error creating account: %v", err)
		}

		err = repo.SaveAccount(account)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("error saving account", func(t *testing.T) {
		t.Parallel()
		id := "2"
		cpf := "987.654.321-00"
		nome := "carlos"
		secret := "zzz"
		account, err := entities.NewAccount(id, cpf, nome, secret)
		if err != nil {
			t.Fatalf("unexpected error creating account: %v", err)
		}

		repo.ErrSave = errors.New("error saving account")
		err = repo.SaveAccount(account)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if err.Error() == "error saving account" {
			t.Errorf("expected error 'error saving account', got %v", err)
		}
	})
}
