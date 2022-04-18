package libbookmarks_test

import (
	"path/filepath"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/pkg/helpers"
	"github.com/JonasMuehlmann/bntp.go/pkg/libbookmarks"
	"github.com/JonasMuehlmann/bntp.go/pkg/libtags"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/stretchr/testify/assert"
)

// ######################
// # ImportMinimalCSV() #
// ######################.
func TestImportMinimalCSVEmpty(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	csv := ""
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVNoHeaderButData(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	csv := "Foo;Bar"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVHeaderNoTitle(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	csv := "dss;Title"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVHeaderNoUrl(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	csv := "dss;Url"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVOnlyHeader(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	csv := "Url;Title"
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVOneEntry(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.NoError(t, err)
}

func TestImportMinimalCSVManyEntries(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar
Foo2;Bar2
Foo3;Bar3`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.NoError(t, err)
}

func TestImportMinimalCSVEntryWithIncompleteUrl(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar
Foo2;Bar2
;Bar3`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestImportMinimalCSVEntryWithIncompleteTitle(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())
	assert.NoError(t, err)

	csv := `Url;Title
Foo;Bar
Foo2;Bar2
Foo3;`
	_, err = file.WriteString(csv)
	assert.NoError(t, err)

	err = libbookmarks.ImportMinimalCSV(db, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.Error(t, err)
}

// ###############
// # ExportCSV() #
// ###############.
func TestExportCSVEmpty(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	bookmarks, err := libbookmarks.GetBookmarks(db, libbookmarks.BookmarkFilter{})
	assert.NoError(t, err)

	err = libbookmarks.ExportCSV(bookmarks, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.Error(t, err)
}

func TestExportCSV(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{HasValue: false})
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo2", "Bar2", optional.Optional[int]{HasValue: false})
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo3", "Bar3", optional.Optional[int]{HasValue: false})
	assert.NoError(t, err)

	bookmarks, err := libbookmarks.GetBookmarks(db, libbookmarks.BookmarkFilter{})
	assert.NoError(t, err)

	err = libbookmarks.ExportCSV(bookmarks, filepath.Join(test.TestDataTempDir, t.Name()))
	assert.NoError(t, err)
}

// #############
// # AddType() #
// #############.
func TestAddTypeEmpty(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "")
	assert.Error(t, err)
}

func TestAddType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)
}

func TestAddTypeTransaction(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libbookmarks.AddType(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

// ################
// # RemoveType() #
// ################.
func TestRemoveTypeEmpty(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.RemoveType(db, nil, "")
	assert.Error(t, err)
}

func TestRemoveTypeNonExistent(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.RemoveType(db, nil, "Foo")
	assert.Error(t, err)
}

func TestRemoveType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.RemoveType(db, nil, "Foo")
	assert.NoError(t, err)
}

func TestRemoveTypeTransaction(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libbookmarks.AddType(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.RemoveType(nil, transaction, "Foo")
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

// ###############
// # ListTypes() #
// ###############.
func TestListTypes(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	types, err := libbookmarks.ListTypes(db)
	assert.NoError(t, err)
	assert.Len(t, types, 1)
	assert.Equal(t, "Foo", types[0])
}

func TestListTypesMany(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo2")
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo3")
	assert.NoError(t, err)

	types, err := libbookmarks.ListTypes(db)
	assert.NoError(t, err)
	assert.Len(t, types, 3)
	assert.Equal(t, []string{"Foo", "Foo2", "Foo3"}, types)
}

func TestListTypesEmpty(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	types, err := libbookmarks.ListTypes(db)
	assert.NoError(t, err)
	assert.Len(t, types, 0)
}

// #################
// # AddBookmark() #
// #################.
func TestAddBookmark(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	typeId, err := helpers.GetIdFromBookmarkType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: int(typeId), HasValue: true})
	assert.NoError(t, err)
}

func TestAddBookmarkTansaction(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	typeId, err := helpers.GetIdFromBookmarkType(db, nil, "Bar")
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(nil, transaction, "Foo", "Bar", optional.Optional[int]{Wrappee: int(typeId), HasValue: true})
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

func TestAddBookmarkNoTitle(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: true})
	assert.Error(t, err)
}

func TestAddBookmarkNoUrl(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "", optional.Optional[int]{Wrappee: 0, HasValue: true})
	assert.Error(t, err)
}

func TestAddBookmarkNoType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)
}

// ####################
// # RemoveBookmark() #
// ####################.
func TestRemoveBookmark(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	typeId, err := helpers.GetIdFromBookmarkType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: int(typeId), HasValue: true})
	assert.NoError(t, err)

	bookmarks, err := libbookmarks.GetBookmarks(db, libbookmarks.BookmarkFilter{})
	assert.NoError(t, err)

	err = libbookmarks.RemoveBookmark(db, nil, bookmarks[0].Id)
	assert.NoError(t, err)
}

func TestRemoveBookmarkTansaction(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)
	err = libbookmarks.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	typeId, err := helpers.GetIdFromBookmarkType(db, nil, "Bar")
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(nil, transaction, "Foo", "Bar", optional.Optional[int]{Wrappee: int(typeId), HasValue: true})
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)

	bookmarks, err := libbookmarks.GetBookmarks(db, libbookmarks.BookmarkFilter{})
	assert.NoError(t, err)

	err = libbookmarks.RemoveBookmark(db, nil, bookmarks[0].Id)
	assert.NoError(t, err)
}

func TestRemoveBookmarkNonExistent(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.RemoveBookmark(db, nil, 0)
	assert.Error(t, err)
}

// ############
// # AddTag() #
// ############.
func TestAddTagToBookmark(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.AddTag(db, nil, bookmarkId, "Foo")
	assert.NoError(t, err)
}

func TestAddTagToBookmarkTransaction(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libbookmarks.AddTag(nil, transaction, bookmarkId, "Foo")
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

func TestAddTagToBookmarkNoBookmark(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddTag(db, nil, 0, "Foo")
	assert.Error(t, err)
}

func TestAddTagToBookmarkNoTag(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.AddTag(db, nil, bookmarkId, "Foo")
	assert.Error(t, err)
}

// ###äää#########
// # RemoveTag() #
// ###äää#########.
func TestRemoveTagFromBookmark(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.AddTag(db, nil, bookmarkId, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.RemoveTag(db, nil, bookmarkId, "Foo")
	assert.NoError(t, err)
}

func TestRemoveTagFromBookmarkTagDoesNotExist(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.RemoveTag(db, nil, bookmarkId, "Foo")
	assert.NoError(t, err)
}

func TestRemoveTagFromBookmarkBookmarkDoesNotExist(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.RemoveTag(db, nil, 0, "Foo")
	assert.NoError(t, err)
}

func TestEditIsRead(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.EditIsRead(db, nil, bookmarkId, true)
	assert.NoError(t, err)
}

func TestEditIsCollection(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.EditIsCollection(db, nil, bookmarkId, true)
	assert.NoError(t, err)
}

func TestEditTitle(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.EditTitle(db, nil, bookmarkId, "Bar")
	assert.NoError(t, err)
}

func TestEditUrl(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.EditUrl(db, nil, bookmarkId, "Bar")
	assert.NoError(t, err)
}

func TestEditType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Baz")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "Bar", optional.Optional[int]{Wrappee: 0, HasValue: false})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.EditType(db, nil, bookmarkId, "Baz")
	assert.NoError(t, err)
}
