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

func (h getBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) rest.Response {
	//vars := mux.Vars(r)
	//id := vars["account_id"]
	//id := r.URL.Query().Get("account_id")
	id := chi.URLParam(r, "account_id")
	if id == "" {
		return rest.BadRequest(nil, errors.New("missing id parameter"))
	}

	balance, err := h.getBalanceUC.GetBalance(id)
	if err != nil {
		logrus.Error(err)
		return rest.InternalServerError(nil, err)
	}

	response := map[string]int{"balance": balance}

	return rest.Created(response)
}

func NewGetBalanceHandler(getBalanceUC usecase.GetBalanceUC) *getBalanceHandler {
	return &getBalanceHandler{getBalanceUC: getBalanceUC}
}
