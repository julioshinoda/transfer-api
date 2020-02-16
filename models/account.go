package models

import "time"

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
