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
			a := AccountService{
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
