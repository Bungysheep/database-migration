package main

import (
	"log"
	"os"

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
	log.Printf("Open migration...")

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

	log.Printf("Close migration...")
	migrationService.Close()

	return nil
}
