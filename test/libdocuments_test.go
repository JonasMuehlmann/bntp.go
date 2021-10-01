package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/libdocuments"
	"github.com/stretchr/testify/assert"
)

// ############
// # AddTag() #
// ############
func TestAddTagToDocumentEmpty(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTag(filePath, "Foo")
	assert.Error(t, err)
}

func TestAddTagToDocumentNoTag(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTag(filePath, "")
	assert.Error(t, err)
}

func TestAddTagToDocument(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTag(filePath, "Foo")
	assert.NoError(t, err)
}

func TestAddTagToDocumentTwice(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTag(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.AddTag(filePath, "Bar")
	assert.NoError(t, err)
}
