package test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	TestDataTempDir = filepath.Join(os.TempDir(), "bntp_tests")
	origTestDbPath  = "../../test/data/bntp_test.db"
)

// GetDB opens a copy of the test DB in memory.
func GetDB(t *testing.T) (*sqlx.DB, error) {
	// Create temp dir if needed
	_, err := os.Stat(TestDataTempDir)

	if os.IsNotExist(err) {
		err := os.Mkdir(TestDataTempDir, 0o755)
		if err != nil {
			return nil, err
		}
	}

	dbName := t.Name() + ".db"

	// Copy original database to new temporary one
	dbOrig, err := os.Open(origTestDbPath)
	if err != nil {
		return nil, err
	}

	defer dbOrig.Close()

	tempTestDbPath := filepath.Join(TestDataTempDir, dbName)

	dbNew, err := os.Create(tempTestDbPath)
	if err != nil {
		return nil, err
	}

	defer dbNew.Close()

	_, err = io.Copy(dbNew, dbOrig)
	if err != nil {
		return nil, err
	}

	// Connect to new temporary database
	db, err := sqlx.Open("sqlite3", tempTestDbPath+"?_foreign_keys=1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
