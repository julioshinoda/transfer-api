package migration

import (
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestCreate(t *testing.T) {
	type args struct {
		migrationFolder string
		migrationTitle  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				migrationFolder: "test/folder",
				migrationTitle:  "Migration_success",
			},
			wantErr: false,
		},
		{
			name: "Fail with invalid folder",
			args: args{
				migrationFolder: "",
				migrationTitle:  "Migration_success",
			},
			wantErr: true,
		},
		{
			name: "Fail with invalid folder",
			args: args{
				migrationFolder: "test/folder",
				migrationTitle:  "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.migrationFolder, tt.args.migrationTitle); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUp(t *testing.T) {
	type args struct {
		migrationFolder string
		connectionURL   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Fail with invalid folder",
			args: args{
				migrationFolder: "",
				connectionURL:   "postgres//user:pass@host?sslMode=disable",
			},
			wantErr: true,
		},
		{
			name: "Fail with invalid url connection",
			args: args{
				migrationFolder: "test/folder",
				connectionURL:   "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Up(tt.args.migrationFolder, tt.args.connectionURL)
		})
	}
}
