package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/jirlon/digitalbank/app/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountUseCase(t *testing.T) {

	t.Parallel()

	type args struct {
		cpf     string
		name    string
		secret  string
		balance int
	}

	tests := []struct {
		nameT   string
		args    args
		setup   func(t *testing.T) CreateAccountUseCase
		want    entities.Account
		wantErr bool
		errMsg  error
	}{
		{
			nameT: "Create new account successfully",
			args: args{
				cpf:     "90948239000",
				name:    "Ronaldo Fenômeno",
				secret:  "r9",
				balance: 1000,
			},
			setup: func(t *testing.T) CreateAccountUseCase {
				return CreateAccountUseCase{
					repo: MockAccountRepository{
						ErrSave:  nil,
						accounts: map[string]entities.Account{},
					},
				}
			},
			want: func() entities.Account {
				account := entities.NewAccountHelper("90948239000", "Ronaldo Fenômeno", "r9", 1000)
				return account
			}(),
			wantErr: false,
			errMsg:  nil,
		},

		{
			nameT: "Error saving account",
			args: args{
				cpf:     "90948239000",
				name:    "Ronaldinho Gaucho",
				secret:  "r10",
				balance: 3000,
			},
			setup: func(t *testing.T) CreateAccountUseCase {
				return CreateAccountUseCase{
					repo: MockAccountRepository{
						ErrSave:  errors.New("error saving account"),
						accounts: map[string]entities.Account{},
					},
				}
			},
			want: func() entities.Account {
				account := entities.NewAccountHelper("90948239000", "Ronaldinho Gaucho", "r10", 3000)
				return account
			}(),
			wantErr: true,
			errMsg:  errors.New("error saving account"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.nameT, func(t *testing.T) {
			t.Parallel()

			account, err := tt.setup(t).CreateAccount(context.Background(), tt.args.cpf, tt.args.name, tt.args.secret, tt.args.balance)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != nil {
					assert.EqualError(t, err, tt.errMsg.Error())
					assert.Empty(t, account)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.GetCPF(), account.GetCPF())
				assert.Equal(t, tt.want.GetName(), account.GetName())
				assert.Equal(t, tt.want.GetSecret(), account.GetSecret())
				assert.Equal(t, tt.want.GetBalance(), account.GetBalance())
				assert.NotEmpty(t, account.GetID())

				assert.NotEmpty(t, account.GetCreatedAt())

			}
		})
	}
}
