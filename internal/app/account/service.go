package account

import (
	"errors"
	"time"

	"github.com/julioshinoda/transfer-api/models"
	"github.com/julioshinoda/transfer-api/pkg/database"
)

//Accounter interface that contains account methods
type Accounter interface {
	GetAccounts() []models.Accounts
	GetBallanceByAccountsID(accountID int64) (int, error)
	CreateAccount(account models.Accounts) error
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
func (a Service) GetBallanceByAccountsID(accountID int64) (int, error) {
	query := "select  ballance from accounts where id = $1"
	result, err := a.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{accountID}})

	if err != nil {
		return 0, err
	}

	if len(result) > 0 {
		account := result[0]
		return int(account.([]interface{})[0].(int32)), nil
	}

	return 0, nil

}

func (a Service) CreateAccount(account models.Accounts) error {
	query := `INSERT INTO accounts( "name", cpf, ballance, created_at) VALUES($1, $2, $3, $4) RETURNING id;`
	result, err := a.DB.QueryExecutor(database.QueryConfig{QueryStr: query,
		Values: []interface{}{account.Name, account.CPF, account.Ballance, time.Now().Format("2006-01-02")}})

	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("account already created")
	}

	return nil
}
