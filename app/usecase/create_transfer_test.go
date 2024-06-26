package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/jirlon/digitalbank/app/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTransferUseCase(t *testing.T) {
	t.Parallel()

	type args struct {
		accountOriginID      string
		accountDestinationID string
		amount               int
	}

	tests := []struct {
		nameT  string
		args   args
		setup  func(t *testing.T) CreateTransferUC
		want   entities.Transfer
		errMsg error
	}{ /*
			{
				nameT: "Create new transfer sucessfully",
				args: args{
					accountOriginID:      "001",
					accountDestinationID: "002",
					amount:               30000,
				},
				setup: func(t *testing.T) CreateTransferUC {
					transferRepo := MockCreateTransferTransferRepository{
						SaveErr:   nil,
						transfers: []entities.Transfer{},
					}

					originAccount := entities.NewAccountHelper("001", "69050488013", "Ronaldo Lima", "supersecret", 40000)
					destinationAccount := entities.NewAccountHelper("002", "85570067051", "maria silva", "senha", 3000)

					accountRepo := MockCreateTransferAccountRepository{
						accounts: map[string]entities.Account{
							originAccount.GetID():      originAccount,
							destinationAccount.GetID(): destinationAccount,
						},
						ErrSave: nil,
						ErrGet:  nil,
					}
					return CreateTransferUC{
						transferRepo: &transferRepo,
						accountRepo:  &accountRepo,
					}
				},
				want: func() entities.Transfer {
					transfer := entities.NewTransferHelper(
						"001",
						"002",
						30000)
					return transfer
				}(),
				errMsg: nil,
			},

			{
				nameT: "Origin account not found",
				args: args{
					accountOriginID:      "001",
					accountDestinationID: "002",
					amount:               30000,
				},
				setup: func(t *testing.T) CreateTransferUC {
					transferRepo := MockCreateTransferTransferRepository{
						SaveErr:   nil,
						transfers: []entities.Transfer{},
					}

					originAccount := entities.NewAccountHelper("001", "69050488013", "Ronaldo Lima", "supersecret", 40000)
					destinationAccount := entities.NewAccountHelper("002", "85570067051", "maria silva", "senha", 3000)

					accountRepo := MockCreateTransferAccountRepository{
						accounts: map[string]entities.Account{
							originAccount.GetID():      originAccount,
							destinationAccount.GetID(): destinationAccount,
						},
						ErrSave: nil,
						ErrGet:  errors.New("origin account not found"),
					}

					return CreateTransferUC{
						transferRepo: &transferRepo,
						accountRepo:  &accountRepo,
					}
				},
				want:   entities.Transfer{},
				errMsg: errors.New("origin account not found"),
			},

			{
				nameT: "Destination account not found",
				args: args{
					accountOriginID:      "001",
					accountDestinationID: "002",
					amount:               30000,
				},
				setup: func(t *testing.T) CreateTransferUC {
					transferRepo := MockCreateTransferTransferRepository{
						SaveErr:   errors.New("destination account not found"),
						transfers: []entities.Transfer{},
					}

					originAccount := entities.NewAccountHelper("001", "69050488013", "Ronaldo Lima", "supersecret", 40000)
					destinationAccount := entities.NewAccountHelper("002", "85570067051", "maria silva", "senha", 3000)

					accountRepo := MockCreateTransferAccountRepository{
						accounts: map[string]entities.Account{
							originAccount.GetID():      originAccount,
							destinationAccount.GetID(): destinationAccount,
						},
						ErrSave: nil,
						ErrGet:  nil,
					}

					return CreateTransferUC{
						transferRepo: &transferRepo,
						accountRepo:  &accountRepo,
					}
				},
				want: func() entities.Transfer {
					transfer := entities.NewTransferHelper(
						"001",
						"003",
						30000)
					return transfer
				}(),
				errMsg: errors.New("destination account not found"),
			},*/

		{
			nameT: "Insufficient founds",
			args: args{
				accountOriginID:      "001",
				accountDestinationID: "002",
				amount:               1000,
			},
			setup: func(t *testing.T) CreateTransferUC {
				transferRepo := MockCreateTransferTransferRepository{
					SaveErr:   errors.New("insufficient founds"),
					transfers: []entities.Transfer{},
				}

				originAccount := entities.NewAccountHelper("001", "69050488013", "Ronaldo Lima", "supersecret", 500)
				destinationAccount := entities.NewAccountHelper("002", "85570067051", "maria silva", "senha", 3000)

				accountRepo := MockCreateTransferAccountRepository{
					accounts: map[string]entities.Account{
						originAccount.GetID():      originAccount,
						destinationAccount.GetID(): destinationAccount,
					},
					ErrSave: nil,
					ErrGet:  nil,
				}

				return CreateTransferUC{
					transferRepo: &transferRepo,
					accountRepo:  &accountRepo,
				}
			},
			want: func() entities.Transfer {
				transfer := entities.NewTransferHelper(
					"001",
					"002",
					1000)
				return transfer
			}(),
			errMsg: errors.New("insufficient founds"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.nameT, func(t *testing.T) {
			//transfer, err := entities.NewTransfer(tt.args.accountOriginID, tt.args.accountDestinationID, tt.args.amount)
			transfer, err := tt.setup(t).CreateTransfer(context.Background(), tt.args.accountOriginID, tt.args.accountDestinationID, tt.args.amount)

			if tt.errMsg != nil {
				assert.EqualError(t, err, tt.errMsg.Error())
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, transfer.GetTransferID())
				//assert.Equal(t, tt.want.GetTransferID(), transfer.GetTransferID())
				assert.Equal(t, tt.want.GetAccountOriginID(), transfer.GetAccountOriginID())
				assert.Equal(t, tt.want.GetAccountDestinationID(), transfer.GetAccountDestinationID())
			}
		})
	}
}
