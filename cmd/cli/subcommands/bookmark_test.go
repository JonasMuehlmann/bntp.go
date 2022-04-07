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
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/cmd/cli/subcommands"
	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libbookmarks"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
)

// ******************************************************************//
//                             --import                             //
// ******************************************************************//.
func TestImportBookmarks(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	file, err := test.CreateTestTempFile(t.Name())

	file.WriteString("Title;Url\nfoo;bar")

	os.Args = []string{"", "bookmark", "--import", path.Join(test.TestDataTempDir, t.Name())}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                             --export                             //
// ******************************************************************//.
func TestExportBookmarskUnfiltered(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "bar", sql.NullInt32{})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--export", path.Join(test.TestDataTempDir, t.Name())}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

func TestExportBookmarksFiltered(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	url := "bar"
	isRead := false
	filter := libbookmarks.BookmarkFilter{Url: &url, IsRead: &isRead}

	filterSerialized, err := json.Marshal(filter)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "bar", sql.NullInt32{})
	os.Args = []string{"", "bookmark", "--export", path.Join(test.TestDataTempDir, t.Name()), "--filter", string(filterSerialized)}
	assert.NoError(t, err)

	subcommands.BookmarkMain(db, helpers.NOPExiter)
	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                              --list                              //
// ******************************************************************//.
func TestListBookmarksUnfiltered(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "bar", sql.NullInt32{})
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Bar", "abc", sql.NullInt32{})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--list"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	stdOutInterceptBuffer.Scan()
	stdOutInterceptBuffer.Scan()
	assert.Contains(t, stdOutInterceptBuffer.Text(), "Foo")
	assert.Contains(t, stdOutInterceptBuffer.Text(), "bar")

	assert.Empty(t, logInterceptBuffer.String())
}

func TestListBookmarksFiltered(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	url := "bar"
	isRead := false
	filter := libbookmarks.BookmarkFilter{Url: &url, IsRead: &isRead}

	filterSerialized, err := json.Marshal(filter)
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "Foo", "bar", sql.NullInt32{})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--list", "--filter", string(filterSerialized)}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	stdOutInterceptBuffer.Scan()
	stdOutInterceptBuffer.Scan()
	assert.Contains(t, stdOutInterceptBuffer.Text(), "Foo")
	assert.Contains(t, stdOutInterceptBuffer.Text(), "bar")

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                            --add-type                            //
// ******************************************************************//.
func TestAddType(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--add-type", "foo"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                           --remove-type                          //
// ******************************************************************//.
func TestRemoveType(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--remove-type", "foo"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                           --list-types                           //
// ******************************************************************//.
func TestListTypes(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	stdOutInterceptBuffer, reader, writer := test.InterceptStdout(t)
	defer test.ResetStdout(t, reader, writer)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Bar")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--list-types"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, "Foo", stdOutInterceptBuffer.Text())

	stdOutInterceptBuffer.Scan()
	assert.Equal(t, "Bar", stdOutInterceptBuffer.Text())

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                               --add                              //
// ******************************************************************//.
func TestAddBookmark(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--add", `{"title": "foo", "url": "bar", "type": "1"}`}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                              --edit                              //
// ******************************************************************//.
func TestEditBookmark(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "foo", "bar", sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit", `{"title": "foo", "url": "bar", "type": "1", "Id": 1}`}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                          --edit-is-read                          //
// ******************************************************************//.
func TestEditIsRead(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "foo", "bar", sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-is-read", "1", "true"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                          --edit-title                            //
// ******************************************************************//.
func TestEditTitle(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "foo", "bar", sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-title", "1", "foo"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                            --edit-url                            //
// ******************************************************************//.
func TestEditUrl(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "foo", "bar", sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-title", "1", "foo"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                            --edit-type                           //
// ******************************************************************//.
func TestEditType(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "foo", "bar", sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-type", "1", "Foo"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                       --edit-is-collection                       //
// ******************************************************************//.
func TestEditIsCollection(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "foo", "bar", sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--edit-is-collection", "1", "true"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                             --add-tag                            //
// ******************************************************************//.
func TestAddTagToBookmark(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "foo", "bar", sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--add-tag", "1", "Foo"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}

// ******************************************************************//
//                           --remove-tag                           //
// ******************************************************************//.
func TestRemoveTagFromBookmark(t *testing.T) {
	logInterceptBuffer := strings.Builder{}
	log.SetOutput(&logInterceptBuffer)

	defer log.SetOutput(os.Stderr)

	db, err := test.GetDB(t)
	assert.NoError(t, err)

	err = libbookmarks.AddType(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddBookmark(db, nil, "foo", "bar", sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	assert.NoError(t, err)

	err = libtags.AddTag(db, nil, "Foo")
	assert.NoError(t, err)

	err = libbookmarks.AddTag(db, nil, 1, "Foo")
	assert.NoError(t, err)

	os.Args = []string{"", "bookmark", "--remove-tag", "1", "Foo"}
	subcommands.BookmarkMain(db, helpers.NOPExiter)

	assert.Empty(t, logInterceptBuffer.String())
}
