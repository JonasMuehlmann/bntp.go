package libdocuments_test

import (
	"os"
	"path/filepath"

	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/libdocuments"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
)

// ##################
// # AddTagToFile() #
// ##################
func TestAddTagToFileEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Foo")
	assert.Error(t, err)
}

func TestAddTagToFileNoTag(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "")
	assert.Error(t, err)
}

func TestAddTagToFile(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Foo")
	assert.NoError(t, err)
}

func TestAddTagToFileTwice(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Bar")
	assert.NoError(t, err)
}

// #######################
// # RemoveTagFromFile() #
// #######################
func TestRemoveTagFromFileEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveTagFromFile(filePath, "Foo")
	assert.Error(t, err)
}

func TestRemoveTagFromFileNoTag(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveTagFromFile(filePath, "")
	assert.Error(t, err)
}

func TestRemoveTagFromFile(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.RemoveTagFromFile(filePath, "Foo")
	assert.NoError(t, err)
}

func TestRemoveTagFromFileTwice(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Bar")
	assert.NoError(t, err)

	err = libdocuments.RemoveTagFromFile(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.RemoveTagFromFile(filePath, "Bar")
	assert.NoError(t, err)
}

// #####################
// # RenameTagInFile() #
// #####################
func TestRenameTagInFileEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RenameTagInFile(filePath, "Foo", "Bar")
	assert.Error(t, err)
}

func TestRenameTagInFileNoOldTag(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RenameTagInFile(filePath, "", "Bar")
	assert.Error(t, err)
}

func TestRenameTagInFileNoNewTag(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RenameTagInFile(filePath, "Foo", "")
	assert.Error(t, err)
}

func TestRenameTagInFile(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.RenameTagInFile(filePath, "Foo", "Bar")
	assert.NoError(t, err)
}

func TestRenameTagInFileTwice(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Bar")
	assert.NoError(t, err)

	err = libdocuments.RenameTagInFile(filePath, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.RenameTagInFile(filePath, "Bar", "Foo")
	assert.NoError(t, err)
}

// #############
// # GetTags() #
// #############
func TestGetTagsEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, err = libdocuments.GetTags(filePath)
	assert.Error(t, err)
}

func TestGetTagsNoTags(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, err = libdocuments.GetTags(filePath)
	assert.Error(t, err)
}

func TestGetTagsOneTag(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Foo")
	assert.NoError(t, err)

	tags, err := libdocuments.GetTags(filePath)
	assert.NoError(t, err)
	assert.Len(t, tags, 1)
	assert.Equal(t, "Foo", tags[0])

}

func TestGetTagsManyTags(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Foo")
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddTagToFile(filePath, "Baz")
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
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, _, err = libdocuments.FindTagsLine(filePath)
	assert.Error(t, err)
}

func TestFindTagsLineFirst(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

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
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

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
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

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

// #############
// # HasTags() #
// #############
func TestHasTagsEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, err = libdocuments.HasTags(filePath, []string{"Foo"})
	assert.Error(t, err)
}

func TestHasTagsNoTags(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Tags"

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, err = libdocuments.HasTags(filePath, []string{"Foo"})
	assert.Error(t, err)
}

func TestHasTagsNoInput(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Tags
Foo
    `

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, err = libdocuments.HasTags(filePath, []string{})
	assert.Error(t, err)
}

func TestHasTagsEmptyStringInput(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Tags
Foo
    `

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, err = libdocuments.HasTags(filePath, []string{""})
	assert.Error(t, err)
}

func TestHasTagsNotFound(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Tags
Foo
    `

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	hasTag, err := libdocuments.HasTags(filePath, []string{"Bar"})
	assert.NoError(t, err)
	assert.False(t, hasTag)
}

func TestHasTags(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Tags
Foo
    `
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	hasTag, err := libdocuments.HasTags(filePath, []string{"Foo"})
	assert.NoError(t, err)
	assert.True(t, hasTag)
}

func TestHasTagsMultipleInDocument(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Tags
Foo,Bar
    `
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	hasTag, err := libdocuments.HasTags(filePath, []string{"Foo"})
	assert.NoError(t, err)
	assert.True(t, hasTag)
}

func TestHasTagsNotFoundMultipleInInput(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Tags
Foo
    `

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	hasTag, err := libdocuments.HasTags(filePath, []string{"Foo", "Bar"})
	assert.NoError(t, err)
	assert.False(t, hasTag)
}

