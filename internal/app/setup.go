package app

/*
	Here is set application setup like middlewares and routes.
*/

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/julioshinoda/transfer-api/internal/app/account"
	"github.com/julioshinoda/transfer-api/internal/app/transfer"
)

//Setup function that start server with all setups
func Setup() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Group(func(r chi.Router) {
		//Accounts routes
		r.Get("/accounts", account.GetAccounts)
		r.Get("/accounts/{account_id}/balance", account.GetBallance)
		r.Post("/accounts", account.CreateAccount)

		//Transfers routes

		r.Get("/transfers", transfer.GetTransfers)
		r.Post("/transfers", transfer.CreateTransfer)
	})
	return r
}
