package migration

import (
	"database/sql"
	"fmt"

	"github.com/bungysheep/database-migration/pkg/config"
	"github.com/bungysheep/database-migration/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

// SOURCEURL - File source url
const SOURCEURL = "file://./migrations"

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
	logger.Log.Info("Creating postgres driver\n")

	driver, err := postgres.WithInstance(db, &postgres.Config{DatabaseName: config.PGDATABASE})
	if err != nil {
		return nil, err
	}

	logger.Log.Info("Opening source\n")

	m, err := migrate.NewWithDatabaseInstance(SOURCEURL, config.PGDATABASE, driver)
	if err != nil {
		return nil, err
	}

	m.Log = &migrationLog{}

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

type migrationLog struct{}

func (ml *migrationLog) Printf(format string, args ...interface{}) {
	logger.Log.Info(fmt.Sprintf(format, args...))
}

func (ml *migrationLog) Verbose() bool {
	return true
}
