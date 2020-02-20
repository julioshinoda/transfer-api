package transfer

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/julioshinoda/transfer-api/mocks"
	"github.com/stretchr/testify/mock"

	"github.com/julioshinoda/transfer-api/models"
	"github.com/julioshinoda/transfer-api/pkg/database"
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
