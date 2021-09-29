package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/libbookmarks"
	"github.com/stretchr/testify/assert"
)

// ######################
// # ImportMinimalCSV() #
// ######################
func TestImportMinimalCSVEmpty(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := ""
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVNoHeaderButData(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := "Foo;Bar"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVHeaderNoTitle(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := "dss;Title"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVHeaderNoUrl(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := "dss;Url"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVOnlyHeader(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := "Url;Title"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVOneEntry(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)
}

func TestImportMinimalCSVManyEntries(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar
Foo2;Bar2
Foo3;Bar3`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)
}

func TestImportMinimalCSVEntryWithIncompleteUrl(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar
Foo2;Bar2
;Bar3`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVEntryWithIncompleteTitle(t *testing.T) {
	db, err := GetDB(t)
	assert.NoError(t, err)

	file, err := os.Create(filepath.Join(testDataTempDir, t.Name()))
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar
Foo2;Bar2
Foo3;`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(testDataTempDir, t.Name()))
	assert.Error(t, err)
}
