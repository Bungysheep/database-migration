package migration

import (
	"database/sql"

	"github.com/bungysheep/database-migration/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

// SOURCE_URL - File source url
const SOURCE_URL = "file://./migrations"

// IMigration - Migration service interface
type IMigration interface {
	Up() error
	Down() error
	Close() error
}

type migrationService struct {
	migrate *migrate.Migrate
}

// NewMigration - Migration service implementation
func NewMigration(db *sql.DB) (IMigration, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(SOURCE_URL, config.PGDATABASE, driver)
	if err != nil {
		return nil, err
	}

	return &migrationService{migrate: m}, nil
}

func (ms *migrationService) Up() error {
	if err := ms.migrate.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (ms *migrationService) Down() error {
	if err := ms.migrate.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (ms *migrationService) Close() error {
	sourceErr, dbErr := ms.migrate.Close()

	if sourceErr != nil {
		return sourceErr
	}

	if dbErr != nil {
		return dbErr
	}

	return nil
}
