package models

import (
	"errors"
	"net/http"
	"time"
)

//Accounts struct data
type Accounts struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Ballance  int       `json:"ballance"`
	CreatedAt time.Time `json:"created_At"`
}

//Transfers struct data
type Transfers struct {
	ID                   int64     `json:"id"`
	AccountOriginID      int64     `json:"account_origin_id"`
	AccountDestinationID int64     `json:"account_destination_id"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_At"`
}

//Bind contains rules for validate request fields
func (a *Accounts) Bind(r *http.Request) error {

	if a.Name == "" {
		return errors.New("missing field name.")
	}

	if a.CPF == "" {
		return errors.New("missing field CPF.")
	}

	if a.Ballance == 0 {
		return errors.New("missing field ballance.")
	}

	return nil
}
