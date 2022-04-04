package test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	_ "embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var TestDataTempDir = filepath.Join(os.TempDir(), "bntp_tests")

//go:embed ../bntp.sql
var sqlSchemaEmbed string

// GetDB opens a copy of the test DB in memory.
func GetDB(t *testing.T) (*sqlx.DB, error) {
	path := filepath.Join("..", "..", "bntp.sql")

	schema, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	schemaCommand := string(schema)

	// Connect to new temporary database
	db, err := sqlx.Open("sqlite3", ":memory:?_foreign_keys=1")
	if err != nil {
		return nil, err
	}

	// Load schema
	_, err = db.Exec(schemaCommand)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTestTempFile(filename string) (*os.File, error) {
	err := os.MkdirAll(TestDataTempDir, os.ModeDir|os.ModePerm)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(filepath.Join(TestDataTempDir, filename))
	if err != nil {
		return nil, err
	}

	return file, nil
}
