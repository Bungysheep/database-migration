package main

import (
	"log"
	"os"

	"github.com/bungysheep/database-migration/pkg/protocol/db"
	_ "github.com/lib/pq"
)

func main() {
	if err := migrate(); err != nil {
		log.Fatalf("ERROR: %v", err)
		os.Exit(1)
	}
}

func migrate() error {
	_, err := db.OpenDbConn()
	if err != nil {
		return err
	}

	return nil
}
