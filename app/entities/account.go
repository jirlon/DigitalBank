package entities

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	id        string
	name      string
	cpf       string
	secret    string
	balance   int
	createdAt time.Time
}

func NewAccountHelper(cpf, name, secret string, balance int) Account {
	account, err := NewAccount(cpf, name, secret, balance)
	if err != nil {
		return Account{}
	}
	return account
}

func ValidaCPF(cpf string) error {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return errors.New("invalid CPF length")
	}

	// Calcula o primeiro dígito verificador
	var soma int
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (10 - i)
	}
	resto := soma % 11
	var primeiroDigitoVerificador int
	if resto < 2 {
		primeiroDigitoVerificador = 0
	} else {
		primeiroDigitoVerificador = 11 - resto
	}

	// Calcula o segundo dígito verificador
	soma = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * (11 - i)
	}
	resto = soma % 11
	var segundoDigitoVerificador int
	if resto < 2 {
		segundoDigitoVerificador = 0
	} else {
		segundoDigitoVerificador = 11 - resto
	}

	// Verifica se os dígitos verificadores são corretos
	if primeiroDigitoVerificador == int(cpf[9]-'0') && segundoDigitoVerificador == int(cpf[10]-'0') {
		return nil
	}
	return errors.New("invalid CPF")
}

func (a Account) GetID() string {
	return a.id
}

func (a Account) GetCPF() string {
	return a.cpf
}

func (a Account) GetName() string {
	return a.name
}

func (a Account) GetSecret() string {
	return a.secret
}

func (a Account) GetBalance() int {
	return a.balance
}

func (a Account) GetCreatedAt() time.Time {
	return a.createdAt
}

func (a *Account) SetBalance(amount int) {
	a.balance += amount
}

func (a *Account) SubtractBalance(amount int) {
	if a.balance >= amount {
		a.balance -= amount
	}
}

func NewAccount(cpf, name, secret string, balance int) (Account, error) {

	if cpf == "" {
		return Account{}, errors.New("empty cpf")
	}
	if name == "" {
		return Account{}, errors.New("empty name")
	}
	if secret == "" {
		return Account{}, errors.New("empty secret")
	}

	err := ValidaCPF(cpf)
	if err != nil {
		return Account{}, err
	}

	// Generates a UUID.
	id := uuid.New().String()

	account := Account{
		id:        id,
		cpf:       cpf,
		name:      name,
		secret:    secret,
		balance:   balance,
		createdAt: time.Now(),
	}

	return account, nil
}

func ParseAccount(id, cpf, name, secret string, balance int, createdAt time.Time) (Account, error) {

	if cpf == "" {
		return Account{}, errors.New("empty cpf")
	}
	if name == "" {
		return Account{}, errors.New("empty name")
	}

	err := ValidaCPF(cpf)
	if err != nil {
		return Account{}, err
	}

	account := Account{
		id:        id,
		cpf:       cpf,
		name:      name,
		secret:    secret,
		balance:   balance,
		createdAt: createdAt,
	}

	return account, nil
}
