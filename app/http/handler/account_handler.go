package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jirlon/digitalbank/app/http/handler/rest"
	"github.com/jirlon/digitalbank/app/usecase"
)

type AccountHandler struct {
	createAccountUC usecase.CreateAccountUseCase
}

type createAccountResponse struct {
	CPF     string `json:"cpf"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func NewAccountHandler(createAccountUC usecase.CreateAccountUseCase) *AccountHandler {
	return &AccountHandler{createAccountUC: createAccountUC}
}

func (h AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) rest.Response {
	var req struct {
		CPF     string `json:"cpf"`
		Name    string `json:"name"`
		Secret  string `json:"secret"`
		Balance int    `json:"balance"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return rest.BadRequest(nil, err)
	}

	account, err := h.createAccountUC.CreateAccount(r.Context(), req.CPF, req.Name, req.Secret, req.Balance)
	if err != nil {
		return rest.InternalServerError(nil, err)
	}

	return rest.Created(createAccountResponse{
		CPF:     account.GetCPF(),
		Name:    account.GetName(),
		Balance: account.GetBalance(),
	})
}
