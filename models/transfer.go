package models

import (
	"errors"
	"net/http"
	"time"
)

//Transfers struct data
type Transfers struct {
	ID                   int64     `json:"id"`
	AccountOriginID      int64     `json:"account_origin_id"`
	AccountDestinationID int64     `json:"account_destination_id"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_At"`
}

//TransfersRequest struct with request data to transfer money
type TransfersRequest struct {
	AccountOriginID      int64 `json:"account_origin_id"`
	AccountDestinationID int64 `json:"account_destination_id"`
	Amount               int   `json:"amount"`
}

//Bind contains rules for validate request fields
func (a *TransfersRequest) Bind(r *http.Request) error {
	if a.AccountOriginID == 0 {
		return errors.New("missing field account_origin_id")
	}
	if a.AccountDestinationID == 0 {
		return errors.New("missing field account_destination_id")
	}
	if a.Amount == 0 {
		return errors.New("missing field amount")
	}
	return nil
}
