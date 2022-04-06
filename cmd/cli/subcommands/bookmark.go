// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package subcommands

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libbookmarks"
	"github.com/docopt/docopt-go"
	"github.com/jmoiron/sqlx"
)

var usageBookmark string = `bntp bookmark - Interact with bookmarks.

Usage:
    bntp bookmark -h | --help
    bntp bookmark -l [--filter FILTER]
    bntp bookmark --import FILE
    bntp bookmark --export FILE [--filter FILTER]
    bntp bookmark (--add-type | --remove-type) TYPE
    bntp bookmark --list-types
    bntp bookmark -a DATA
    bntp bookmark -r BOOKMARK_ID
    bntp bookmark -E NEW_DATA
    bntp bookmark --edit-is-read BOOKMARK_ID IS_READ
    bntp bookmark --edit-title BOOKMARK_ID TITLE
    bntp bookmark --edit-url BOOKMARK_ID TITLE
    bntp bookmark --edit-type BOOKMARK_ID TITLE
    bntp bookmark --edit-is-collection BOOKMARK_ID IS_COLLECTION
    bntp bookmark (--add-tag | --remove-tag) BOOKMARK_ID TAG

Options:
    -h --help               Show this screen.
    -i --import             Import bookmarks.
    -e --export             Export bookmarks.
    -l --list               List bookmarks.
    --filter                A filter to apply when searching for bookmarks.
    --add-type              Add a bookmark type.
    --remove-type           Remove a bookmark type.
    --list-types            List the bookmark types.
    -a --add                Add a bookmark.
    -r --remove             Remove a bookmark.
    -E --edit               Edit a bookmark.
    --edit-is-read          Change the is read status of a bookmark.
    --edit-title            Change the title of a bookmark.
    --edit-url              Change the url of a bookmark.
    --edit-type             Change the type of a bookmark.
    --edit-is-collection    Change the is collection status of a bookmark.
    --add-tag               Add a tag to a bookmark.
    --remove-tag            Remove a tag from a bookmark.
`

func BookmarkMain(db *sqlx.DB, exiter func(int)) {
	arguments, err := docopt.ParseDoc(usageBookmark)
	helpers.OnError(err, helpers.MakeFatalLogger(exiter))

	// ******************************************************************//
	if isSet, ok := arguments["--import"]; ok && isSet.(bool) {
		source, err := arguments.String("FILE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.ImportMinimalCSV(db, source)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--export"]; ok && isSet.(bool) {
		source, err := arguments.String("FILE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		filter := libbookmarks.BookmarkFilter{}
		if isSet, ok := arguments["--filter"]; ok && isSet.(bool) {
			filterRaw, err := arguments.String("FILTER")
			helpers.OnError(err, helpers.MakeFatalLogger(exiter))

			err = json.Unmarshal([]byte(filterRaw), &filter)
			helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		}

		bookmarks, err := libbookmarks.GetBookmarks(db, filter)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.ExportCSV(bookmarks, source)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--list"]; ok && isSet.(bool) {
		filter := libbookmarks.BookmarkFilter{}
		if isSet, ok := arguments["--filter"]; ok && isSet.(bool) {
			filterRaw, err := arguments.String("FILTER")
			helpers.OnError(err, helpers.MakeFatalLogger(exiter))

			err = json.Unmarshal([]byte(filterRaw), &filter)
			helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		}

		bookmarks, err := libbookmarks.GetBookmarks(db, filter)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		for _, bookmark := range bookmarks {
			fmt.Println(bookmark)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-type"]; ok && isSet.(bool) {
		type_, err := arguments.String("TYPE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.AddType(db, nil, type_)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-type"]; ok && isSet.(bool) {
		type_, err := arguments.String("TYPE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.RemoveType(db, nil, type_)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--list-types"]; ok && isSet.(bool) {
		types, err := libbookmarks.ListTypes(db)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		for _, type_ := range types {
			fmt.Println(type_)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add"]; ok && isSet.(bool) {
		var data map[string]string
		dataRaw, err := arguments.String("DATA")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = json.Unmarshal([]byte(dataRaw), &data)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		title, ok := data["title"]
		if !ok {
			log.Panic("Missing parameter title in DATA")
		}

		url, ok := data["url"]
		if !ok {
			log.Panic("Missing parameter url in DATA")
		}

		var type_ sql.NullInt32
		typeRaw, ok := data["type"]
		if !ok {
			type_.Valid = false
		}

		typeInt, err := strconv.ParseInt(typeRaw, 10, 32)

		type_.Int32 = int32(typeInt)

		err = libbookmarks.AddBookmark(db, nil, title, url, type_)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit"]; ok && isSet.(bool) {
		var data libbookmarks.Bookmark
		dataRaw, err := arguments.String("NEW_DATA")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = json.Unmarshal([]byte(dataRaw), &data)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.EditBookmark(db, nil, data)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-is-read"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		isReadRaw, err := arguments.String("IS_READ")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		isRead, err := strconv.ParseBool(isReadRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.EditIsRead(db, nil, ID, isRead)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-title"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		title, err := arguments.String("TITLE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.EditTitle(db, nil, ID, title)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-url"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		url, err := arguments.String("URL")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.EditUrl(db, nil, ID, url)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-type"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		type_, err := arguments.String("TYPE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.EditType(db, nil, ID, type_)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-is-collection"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		isCollectionRaw, err := arguments.String("IS_COLLECTION")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		isCollection, err := strconv.ParseBool(isCollectionRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.EditIsCollection(db, nil, ID, isCollection)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-tag"]; ok && isSet.(bool) {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.AddType(db, nil, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.AddTag(db, nil, ID, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-tag"]; ok && isSet.(bool) {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		IDRaw, err := arguments.String("BOOKMARK_ID")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.RemoveType(db, nil, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		ID, err := strconv.Atoi(IDRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libbookmarks.RemoveTag(db, nil, ID, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
	}
}
