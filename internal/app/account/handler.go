package account

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/julioshinoda/transfer-api/pkg/database/postgres"
)

//GetAccounts function to return accounts list
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accountService := NewAccountService(AccountService{DB: postgres.GetDBConn()})
	accountService.GetAccounts()
	render.JSON(w, r, accountService.GetAccounts())
}
