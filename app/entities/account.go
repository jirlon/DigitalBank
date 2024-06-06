package entities

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Represents a user account in the digital bank.
type Account struct {
	id         string
	name       string
	cpf        string
	secret     string
	balance    int64
	created_at time.Time
}

// constructor
func NewAccount(id, cpf, name, secret string) (*Account, error) {

	//Creates a new instance of Account.
	account := &Account{
		id:         id,
		cpf:        cpf,
		name:       name,
		secret:     secret,
		balance:    0,
		created_at: time.Now(),
	}

	return account, nil
}

// Check if any fields are empty.
func EmptyField(cpf, name, secret string) bool {
	if cpf == "" || name == "" || secret == "" {
		return true
	}
	return false
}

// Validates a Brazilian CPF.
func ValidaCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Verifica se o CPF tem 11 dígitos
	if len(cpf) != 11 {
		return false
	}

	// Verifica se todos os dígitos são iguais
	if matched, _ := regexp.MatchString(`(\d)\1{10}`, cpf); matched {
		return false
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
	return primeiroDigitoVerificador == int(cpf[9]-'0') && segundoDigitoVerificador == int(cpf[10]-'0')
}

// Returns the account's ID.
func (a *Account) GetID() string {
	return a.id
}

// Returns the account's CPF.
func (a *Account) GetCPF() string {
	return a.cpf
}

// Returns the account's name.
func (a *Account) GetName() string {
	return a.name
}

// Returns the account's secret.
func (a *Account) GetSecret() string {
	return a.secret
}

// Returns the account's balance.
func (a *Account) GetBalance() int64 {
	return a.balance
}

// Returns the account creation date.
func (a *Account) GetCreatedAt() time.Time {
	return a.created_at
}
