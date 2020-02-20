package transfer

import (
	"net/http"

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