func TestHasTagsMultipleInInput(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Tags
Foo,Bar
    `

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	hasTag, err := libdocuments.HasTags(filePath, []string{"Foo", "Bar"})
	assert.NoError(t, err)
	assert.True(t, hasTag)
}
func TestHasTagsMultipleInInputReverseOrder(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Tags
Bar,Foo
    `

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	hasTag, err := libdocuments.HasTags(filePath, []string{"Foo", "Bar"})
	assert.NoError(t, err)
	assert.True(t, hasTag)
}

// ###########################
// # FindDocumentsWithTags() #
// ###########################
func TestFindDocumentsWithTagsNoInput(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	_, err = libdocuments.FindDocumentsWithTags(db, []string{})
	assert.Error(t, err)
}

func TestFindDocumentsWithTagsEmptyStringInput(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	_, err = libdocuments.FindDocumentsWithTags(db, []string{""})
	assert.Error(t, err)
}

func TestFindDocumentsWithTagsNoDocuments(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	_, err = libdocuments.FindDocumentsWithTags(db, []string{"Foo"})
	assert.Error(t, err)
}

func TestFindDocumentsWithTagsNoneFound(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	_, err = libdocuments.FindDocumentsWithTags(db, []string{"Bar"})
	assert.Error(t, err)
}

func TestFindDocumentsWithTags(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddTag(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	_, err = libdocuments.FindDocumentsWithTags(db, []string{"Bar"})
	assert.NoError(t, err)
}

// ####################
// # FindLinksLines() #
// ####################
func TestFindLinkLinesEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, _, _, err = libdocuments.FindLinksLines(filePath)
	assert.Error(t, err)
}

func TestFindLinkLinesHeaderOnly(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Links"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, _, _, err = libdocuments.FindLinksLines(filePath)
	assert.NoError(t, err)
}

func TestFindLinkLinesHeaderButNoLinks(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Links
Foo`

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, _, _, err = libdocuments.FindLinksLines(filePath)
	assert.Error(t, err)
}

func TestFindLinkLinesOneLink(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Links
- (Foo)[Bar]`

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	lineNrFirstLink, lineNrLastLink, links, err := libdocuments.FindLinksLines(filePath)
	assert.NoError(t, err)
	assert.Equal(t, 1, lineNrFirstLink)
	assert.Equal(t, 1, lineNrLastLink)
	assert.Len(t, links, 1)
	assert.Equal(t, "- (Foo)[Bar]", links[0])
}

func TestFindLinkLinesManyLinks(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Links
- (Foo)[Bar]
- (Foo)[Bar]
- (Foo)[Bar]`

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	lineNrFirstLink, lineNrLastLink, links, err := libdocuments.FindLinksLines(filePath)
	assert.NoError(t, err)
	assert.Equal(t, 1, lineNrFirstLink)
	assert.Equal(t, 3, lineNrLastLink)
	assert.Len(t, links, 3)
	assert.Equal(t, []string{"- (Foo)[Bar]", "- (Foo)[Bar]", "- (Foo)[Bar]"}, links)
}

// ########################
// # FindBacklinksLines() #
// ########################
func TestFindBacklinksLinesEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, _, _, err = libdocuments.FindBacklinksLines(filePath)
	assert.Error(t, err)
}

func TestFindBacklinksLinesHeaderOnly(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Backlinks"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, _, _, err = libdocuments.FindBacklinksLines(filePath)
	assert.NoError(t, err)
}

func TestFindBacklinksLinesHeaderButNoBacklinks(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Backlinks
Foo`

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	_, _, _, err = libdocuments.FindBacklinksLines(filePath)
	assert.Error(t, err)
}

func TestFindBacklinksLinesOneLink(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Backlinks
- (Foo)[Bar]`

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	lineNrFirstLink, lineNrLastLink, links, err := libdocuments.FindBacklinksLines(filePath)
	assert.NoError(t, err)
	assert.Equal(t, 1, lineNrFirstLink)
	assert.Equal(t, 1, lineNrLastLink)
	assert.Len(t, links, 1)
	assert.Equal(t, "- (Foo)[Bar]", links[0])
}

func TestFindBacklinksLinesManyBacklinks(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Backlinks
- (Foo)[Bar]
- (Foo)[Bar]
- (Foo)[Bar]`

	_, err = file.WriteString(document)
	assert.NoError(t, err)

	lineNrFirstLink, lineNrLastLink, links, err := libdocuments.FindBacklinksLines(filePath)
	assert.NoError(t, err)
	assert.Equal(t, 1, lineNrFirstLink)
	assert.Equal(t, 3, lineNrLastLink)
	assert.Len(t, links, 3)
	assert.Equal(t, []string{"- (Foo)[Bar]", "- (Foo)[Bar]", "- (Foo)[Bar]"}, links)
}

