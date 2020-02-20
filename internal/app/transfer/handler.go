package transfer

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/julioshinoda/transfer-api/models"
	"github.com/julioshinoda/transfer-api/pkg/database/postgres"
	"github.com/julioshinoda/transfer-api/pkg/rest"
)

//GetTransfers function to return accounts list
func GetTransfers(w http.ResponseWriter, r *http.Request) {
	service := NewTransfersService(Service{DB: postgres.GetDBConn()})
	transfersList, err := service.GetTransfers()
	if err != nil {
		rest.RespondwithJSON(w, 500, map[string]string{"message": err.Error()})
		return
	}
	rest.RespondwithJSON(w, 500, transfersList)
}

//CreateTransfer handler to transfer money service
func CreateTransfer(w http.ResponseWriter, r *http.Request) {
	request := &models.TransfersRequest{}
	if err := render.Bind(r, request); err != nil {
		rest.RespondwithJSON(w, 400, map[string]interface{}{"message": err.Error()})
		return
	}

	service := NewTransfersService(Service{DB: postgres.GetDBConn()})

	if err := service.TransferMoney(request.AccountOriginID, request.AccountDestinationID, request.Amount); err != nil {
		rest.RespondwithJSON(w, 400, map[string]string{"message": err.Error()})
	}
	w.WriteHeader(200)
	return
}
