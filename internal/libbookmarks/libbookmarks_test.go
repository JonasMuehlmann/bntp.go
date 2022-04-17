package libbookmarks_test

import (
	"encoding/csv"
	"io"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libbookmarks"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/gocarina/gocsv"
	"github.com/stretchr/testify/assert"
)

// ##########################
// # DeserializeBookmarks() #
// ##########################.
func TestDeserializeBookmarksEmpty(t *testing.T) {
	csv := ""
	_, err := libbookmarks.DeserializeBookmarks(csv)
	assert.ErrorAs(t, err, &helpers.DeserializationError{})
}

func TestDeserializeBookmarksOnlyHeader(t *testing.T) {
	csv := "Url,Title"
	_, err := libbookmarks.DeserializeBookmarks(csv)
	assert.ErrorAs(t, err, &helpers.IneffectiveOperationError{})
}

func TestDeserializeBookmarksOneEntry(t *testing.T) {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ','

		return r
	})

	csv := `Url,Title
Foo,Bar`
	bookmarkList, err := libbookmarks.DeserializeBookmarks(csv)
	assert.NoError(t, err)
	assert.Len(t, bookmarkList, 1)
	assert.Equal(t, libbookmarks.Bookmark{Url: "Foo", Title: optional.Make("Bar")}, bookmarkList[0])
}

func TestDeserializeBookmarksManyEntries(t *testing.T) {
	csv := `Url,Title
Foo,Bar
Foo2,Bar2
`
	bookmarkList, err := libbookmarks.DeserializeBookmarks(csv)
	assert.NoError(t, err)
	assert.Len(t, bookmarkList, 2)
	assert.Equal(t, libbookmarks.Bookmark{Url: "Foo", Title: optional.Make("Bar")}, bookmarkList[0])
	assert.Equal(t, libbookmarks.Bookmark{Url: "Foo2", Title: optional.Make("Bar2")}, bookmarkList[1])
}

// ########################
// # SerializeBookmarks() #
// ########################.
func TestSerializeBookmarksEmpty(t *testing.T) {
	_, err := libbookmarks.SerializeBookmarks([]libbookmarks.Bookmark{})
	assert.ErrorAs(t, err, &helpers.IneffectiveOperationError{})
}

func TestSerializeBookmarks(t *testing.T) {
	bookmarks := []libbookmarks.Bookmark{
		{Title: optional.Make("Foo"), Url: "Bar"},
		{Title: optional.Make("Foo1"), Url: "Bar1"},
	}

	serializedBookmarks, err := libbookmarks.SerializeBookmarks(bookmarks)
	assert.Contains(t, serializedBookmarks, "Title")
	assert.Contains(t, serializedBookmarks, "Url")

	assert.Contains(t, serializedBookmarks, "Foo")
	assert.Contains(t, serializedBookmarks, "Foo1")

	assert.Contains(t, serializedBookmarks, "Bar")
	assert.Contains(t, serializedBookmarks, "Bar1")
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

	_, err = helpers.GetIdFromBookmarkType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Make("Bar")})
	assert.NoError(t, err)
}

func TestAddBookmarkTansaction(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	_, err = helpers.GetIdFromBookmarkType(db, nil, "Bar")
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(nil, transaction, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Make("Bar")})
	assert.NoError(t, err)

	err = transaction.Commit()
	assert.NoError(t, err)
}

func TestAddBookmarkNoUrl(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "", Type: optional.Optional[string]{}})
	assert.Error(t, err)
}

func TestAddBookmarkNoType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
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

	_, err = helpers.GetIdFromBookmarkType(db, nil, "Bar")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Make("Bar")})
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

	_, err = helpers.GetIdFromBookmarkType(db, nil, "Bar")
	assert.NoError(t, err)

	transaction, err := db.Beginx()
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(nil, transaction, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Make("Bar")})
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

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddTag(db, nil, bookmarkId, "Foo")
	assert.NoError(t, err)
}

func TestAddTagToBookmarkTransaction(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Foo")
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

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
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

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Foo")
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

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
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

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.EditIsRead(db, nil, bookmarkId, true)
	assert.NoError(t, err)
}

func TestEditIsCollection(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.EditIsCollection(db, nil, bookmarkId, true)
	assert.NoError(t, err)
}

func TestEditTitle(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.EditTitle(db, nil, bookmarkId, "Bar")
	assert.NoError(t, err)
}

func TestEditUrl(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Foo")
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

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "Bar", Type: optional.Optional[string]{}})
	assert.NoError(t, err)

	bookmarkId, err := helpers.GetIdFromBookmark(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.EditType(db, nil, bookmarkId, "Baz")
	assert.NoError(t, err)
}
