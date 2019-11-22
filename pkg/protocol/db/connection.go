package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bungysheep/database-migration/pkg/config"
)

// OpenDbConn opens connection to postgres database
func OpenDbConn() (*sql.DB, error) {
	log.Printf("Database connection is opening...")

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", config.PGHOST, config.PGPORT, config.PGDATABASE, config.PGUSER, config.PGPASSWORD, config.PGSSLMODE)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
