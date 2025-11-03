package migrations

import (
	"embed"
	"errors"
	"fmt"
	"log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
)

//go:embed sql/*.sql
var sqlFiles embed.FS

type Migrator struct {
	srcDriver source.Driver
}

func NewMigrator() (*Migrator, error) {
	d, err := iofs.New(sqlFiles, "sql")
	if err != nil {
		return nil, fmt.Errorf("unable to create migration source: %v", err)
	}
	return &Migrator{
		srcDriver: d,
	}, nil
}

func MustNewMigrator() *Migrator {
	migrator, err := NewMigrator()
	if err != nil {
		panic(err)
	}
	return migrator
}

func (m *Migrator) ApplyMigrations(db *sqlx.DB) error {
	stdDB := db.DB

	driver, err := postgres.WithInstance(stdDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("unable to create db instance: %v", err)
	}

	migrator, err := migrate.NewWithInstance("iofs", m.srcDriver, "postgres", driver)
	if err != nil {
		return fmt.Errorf("unable to create migration: %v", err)
	}

	defer func() {
		if migrator != nil {
			migrator.Close()
		}
	}()

	if err = migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("unable to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
