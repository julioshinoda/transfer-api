package account

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/julioshinoda/transfer-api/models"
	"github.com/julioshinoda/transfer-api/pkg/database/postgres"
	"github.com/julioshinoda/transfer-api/pkg/rest"
)

//GetAccounts function to return accounts list
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accountService := NewAccountService(Service{DB: postgres.GetDBConn()})
	accountService.GetAccounts()
	rest.RespondwithJSON(w, 200, accountService.GetAccounts())
}

//GetBallance return ballance from accounts ID
func GetBallance(w http.ResponseWriter, r *http.Request) {
	reqAccountID := chi.URLParam(r, "account_id")
	accountID, err := strconv.Atoi(reqAccountID)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	accountService := NewAccountService(Service{DB: postgres.GetDBConn()})
	ballance, err := accountService.GetBallanceByAccountsID(int64(accountID))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	rest.RespondwithJSON(w, 200, map[string]int{"ballance": ballance})
	return
}

//CreateBallance handler to bind request and send to create service
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	accountReq := &models.Accounts{}
	if err := render.Bind(r, accountReq); err != nil {
		rest.RespondwithJSON(w, 400, map[string]interface{}{"message": err.Error()})
		return
	}
	accountService := NewAccountService(Service{DB: postgres.GetDBConn()})
	err := accountService.CreateAccount(*accountReq)
	if err != nil {
		rest.RespondwithJSON(w, 400, map[string]string{"message": err.Error()})
		return
	}
	w.WriteHeader(201)
	return
}
