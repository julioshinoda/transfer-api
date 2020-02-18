package account

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/julioshinoda/transfer-api/pkg/database/postgres"
)

//GetAccounts function to return accounts list
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accountService := NewAccountService(Service{DB: postgres.GetDBConn()})
	accountService.GetAccounts()
	render.JSON(w, r, accountService.GetAccounts())
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

	w.WriteHeader(200)
	render.JSON(w, r, map[string]float64{"ballance": ballance})
	return
}