// #############
// # AddLink() #
// #############
func TestAddLinkToFileEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddLink(filePath, "Foo")
	assert.Error(t, err)
}

func TestAddLinkToFileHeaderOnly(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Links"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddLink(filePath, "Foo")
	assert.NoError(t, err)
}

func TestAddLinkToFileOneLink(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Links
- (Foo)[Bar]`
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddLink(filePath, "Foo")
	assert.NoError(t, err)
}

// #################
// # AddBackLink() #
// #################
func TestAddBacklinkToFileEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddBacklink(filePath, "Foo")
	assert.Error(t, err)
}

func TestAddBacklinkToFileHeaderOnly(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Backlinks"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddBacklink(filePath, "Foo")
	assert.NoError(t, err)
}

func TestAddBacklinkToFileOneBacklink(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Backlinks
- (Foo)[Bar]`
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.AddBacklink(filePath, "Bar")
	assert.NoError(t, err)
}

// ################
// # RemoveLink() #
// ################
func RemoveLinkFromFileEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveLink(filePath, "Foo")
	assert.Error(t, err)
}

func RemoveLinkFromFileHeaderOnly(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Links"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveLink(filePath, "Foo")
	assert.NoError(t, err)
}

func RemoveLinkFromFileNoMatch(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Backlinks
    - ()[Foo]`
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveLink(filePath, "Bar")
	assert.NoError(t, err)
}

func RemoveLinkFromFile(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Backlinks
    - ()[Foo]`
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveLink(filePath, "Foo")
	assert.NoError(t, err)
}

// ####################
// # RemoveBacklink() #
// ####################
func RemoveBacklinkFromFileEmpty(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := ""
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveBacklink(filePath, "Foo")
	assert.Error(t, err)
}

func RemoveBacklinkFromFileHeaderOnly(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := "# Backlinks"
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveBacklink(filePath, "Foo")
	assert.NoError(t, err)
}

func RemoveBacklinkFromFileNoMatch(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Backlinks
    - ()[Foo]`
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveBacklink(filePath, "Bar")
	assert.NoError(t, err)
}

func RemoveBacklinkFromFile(t *testing.T) {
	filePath := filepath.Join(test.TestDataTempDir, t.Name())

	file, err := os.Create(filePath)
	assert.NoError(t, err)

	document := `# Backlinks
    - ()[Foo]`
	_, err = file.WriteString(document)
	assert.NoError(t, err)

	err = libdocuments.RemoveBacklink(filePath, "Foo")
	assert.NoError(t, err)
}

// #################
// # AddDocument() #
// #################
func TestAddDocumentTypeDoesNotExist(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.Error(t, err)
}

func TestAddDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)
}

func TestAddDocumentDuplicate(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.Error(t, err)
}

func TestAddDocumentEmptyTag(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "")
	assert.Error(t, err)
}

// ####################
// # RemoveDocument() #
// ####################
func TestRemoveDocumentDocumentDoesNotExist(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.RemoveDocument(db, nil, "Foo")
	assert.Error(t, err)
}

func TestRemoveDocument(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.RemoveDocument(db, nil, "Foo")
	assert.NoError(t, err)
}

// ####################
// # RenameDocument() #
// ####################
func TestRenameDocumentDocumentDoesNotExist(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.RenameDocument(db, nil, "Foo", "Bar")
	assert.Error(t, err)
}

func TestRenameDoucment(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.RenameDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)
}

func TestRenameDoucmentNewNameExists(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Bar", "Bar")
	assert.NoError(t, err)

	err = libdocuments.RenameDocument(db, nil, "Foo", "Bar")
	assert.Error(t, err)
}

// ########################
// # ChangeDocumentType() #
// ########################
func TestChangeDocumentTypeDocumentDoesNotExist(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.ChangeDocumentType(db, nil, "Foo", "Bar")
	assert.Error(t, err)
}

func TestChangeDocumentTypeTypeDoesNotExist(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.ChangeDocumentType(db, nil, "Foo", "Baz")
	assert.Error(t, err)
}

func TestChangeDocumentType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libdocuments.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libdocuments.AddDocument(db, nil, "Foo", "Bar")
	assert.NoError(t, err)

	err = libdocuments.ChangeDocumentType(db, nil, "Foo", "Bar")
	assert.NoError(t, err)
}
