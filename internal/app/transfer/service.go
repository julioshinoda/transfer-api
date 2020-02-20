package transfer

import (
	"errors"
	"time"

	"github.com/julioshinoda/transfer-api/models"
	"github.com/julioshinoda/transfer-api/pkg/database"
)

//Transferer transfer interface that contains all services
type Transferer interface {
	GetTransfers() ([]models.Transfers, error)
	TransferMoney(origin int64, destination int64, value int) error
}

//Service struct that implements Accounter interface
type Service struct {
	DB database.SQLInterface
}

//NewTransfersService return Transferer interface
func NewTransfersService(transfer Service) Transferer {
	return transfer
}

//GetTransfers return all made transfers
func (t Service) GetTransfers() ([]models.Transfers, error) {
	query := "select id,account_origin_id,account_destination_id,amount,created_at from transfers"
	result, err := t.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{}})

	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		var transfersList []models.Transfers
		for _, row := range result {
			transfer := models.Transfers{
				ID:                   row.([]interface{})[0].(int64),
				AccountOriginID:      row.([]interface{})[1].(int64),
				AccountDestinationID: row.([]interface{})[2].(int64),
				Amount:               int(row.([]interface{})[3].(int32)),
				CreatedAt:            row.([]interface{})[4].(time.Time),
			}
			transfersList = append(transfersList, transfer)
		}
		return transfersList, nil
	}
	return nil, nil
}

//TransferMoney service to apply all transfer rules and transfer money between accounts
func (t Service) TransferMoney(origin int64, destination int64, value int) error {
	queryOrigin := "select ballance from accounts where id = $1 and ballance >= $2"
	resultQueryOrigin, err := t.DB.QueryExecutor(database.QueryConfig{QueryStr: queryOrigin,
		Values: []interface{}{origin, value}})

	if err != nil {
		return err
	}
	if len(resultQueryOrigin) == 0 {
		return errors.New("insufficient balance")
	}

	errTransaction := t.DB.TransactionExecutor([]database.QueryConfig{
		database.QueryConfig{
			QueryStr: `update accounts set ballance = (ballance - $1) where id = $2`,
			Values:   []interface{}{value, origin},
		},
		database.QueryConfig{
			QueryStr: `update accounts set ballance = (ballance + $1) where id = $2`,
			Values:   []interface{}{value, destination},
		},
		database.QueryConfig{
			QueryStr: `INSERT INTO transfers
			( account_origin_id, account_destination_id, amount, created_at)
			VALUES($1, $2, $3, $4 );`,
			Values: []interface{}{origin, destination, value, time.Now().Format("2006-01-02")},
		},
	})
	return errTransaction
}
