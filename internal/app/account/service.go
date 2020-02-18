package account

import (
	"time"

	"github.com/julioshinoda/transfer-api/models"
	"github.com/julioshinoda/transfer-api/pkg/database"
)

//Accounter interface that contains account methods
type Accounter interface {
	GetAccounts() []models.Accounts
	GetBallanceByAccountsID(accountID int64) (float64, error)
}

//Service struct that implements Accounter interface
type Service struct {
	DB database.SQLInterface
}

//NewAccountService return Accounter interface
func NewAccountService(account Service) Accounter {
	return account
}

//GetAccounts list all accounts
func (a Service) GetAccounts() []models.Accounts {
	query := "select id, name, cpf, ballance, created_at from accounts"
	result, err := a.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{}})

	if err != nil {
		return []models.Accounts{}
	}
	accounts := []models.Accounts{}
	if len(result) > 0 {
		for _, row := range result {
			account := models.Accounts{
				ID:        row.([]interface{})[0].(int64),
				Name:      row.([]interface{})[1].(string),
				CPF:       row.([]interface{})[2].(string),
				Ballance:  int(row.([]interface{})[3].(int32)),
				CreatedAt: row.([]interface{})[4].(time.Time),
			}
			accounts = append(accounts, account)
		}

		return accounts
	}

	return []models.Accounts{}
}

//GetBallanceByAccountsID returns ballance from an account by a given account ID
func (a Service) GetBallanceByAccountsID(accountID int64) (float64, error) {
	query := "select  ballance from accounts where id = $1"
	result, err := a.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{accountID}})

	if err != nil {
		return 0, err
	}

	if len(result) > 0 {
		account := result[0]
		return float64(account.([]interface{})[0].(int32)) / 100, nil
	}

	return 0, nil

}
