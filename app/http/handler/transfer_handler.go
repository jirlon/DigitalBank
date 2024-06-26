package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jirlon/digitalbank/app/http/handler/rest"
	"github.com/jirlon/digitalbank/app/usecase"
	"github.com/sirupsen/logrus"
)

type transferHandler struct {
	createTransferUC usecase.CreateTransferUC
}

type transferRequest struct {
	OriginID      string `json:"origin_id"`
	DestinationID string `json:"destination_id"`
	Amount        int    `json:"amount"`
}

type transferResponse struct {
	TransferID           string    `json:"transfer_id"`
	AccountOriginID      string    `json:"account_origin_id"`
	AccountDestinationID string    `json:"account_destination_id"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

func (h transferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) rest.Response {
	var req transferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logrus.Error(err)
		return rest.BadRequest(nil, err)
	}

	transfer, err := h.createTransferUC.CreateTransfer(r.Context(), req.OriginID, req.DestinationID, req.Amount)
	if err != nil {
		return rest.InternalServerError(nil, err)
	}

	response := transferResponse{
		TransferID:           transfer.GetTransferID(),
		AccountOriginID:      transfer.GetAccountOriginID(),
		AccountDestinationID: transfer.GetAccountDestinationID(),
		Amount:               transfer.GetAmount(),
		CreatedAt:            transfer.GetCreatedAt(),
	}
	return rest.OK(response)
}

func NewCreateTransferHamdler(createTransferUC usecase.CreateTransferUC) *transferHandler {
	return &transferHandler{createTransferUC: createTransferUC}
}
