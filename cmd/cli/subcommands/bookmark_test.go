// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the"Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED"AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package subcommands_test

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/cmd/cli/subcommands"
	"github.com/JonasMuehlmann/bntp.go/internal/libbookmarks"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/stretchr/testify/assert"
)

// ******************************************************************//
//                             --import                             //
// ******************************************************************//.
func TestImportBookmarks(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())

	file.WriteString("Title;Url\nfoo;bar")

	os.Args = []string{"", "bookmark", "--import", path.Join(test.TestDataTempDir, t.Name())}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                             --export                             //
// ******************************************************************//.
func TestExportBookmarskUnfiltered(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "bar"})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--export", path.Join(test.TestDataTempDir, t.Name())}
	err = subcommands.BookmarkMain(db)
	assert.NoError(t, err)
}

func TestExportBookmarksFiltered(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	url := optional.Optional[string]{Wrappee: "bar", HasValue: true}
	isRead := optional.Optional[bool]{Wrappee: false, HasValue: true}
	filter := libbookmarks.BookmarkFilter{Url: url, IsRead: isRead}

	filterSerialized, err := json.Marshal(filter)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "bar"})
	os.Args = []string{"", "bookmark", "--export", path.Join(test.TestDataTempDir, t.Name()), "--filter", string(filterSerialized)}
	assert.NoError(t, err)

	err = subcommands.BookmarkMain(db)
	assert.NoError(t, err)
}

// ******************************************************************//
//                              --list                              //
// ******************************************************************//.
func TestListBookmarksUnfiltered(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "bar"})
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "abc"})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--list"}
	err = subcommands.BookmarkMain(db)

	stdOutInterceptBuffer.Scan()
	stdOutInterceptBuffer.Scan()
	assert.Contains(t, stdOutInterceptBuffer.Text(), "Foo")
	assert.Contains(t, stdOutInterceptBuffer.Text(), "bar")

	assert.NoError(t, err)
}

func TestListBookmarksFiltered(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	url := optional.Optional[string]{Wrappee: "bar", HasValue: true}
	isRead := optional.Optional[bool]{Wrappee: false, HasValue: true}
	filter := libbookmarks.BookmarkFilter{Url: url, IsRead: isRead}

	filterSerialized, err := json.Marshal(filter)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("Foo"), Url: "bar"})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--list", "--filter", string(filterSerialized)}
	err = subcommands.BookmarkMain(db)

	stdOutInterceptBuffer.Scan()
	stdOutInterceptBuffer.Scan()
	assert.Contains(t, stdOutInterceptBuffer.Text(), "Foo")
	assert.Contains(t, stdOutInterceptBuffer.Text(), "bar")

	assert.NoError(t, err)
}

// ******************************************************************//
//                            --add-type                            //
// ******************************************************************//.
func TestAddType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--add-type", "foo"}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                           --remove-type                          //
// ******************************************************************//.
func TestRemoveType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--remove-type", "foo"}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                           --list-types                           //
// ******************************************************************//.
func TestListTypes(t *testing.T) {
	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--list-types"}
	err = subcommands.BookmarkMain(db)

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, "Foo", stdOutInterceptBuffer.Text())

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, "Bar", stdOutInterceptBuffer.Text())

	assert.NoError(t, err)
}

// ******************************************************************//
//                               --add                              //
// ******************************************************************//.
func TestAddBookmark(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--add", `{"title": "foo", "url": "bar", "type": "Foo"}`}
	err = subcommands.BookmarkMain(db)
	assert.NoError(t, err)
}

func TestAddBookmarkNoUrl(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--add", `{"title": "bar", "type": "Foo"}`}
	err = subcommands.BookmarkMain(db)

	assert.ErrorAs(t, err, &subcommands.IncompleteCompoundParameterError{})
}

func TestAddBookmarkNoType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--add", `{"title": "foo", "url": "bar"}`}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                              --edit                              //
// ******************************************************************//.
func TestEditBookmark(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("foo"), Url: "bar"})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit", `{"title": "foo", "url": "bar", "type": "1", "Id": 1}`}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                          --edit-is-read                          //
// ******************************************************************//.
func TestEditIsRead(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("foo"), Url: "bar"})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-is-read", "1", "true"}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                          --edit-title                            //
// ******************************************************************//.
func TestEditTitle(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("foo"), Url: "bar"})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-title", "1", "foo"}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                            --edit-url                            //
// ******************************************************************//.
func TestEditUrl(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("foo"), Url: "bar"})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-url", "1", "foo"}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                            --edit-type                           //
// ******************************************************************//.
func TestEditType(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("foo"), Url: "bar"})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-type", "1", "Foo"}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                       --edit-is-collection                       //
// ******************************************************************//.
func TestEditIsCollection(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("foo"), Url: "bar"})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-is-collection", "1", "true"}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                             --add-tag                            //
// ******************************************************************//.
func TestAddTagToBookmark(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("foo"), Url: "bar"})
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--add-tag", "1", "Foo"}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}

// ******************************************************************//
//                           --remove-tag                           //
// ******************************************************************//.
func TestRemoveTagFromBookmark(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, libbookmarks.Bookmark{Title: optional.Make("foo"), Url: "bar"})
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddTag(db, nil, 1, "Foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--remove-tag", "1", "Foo"}
	err = subcommands.BookmarkMain(db)

	assert.NoError(t, err)
}
