package account

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/julioshinoda/transfer-api/mocks"
	"github.com/julioshinoda/transfer-api/models"
	"github.com/julioshinoda/transfer-api/pkg/database"
	"github.com/stretchr/testify/mock"
)

func TestAccountService_GetAccounts(t *testing.T) {
	acc1 := models.Accounts{
		ID:        int64(1),
		Name:      "account1",
		CPF:       "12312312300",
		Ballance:  13000,
		CreatedAt: time.Now(),
	}
	type fields struct {
		DB database.SQLInterface
	}
	tests := []struct {
		name   string
		fields fields
		want   []models.Accounts
	}{
		{
			name: "success to return accounts",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			want: []models.Accounts{
				acc1,
			},
		},
		{
			name: "error to execute query accounts",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			want: []models.Accounts{},
		},
		{
			name: "not found any accounts",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			want: []models.Accounts{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Service{
				DB: tt.fields.DB,
			}
			switch tt.name {
			case "success to return accounts":
				a.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select id, name, cpf, ballance, created_at from accounts"
				})).Return([]interface{}{[]interface{}{int64(1), "account1", "12312312300", int32(13000), acc1.CreatedAt}}, nil)
			case "error to execute query accounts":
				a.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select id, name, cpf, ballance, created_at from accounts"
				})).Return([]interface{}{}, errors.New("some error"))
			case "not found any accounts":
				a.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select id, name, cpf, ballance, created_at from accounts"
				})).Return([]interface{}{}, nil)
			}
			if got := a.GetAccounts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountService.GetAccounts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetBallanceByAccountsID(t *testing.T) {
	type fields struct {
		DB database.SQLInterface
	}
	type args struct {
		accountID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success to get ballance",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			args: args{
				accountID: int64(1),
			},
			want: int(12025),
		},
		{
			name: "error to get ballance",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			args: args{
				accountID: int64(1),
			},
			want:    int(0),
			wantErr: true,
		},
		{
			name: "not found ballance",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			args: args{
				accountID: int64(1),
			},
			want:    int(0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Service{
				DB: tt.fields.DB,
			}
			switch tt.name {
			case "success to get ballance":
				a.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select  ballance from accounts where id = $1"
				})).Return([]interface{}{[]interface{}{int32(12025)}}, nil)
			case "error to get ballance":
				a.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select  ballance from accounts where id = $1"
				})).Return([]interface{}{[]interface{}{int32(0)}}, errors.New("error_on_query"))
			case "not found ballance":
				a.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select  ballance from accounts where id = $1"
				})).Return([]interface{}{}, nil)
			}
			got, err := a.GetBallanceByAccountsID(tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetBallanceByAccountsID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.GetBallanceByAccountsID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateAccount(t *testing.T) {
	type fields struct {
		DB database.SQLInterface
	}
	type args struct {
		account models.Accounts
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success to create account",
			fields:  fields{DB: &mocks.SQLInterface{}},
			args:    args{account: models.Accounts{Name: "test account", CPF: "12312312345", Ballance: 150000}},
			wantErr: false,
		},
		{
			name:    "error on create account",
			fields:  fields{DB: &mocks.SQLInterface{}},
			args:    args{account: models.Accounts{Name: "test account", CPF: "12312312345", Ballance: 150000}},
			wantErr: true,
		},
		{
			name:    "error for duplicated CPF",
			fields:  fields{DB: &mocks.SQLInterface{}},
			args:    args{account: models.Accounts{Name: "test account", CPF: "12312312345", Ballance: 150000}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Service{
				DB: tt.fields.DB,
			}
			switch tt.name {
			case "success to create account":
				a.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == `INSERT INTO accounts( "name", cpf, ballance, created_at) VALUES($1, $2, $3, $4) RETURNING id;`
				})).Return([]interface{}{[]interface{}{int32(0)}}, nil)
			case "error on create account":
				a.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == `INSERT INTO accounts( "name", cpf, ballance, created_at) VALUES($1, $2, $3, $4) RETURNING id;`
				})).Return([]interface{}{}, errors.New("error"))
			case "error for duplicated CPF":
				a.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == `INSERT INTO accounts( "name", cpf, ballance, created_at) VALUES($1, $2, $3, $4) RETURNING id;`
				})).Return([]interface{}{}, nil)
			}

			if err := a.CreateAccount(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
