package handler

import (
	"net/http"
	"time"

	"github.com/jirlon/digitalbank/app/http/handler/rest"
	"github.com/jirlon/digitalbank/app/usecase"
	"github.com/sirupsen/logrus"
)

type listAccountHandler struct {
	listAccountUC usecase.ListAccountUC
}

type listAccountResponse struct {
	Accounts []accountResponse `json:"accounts"`
}

type accountResponse struct {
	ID        string    `json:"id"`
	CPF       string    `json:"cpf"`
	Name      string    `json:"name"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func (h listAccountHandler) ListAccounts(w http.ResponseWriter, r *http.Request) rest.Response {
	accounts, err := h.listAccountUC.ListAccount()
	if err != nil {
		logrus.Error(err)
		return rest.InternalServerError(nil, err)
	}

	var accountResponses []accountResponse
	for _, account := range accounts {
		accountResponses = append(accountResponses, accountResponse{
			ID:        account.GetID(),
			CPF:       account.GetCPF(),
			Name:      account.GetName(),
			Balance:   account.GetBalance(),
			CreatedAt: account.GetCreatedAt(),
		})
	}

	response := listAccountResponse{
		Accounts: accountResponses,
	}

	return rest.Created(response)
}

func NewListAccountHandler(listAccountUC usecase.ListAccountUC) *listAccountHandler {
	return &listAccountHandler{listAccountUC: listAccountUC}
}
