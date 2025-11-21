package database

import (
	"fmt"
	"igaming-platform/internal/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrationsUp() {
	env := config.LoadEnv()

	dsn := fmt.Sprintf(
		"mysql://%s:%s@tcp(%s:%s)/%s",
		env.DB.Username,
		env.DB.Password,
		env.DB.Host,
		env.DB.Port,
		env.DB.Database,
	)

	migration, err := migrate.New(
		"file://./internal/database/migrations",
		dsn,
	)

	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
