package main

import (
	"fmt"
	"os"

	"github.com/julioshinoda/transfer-api/tools/migration"
)

func main() {
	err := migration.Up(os.Getenv("MIGRATION_FOLDER"), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("migration", "Migration failed", err.Error())
	}

}
