package account

import (
	"time"

	"github.com/julioshinoda/transfer-api/models"
	"github.com/julioshinoda/transfer-api/pkg/database"
)

//AccountsInterface interface that contains account methods
type Accounter interface {
	GetAccounts() []models.Accounts
}

type AccountService struct {
	DB database.SQLInterface
}

//NewAccountService return Accounter interface
func NewAccountService(account AccountService) Accounter {
	return account
}

//GetAccounts list all accounts
func (a AccountService) GetAccounts() []models.Accounts {
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
