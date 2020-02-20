package models

import (
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

//Bind contains rules for validate request fields
func (a *Transfers) Bind(r *http.Request) error {

	return nil
}
