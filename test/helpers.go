package test

import (
	"bufio"
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "embed"

	"github.com/JonasMuehlmann/goaoi"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var TestDataTempDir = filepath.Join(os.TempDir(), "bntp_tests")

// GetDB opens a copy of the test DB in memory.
func GetDB() (*sql.DB, error) {
	// Connect to new temporary database
	// db, err := sql.Open("sqlite3", ":memory:?_foreign_keys=1")
	// db, err := sql.Open("sqlite3", "file::memory:?_foreign_keys=1")
	db, err := sql.Open("sqlite3", "file:"+uuid.NewString()+"?mode=memory&cache=shared&_foreign_keys=1")
	if err != nil {
		return nil, err
	}
	// db.SetMaxIdleConns(1)
	// db.SetMaxOpenConns(1)

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	parentDirs := strings.Split(cwd, string(os.PathSeparator))

	iProjectRoot, err := goaoi.FindIfSlice(parentDirs, goaoi.AreEqualPartial("bntp.go"))
	if err != nil {
		return nil, err
	}

	schemaFilePath := string(os.PathSeparator) + filepath.Join(filepath.Join(parentDirs[:iProjectRoot+1]...), "schema", "bntp_sqlite.sql")

	sqlSchema, err := ioutil.ReadFile(schemaFilePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(sqlSchema))
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

func InterceptStdout(t *testing.T) (*bufio.Scanner, *os.File, *os.File) {
	reader, writer, err := os.Pipe()
	if err != nil {
		assert.Fail(t, "Error creating pipe: %v", err)
	}

	os.Stdout = writer

	return bufio.NewScanner(reader), reader, writer
}

func ResetStdout(t *testing.T, reader *os.File, writer *os.File) {
	err := reader.Close()
	if err != nil {
		assert.Fail(t, "Error closing reader: %v", err)
	}

	err = writer.Close()
	if err != nil {
		assert.Fail(t, "Error closing writer: %v", err)
	}

	os.Stdout = os.Stderr
}

func HandlePanic(t *testing.T, name string) {
	if err := recover(); err != nil {
		assert.Fail(t, name, err)
	}
}
