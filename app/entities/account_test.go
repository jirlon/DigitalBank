package entities

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	t.Parallel()

	type args struct {
		cpf     string
		name    string
		secret  string
		balance int
	}

	tests := []struct {
		nameT  string
		args   args
		want   Account
		errMsg error
	}{
		{
			nameT: "Create new account successfully",
			args: args{
				cpf:     "90948239000",
				name:    "Ronaldo Fenômeno",
				secret:  "r9",
				balance: 1000,
			},
			want: Account{
				cpf:     "90948239000",
				name:    "Ronaldo Fenômeno",
				secret:  "r9",
				balance: 1000,
			},
			errMsg: nil,
		},
		{
			nameT: "Invalid CPF length",
			args: args{
				cpf:     "123456789",
				name:    "Chico Buarque",
				secret:  "tiochico",
				balance: 2000,
			},
			want:   Account{},
			errMsg: errors.New("invalid CPF length"),
		},

		{
			nameT: "Invalid CPF",
			args: args{
				cpf:     "12345678901",
				name:    "machado de Assis",
				secret:  "DomCasmurro",
				balance: 3000,
			},
			want:   Account{},
			errMsg: errors.New("invalid CPF"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.nameT, func(t *testing.T) {

			account, err := NewAccount(tt.args.cpf, tt.args.name, tt.args.secret, tt.args.balance)

			if tt.errMsg != nil {
				assert.EqualError(t, err, tt.errMsg.Error())
			} else {
				assert.NotEmpty(t, account.id)
				account.id = ""
				assert.NotEmpty(t, account.createdAt)
				account.createdAt = time.Time{}
				assert.Equal(t, tt.want, account)
			}
		})
	}
}
