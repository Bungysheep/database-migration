package main

import (
	"log"
	"os"

	"github.com/bungysheep/database-migration/pkg/logger"
	"github.com/bungysheep/database-migration/pkg/protocol/db"
	"github.com/bungysheep/database-migration/pkg/service/migration"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	if err := runMigration(); err != nil {
		log.Fatalf("ERROR: %v", err)
		os.Exit(1)
	}
}

func runMigration() error {
	if err := logger.InitLog(); err != nil {
		return err
	}

	logger.Log.Info("Starting database migration\n")

	db, err := db.OpenDbConn()
	if err != nil {
		return err
	}

	migrationService, err := migration.NewMigration(db)
	if err != nil {
		return err
	}

	if err := migrationService.Up(); err != nil {
		return err
	}

	migrationService.Close()

	logger.Log.Info("End of database migration\n")

	return nil
}
