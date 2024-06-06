package entities

import "testing"

func TestEmptyField(t *testing.T) {
	t.Run("Invalid input", func(t *testing.T) {
		t.Parallel()
		cpf := "123.456.789.00"
		name := ""
		secret := ""
		if EmptyField(cpf, name, secret) {
			t.Errorf("invalid input, some field is empty")
		}
	})
}

func TestValidaCPF(t *testing.T) {
	t.Run("Valid CPF", func(t *testing.T) {
		t.Parallel()
		validaCPF := "123.456.789-09"
		if !ValidaCPF(validaCPF) {
			t.Errorf("expected CPF %v to be valid", validaCPF)
		}
	})

	t.Run("Invalid CPF", func(t *testing.T) {
		t.Parallel()
		invalidCPF := "123.456.789-00"
		if ValidaCPF(invalidCPF) {
			t.Errorf("expected CPF %v to be invalid", invalidCPF)
		}
	})

	t.Run("CPF with all identical digits", func(t *testing.T) {
		t.Parallel()
		identicalCPF := "333.333.333-33"
		if ValidaCPF(identicalCPF) {
			t.Errorf("expected CPF %v to be invalid", identicalCPF)
		}
	})

	t.Run("CPF with less than 11 digits", func(t *testing.T) {
		t.Parallel()
		shortCPF := "123.456.789-0"
		if ValidaCPF(shortCPF) {
			t.Errorf("expected CPF %v to be invalid", shortCPF)
		}
	})
}

func TestNewAccount(t *testing.T) {
	t.Run("Create new account successfully", func(t *testing.T) {
		t.Parallel()
		id := "1"
		cpf := "123.456.789-09"
		name := "jirlon"
		secret := "BoaNoite"
		account, err := NewAccount(id, cpf, name, secret)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if account.GetID() != id {
			t.Errorf("expected ID %v, got %v", id, account.GetID())
		}
		if account.GetCPF() != cpf {
			t.Errorf("expected CPF %v, got %v", cpf, account.GetCPF())
		}
		if account.GetName() != name {
			t.Errorf("expected nome %v, got %v", name, account.GetName())
		}
		if account.GetSecret() != secret {
			t.Errorf("expected nome %v, got %v", secret, account.GetSecret())
		}
		if account.GetBalance() != 0 {
			t.Errorf("expected balance 0, got %v", account.GetBalance())
		}
		if account.GetCreatedAt().IsZero() {
			t.Errorf("expected non-zero CreatedAt")
		}
	})
}
