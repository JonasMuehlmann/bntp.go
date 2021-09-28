package tests

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// GetDB opens a copy of the test DB in memory.
func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:data/bntp_test.db?mode=memory")
	if err != nil {
		log.Fatal("Could not load database")
	}

	return db, nil
}
