package migration

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//Up method to update to last database migration
func Up(migrationFolder string, connectionURL string) error {
	if migrationFolder == "" {
		return errors.New("Invalid migration folder")
	}
	if connectionURL == "" {
		return errors.New("Invalid migration connectionURL")
	}
	m, err := migrate.New(
		"file://"+migrationFolder,
		connectionURL)

	if err != nil {
		log.Println(err.Error())
	}
	if err := m.Up(); err != nil {
		log.Println(err.Error())
	}
	return nil
}
