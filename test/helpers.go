package test

import (
	"bufio"
	"bytes"
	"database/sql"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	_ "embed"

	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/goaoi/functional"
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

	iProjectRoot, err := goaoi.FindIfSlice(parentDirs, functional.AreEqualPartial("bntp.go"))
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

type OutputValidator func(t *testing.T, output string, message string) bool

func ValidatorEqual(expected string) OutputValidator {
	return func(t *testing.T, output string, message string) bool {
		return assert.Equal(t, expected, output, message)
	}
}

func ValidatorEmpty(t *testing.T, actual string, message string) bool {
	return assert.Empty(t, actual, message)
}

func ValidatorContains(mustHaves ...string) OutputValidator {
	return func(t *testing.T, actual string, message string) bool {
		for _, mustHave := range mustHaves {
			if !assert.Contains(t, actual, mustHave, message) {
				return false
			}
		}

		return true
	}
}

// zupa@https://stackoverflow.com/a/36226525/14977798
type Buffer struct {
	b bytes.Buffer
	m sync.Mutex
}

func NewBufferString(s string) *Buffer {
	return &Buffer{b: *bytes.NewBufferString(s)}
}

func (b *Buffer) Read(p []byte) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Read(p)
}
func (b *Buffer) Write(p []byte) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Write(p)
}
func (b *Buffer) String() string {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.String()
}

func (b *Buffer) Bytes() []byte {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Bytes()
}
func (b *Buffer) Cap() int {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Cap()
}
func (b *Buffer) Grow(n int) {
	b.m.Lock()
	defer b.m.Unlock()
	b.b.Grow(n)
}
func (b *Buffer) Len() int {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Len()
}
func (b *Buffer) Next(n int) []byte {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.Next(n)
}
func (b *Buffer) ReadByte() (c byte, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.ReadByte()
}
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.ReadBytes(delim)
}
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.ReadFrom(r)
}
func (b *Buffer) ReadRune() (r rune, size int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.ReadRune()
}
func (b *Buffer) ReadString(delim byte) (line string, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.ReadString(delim)
}
func (b *Buffer) Reset() {
	b.m.Lock()
	defer b.m.Unlock()
	b.b.Reset()
}
func (b *Buffer) Truncate(n int) {
	b.m.Lock()
	defer b.m.Unlock()
	b.b.Truncate(n)
}
func (b *Buffer) UnreadByte() error {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.UnreadByte()
}
func (b *Buffer) UnreadRune() error {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.UnreadRune()
}
func (b *Buffer) WriteByte(c byte) error {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.WriteByte(c)
}
func (b *Buffer) WriteRune(r rune) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.WriteRune(r)
}
func (b *Buffer) WriteString(s string) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.WriteString(s)
}
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.WriteTo(w)
}
