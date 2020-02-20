package transfer

import (
	"time"

	"github.com/julioshinoda/transfer-api/models"
	"github.com/julioshinoda/transfer-api/pkg/database"
)

type Transferer interface {
	GetTransfers() ([]models.Transfers, error)
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
