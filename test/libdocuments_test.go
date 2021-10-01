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

// ###############
// # RemoveTag() #
// ###############
func TestRemoveTagFromDocumentEmpty(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveTag(filePath, "Foo")
	assert.Error(t, err)
}

func TestRemoveTagFromDocumentNoTag(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveTag(filePath, "")
	assert.Error(t, err)
}

func TestRemoveTagFromDocument(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTag(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.RemoveTag(filePath, "Foo")
	assert.NoError(t, err)
}

func TestRemoveTagFromDocumentTwice(t *testing.T) {
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

	err = libdocuments.RemoveTag(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.RemoveTag(filePath, "Bar")
	assert.NoError(t, err)
}

// ###############
// # RenameTag() #
// ###############
func TestRenameTagFromDocumentEmpty(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RenameTag(filePath, "Foo", "Bar")
	assert.Error(t, err)
}

func TestRenameTagFromDocumentNoOldTag(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RenameTag(filePath, "", "Bar")
	assert.Error(t, err)
}

func TestRenameTagFromDocumentNoNewTag(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RenameTag(filePath, "Foo", "")
	assert.Error(t, err)
}

func TestRenameTagFromDocument(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTag(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.RenameTag(filePath, "Foo", "Bar")
	assert.NoError(t, err)
}

func TestRenameTagFromDocumentTwice(t *testing.T) {
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

	err = libdocuments.RenameTag(filePath, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.RenameTag(filePath, "Bar", "Foo")
	assert.NoError(t, err)
}

// #############
// # GetTags() #
// #############
func TestGetTagsEmpty(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, err = libdocuments.GetTags(filePath)
	assert.Error(t, err)
}

func TestGetTagsNoTags(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, err = libdocuments.GetTags(filePath)
	assert.Error(t, err)
}

func TestGetTagsOneTag(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTag(filePath, "Foo")
	assert.NoError(t, err)

	tags, err := libdocuments.GetTags(filePath)
	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, "Foo", tags[0])

}

func TestGetTagsManyTags(t *testing.T) {
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

	err = libdocuments.AddTag(filePath, "Baz")
	assert.NoError(t, err)

	tags, err := libdocuments.GetTags(filePath)
	assert.NoError(t, err)
	assert.Len(t, tags, 3)
	assert.Equal(t, []string{"Foo", "Bar", "Baz"}, tags)

}

// ##################
// # FindTagsLine() #
// ##################
func TestFindTagsLineEmpty(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, _, err = libdocuments.FindTagsLine(filePath)
	assert.Error(t, err)
}

func TestFindTagsLineFirst(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	i, _, err := libdocuments.FindTagsLine(filePath)
	assert.NoError(t, err)
	assert.Equal(t, i, 1)
}

func TestFindTagsLineLast(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `


# Tags
`
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	i, _, err := libdocuments.FindTagsLine(filePath)
	assert.NoError(t, err)
	assert.Equal(t, i, 4)
}

func TestFindTagsLineMiddle(t *testing.T) {
	filePath := filepath.Join(testDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `
# Tags

`
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	i, _, err := libdocuments.FindTagsLine(filePath)
	assert.NoError(t, err)
	assert.Equal(t, i, 2)
}
