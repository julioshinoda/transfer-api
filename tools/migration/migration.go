package migration

import (
	"errors"
	"log"
	"os/exec"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//Create  method to create migrations up and down files.
func Create(migrationFolder string, migrationTitle string) error {
	if migrationFolder == "" {
		return errors.New("Invalid migration folder")
	}
	if migrationTitle == "" {
		return errors.New("Invalid migration title")
	}
	exec.Command("mkdir", "-p", migrationFolder).Run()
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", migrationFolder, "-seq", migrationTitle)
	return cmd.Run()
}

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
