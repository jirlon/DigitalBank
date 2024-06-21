package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jirlon/digitalbank/app/http/handler/rest"
	"github.com/jirlon/digitalbank/app/usecase"
	"github.com/sirupsen/logrus"
)

type getBalanceHandler struct {
	getBalanceUC usecase.GetBalanceUC
}

type getBalanceResponse struct {
	Balance int `json:"balance"`
}

func (h getBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) rest.Response {
	id := chi.URLParam(r, "account_id")
	if id == "" {
		return rest.BadRequest(nil, errors.New("missing account_id parameter"))
	}

	balance, err := h.getBalanceUC.GetBalance(id)
	if err != nil {
		logrus.Error(err)
		return rest.InternalServerError(nil, err)
	}

	response := getBalanceResponse{Balance: balance}

	return rest.OK(response)
}

func NewGetBalanceHandler(getBalanceUC usecase.GetBalanceUC) *getBalanceHandler {
	return &getBalanceHandler{getBalanceUC: getBalanceUC}
}
