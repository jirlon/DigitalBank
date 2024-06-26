package entities

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTransfer(t *testing.T) {
	t.Parallel()

	type args struct {
		accountOriginID      string
		accountDestinationID string
		amount               int
	}

	tests := []struct {
		nameT  string
		args   args
		want   Transfer
		errMsg error
	}{
		{
			nameT: "Create new transfer sucessfully",
			args: args{
				accountOriginID:      "ca5ad981-d471-4172-9a46-8ecf1551307b",
				accountDestinationID: "8b72a4e1-a7b1-473d-be6a-0f20683dc3d7",
				amount:               30000,
			},
			want: Transfer{
				accountOriginID:      "ca5ad981-d471-4172-9a46-8ecf1551307b",
				accountDestinationID: "8b72a4e1-a7b1-473d-be6a-0f20683dc3d7",
				amount:               30000,
			},
			errMsg: nil,
		},

		{
			nameT: "Origin and destination account are the same",
			args: args{
				accountOriginID:      "ca5ad981-d471-4172-9a46-8ecf1551307b",
				accountDestinationID: "ca5ad981-d471-4172-9a46-8ecf1551307b",
				amount:               30000,
			},
			want:   Transfer{},
			errMsg: errors.New("origin and destination account are the same"),
		},

		{
			nameT: "Invalid amount",
			args: args{
				accountOriginID:      "ca5ad981-d471-4172-9a46-8ecf1551307b",
				accountDestinationID: "8b72a4e1-a7b1-473d-be6a-0f20683dc3d7",
				amount:               -1,
			},
			want:   Transfer{},
			errMsg: errors.New("amount should be greater than zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.nameT, func(t *testing.T) {
			transfer, err := NewTransfer(tt.args.accountOriginID, tt.args.accountDestinationID, tt.args.amount)

			if tt.errMsg != nil {
				assert.EqualError(t, err, tt.errMsg.Error())
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, transfer.id)
				transfer.id = ""
				assert.NotEmpty(t, transfer.createdAt)
				transfer.createdAt = time.Time{}
				assert.Equal(t, tt.want, transfer)
			}
		})
	}
}
