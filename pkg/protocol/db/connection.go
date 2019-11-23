package db

import (
	"database/sql"

	"github.com/bungysheep/database-migration/pkg/config"
	"github.com/bungysheep/database-migration/pkg/logger"
)

// OpenDbConn opens connection to postgres database
func OpenDbConn() (*sql.DB, error) {
	logger.Log.Info("Openning database\n")

	db, err := sql.Open("postgres", config.CONNECTIONSTRING)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
