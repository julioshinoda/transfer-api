.EXPORT_ALL_VARIABLES:

PORT=9011
DATABASE_URL=postgres://transfer:transfer@localhost:5433/transfer?sslmode=disable

MIGRATION_FOLDER=db/migrations

run: 
	CompileDaemon --build="go build cmd/rest/server.go" --command=./server

migration:
	go run cmd/migration/migration.go

test :
	go clean -testcache
	go test ./... -race -coverprofile cp.out
	go tool cover -html=./cp.out -o cover.html

migration-create :
	go run tools/migration/create/create_migration.go -title=$(TITLE)

