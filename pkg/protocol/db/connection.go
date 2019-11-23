package db

import (
	"database/sql"
	"fmt"

	"github.com/bungysheep/database-migration/pkg/config"
	"github.com/bungysheep/database-migration/pkg/logger"
)

// OpenDbConn opens connection to postgres database
func OpenDbConn() (*sql.DB, error) {
	logger.Log.Info("Openning database\n")

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
