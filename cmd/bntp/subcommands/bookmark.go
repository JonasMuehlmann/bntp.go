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
	"log"
	"strconv"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libbookmarks"
	"github.com/docopt/docopt-go"
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
    bntp bookmark -e NEW_DATA
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
    -e --edit               Edit a bookmark.
    --edit-is-read          Change the is read status of a bookmark.
    --edit-title            Change the title of a bookmark.
    --edit-url              Change the url of a bookmark.
    --edit-type             Change the type of a bookmark.
    --edit-is-collection    Change the is collection status of a bookmark.
    --add-tag               Add a tag to a bookmark.
    --remove-tag            Remove a tag from a bookmark.
`

func BookmarkMain() {
	arguments, err := docopt.ParseDoc(usageBookmark)
	OnError(err, log.Fatal)

	db, err := helpers.GetDefaultDB()
	OnError(err, log.Fatal)

	//******************************************************************//
	if _, ok := arguments["--import"]; ok {
		source, err := arguments.String("FILE")
		OnError(err, log.Fatal)

		err = libbookmarks.ImportMinimalCSV(db, source)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--export"]; ok {
		source, err := arguments.String("FILE")
		OnError(err, log.Fatal)

		filter := libbookmarks.BookmarkFilter{}
		if _, ok := arguments["--filter"]; ok {
			filterRaw, err := arguments.String("FILTER")
			OnError(err, log.Fatal)

			err = json.Unmarshal([]byte(filterRaw), &filter)
			OnError(err, log.Fatal)
		}

		bookmarks, err := libbookmarks.GetBookmarks(db, filter)
		OnError(err, log.Fatal)

		err = libbookmarks.ExportCSV(bookmarks, source)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--list"]; ok {
		filter := libbookmarks.BookmarkFilter{}
		if _, ok := arguments["--filter"]; ok {
			filterRaw, err := arguments.String("FILTER")
			OnError(err, log.Fatal)

			err = json.Unmarshal([]byte(filterRaw), &filter)
			OnError(err, log.Fatal)
		}

		bookmarks, err := libbookmarks.GetBookmarks(db, filter)
		OnError(err, log.Fatal)

		for bookmark := range bookmarks {
			println(bookmark)
		}
		//******************************************************************//
	} else if _, ok := arguments["--add-type"]; ok {
		type_, err := arguments.String("TYPE")
		OnError(err, log.Fatal)

		err = libbookmarks.AddType(db, nil, type_)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--remove-type"]; ok {
		type_, err := arguments.String("TYPE")
		OnError(err, log.Fatal)

		err = libbookmarks.RemoveType(db, nil, type_)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--list-types"]; ok {
		types, err := libbookmarks.ListTypes(db)
		OnError(err, log.Fatal)

		for type_ := range types {
			println(type_)
		}
		//******************************************************************//
	} else if _, ok := arguments["--add"]; ok {
		var data map[string]string
		dataRaw, err := arguments.String("DATA")
		OnError(err, log.Fatal)

		err = json.Unmarshal([]byte(dataRaw), &data)
		OnError(err, log.Fatal)

		title, ok := data["title"]
		if !ok {
			log.Fatal("Missing parameter title in DATA")
		}

		url, ok := data["url"]
		if !ok {
			log.Fatal("Missing parameter url in DATA")
		}

		var type_ sql.NullInt32
		typeRaw, ok := data["type"]
		if !ok {
			type_.Valid = false
		}

		typeInt, err := strconv.ParseInt(typeRaw, 10, 32)

		type_.Int32 = int32(typeInt)

		err = libbookmarks.AddBookmark(db, nil, title, url, type_)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--edit"]; ok {
		var data libbookmarks.Bookmark
		dataRaw, err := arguments.String("NEW_DATA")
		OnError(err, log.Fatal)

		err = json.Unmarshal([]byte(dataRaw), &data)
		OnError(err, log.Fatal)

		err = libbookmarks.EditBookmark(db, nil, data)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--edit-is-read"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		OnError(err, log.Fatal)

		isReadRaw, err := arguments.String("IS_READ")
		OnError(err, log.Fatal)

		ID, err := strconv.Atoi(IDRaw)
		OnError(err, log.Fatal)

		isRead, err := strconv.ParseBool(isReadRaw)
		OnError(err, log.Fatal)

		err = libbookmarks.EditIsRead(db, nil, ID, isRead)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--edit-title"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		OnError(err, log.Fatal)

		title, err := arguments.String("TITLE")
		OnError(err, log.Fatal)

		ID, err := strconv.Atoi(IDRaw)
		OnError(err, log.Fatal)

		err = libbookmarks.EditTitle(db, nil, ID, title)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--edit-url"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		OnError(err, log.Fatal)

		url, err := arguments.String("URL")
		OnError(err, log.Fatal)

		ID, err := strconv.Atoi(IDRaw)
		OnError(err, log.Fatal)

		err = libbookmarks.EditUrl(db, nil, ID, url)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--edit-type"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		OnError(err, log.Fatal)

		type_, err := arguments.String("TYPE")
		OnError(err, log.Fatal)

		ID, err := strconv.Atoi(IDRaw)
		OnError(err, log.Fatal)

		err = libbookmarks.EditType(db, nil, ID, type_)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--edit-is-collection"]; ok {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		OnError(err, log.Fatal)

		isCollectionRaw, err := arguments.String("IS_COLLECTION")
		OnError(err, log.Fatal)

		ID, err := strconv.Atoi(IDRaw)
		OnError(err, log.Fatal)

		isCollection, err := strconv.ParseBool(isCollectionRaw)
		OnError(err, log.Fatal)

		err = libbookmarks.EditIsCollection(db, nil, ID, isCollection)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--add-tag"]; ok {
		tag, err := arguments.String("TAG")
		OnError(err, log.Fatal)

		IDRaw, err := arguments.String("BOOKMARK_ID")
		OnError(err, log.Fatal)

		err = libbookmarks.AddType(db, nil, tag)
		OnError(err, log.Fatal)

		ID, err := strconv.Atoi(IDRaw)
		OnError(err, log.Fatal)

		err = libbookmarks.AddTag(db, nil, ID, tag)
		OnError(err, log.Fatal)
		//******************************************************************//
	} else if _, ok := arguments["--remove-tag"]; ok {
		tag, err := arguments.String("TAG")
		OnError(err, log.Fatal)

		IDRaw, err := arguments.String("BOOKMARK_ID")
		OnError(err, log.Fatal)

		err = libbookmarks.RemoveType(db, nil, tag)
		OnError(err, log.Fatal)

		ID, err := strconv.Atoi(IDRaw)
		OnError(err, log.Fatal)

		err = libbookmarks.RemoveTag(db, nil, ID, tag)
		OnError(err, log.Fatal)
	}
}
