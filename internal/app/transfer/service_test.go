package transfer

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

func TestService_GetTransfers(t *testing.T) {
	now := time.Now()
	transfer := models.Transfers{
		ID:                   int64(1),
		AccountDestinationID: int64(2),
		AccountOriginID:      int64(1),
		Amount:               1250,
		CreatedAt:            now,
	}
	type fields struct {
		DB database.SQLInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Transfers
		wantErr bool
	}{
		{
			name: "success to get transfers list",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			want: []models.Transfers{transfer},
		},
		{
			name: "error to get transfers list",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not found any transfer",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transfer := Service{
				DB: tt.fields.DB,
			}
			switch tt.name {
			case "success to get transfers list":
				transfer.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select id,account_origin_id,account_destination_id,amount,created_at from transfers"
				})).Return([]interface{}{[]interface{}{int64(1), int64(1), int64(2), int32(1250), now}}, nil)
			case "error to get transfers list":
				transfer.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select id,account_origin_id,account_destination_id,amount,created_at from transfers"
				})).Return(nil, errors.New("error"))
			case "not found any transfer":
				transfer.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select id,account_origin_id,account_destination_id,amount,created_at from transfers"
				})).Return(nil, nil)
			}
			got, err := transfer.GetTransfers()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetTransfers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_TransferMoney(t *testing.T) {
	type fields struct {
		DB database.SQLInterface
	}
	type args struct {
		origin      int64
		destination int64
		value       int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success to transfer money",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			args: args{
				origin:      int64(1),
				destination: int64(2),
				value:       1200,
			},
		},
		{
			name: "error on query origin accounts balance",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			args: args{
				origin:      int64(1),
				destination: int64(2),
				value:       10000,
			},
			wantErr: true,
		},
		{
			name: "insufficient balance",
			fields: fields{
				DB: &mocks.SQLInterface{},
			},
			args: args{
				origin:      int64(1),
				destination: int64(2),
				value:       10000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transfer := Service{
				DB: tt.fields.DB,
			}
			switch tt.name {
			case "success to transfer money":
				transfer.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select ballance from accounts where id = $1 and ballance >= $2"
				})).Return([]interface{}{[]interface{}{int32(15000)}}, nil)
				transfer.DB.(*mocks.SQLInterface).On("TransactionExecutor", mock.Anything).Return(nil)
			case "insufficient balance":
				transfer.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select ballance from accounts where id = $1 and ballance >= $2"
				})).Return([]interface{}{}, nil)
			case "error on query origin accounts balance":
				transfer.DB.(*mocks.SQLInterface).On("QueryExecutor", mock.MatchedBy(func(conn database.QueryConfig) bool {
					return conn.QueryStr == "select ballance from accounts where id = $1 and ballance >= $2"
				})).Return(nil, errors.New("error"))
			}

			if err := transfer.TransferMoney(tt.args.origin, tt.args.destination, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Service.TransferMoney() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
