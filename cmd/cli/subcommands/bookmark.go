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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/libbookmarks"
	"github.com/docopt/docopt-go"
	"github.com/jmoiron/sqlx"
)

var usageBookmark string = `bntp bookmark - Interact with bookmarks.

Usage:
    bntp bookmark -h | --help
    bntp bookmark -l [--filter FILTER] [--format FORMAT]
    bntp bookmark --import FILE
    bntp bookmark --export FILE [--filter FILTER]
    bntp bookmark (--add-type | --remove-type) TYPE
    bntp bookmark --list-types
    bntp bookmark -a DATA
    bntp bookmark -r BOOKMARK_ID
    bntp bookmark -E NEW_DATA
    bntp bookmark --edit-is-read BOOKMARK_ID IS_READ
    bntp bookmark --edit-title BOOKMARK_ID TITLE
    bntp bookmark --edit-url BOOKMARK_ID URL
    bntp bookmark --edit-type BOOKMARK_ID TYPE
    bntp bookmark --edit-is-collection BOOKMARK_ID IS_COLLECTION
    bntp bookmark (--add-tag | --remove-tag) BOOKMARK_ID TAG

Options:
    -h --help               Show this screen.
    -i --import             Import bookmarks.
    -e --export             Export bookmarks.
    -l --list               List bookmarks.
    --filter=FILTER         A filter to apply when searching for bookmarks.
    --format=FORLAT         A format to return data in [default: csv]
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

var AllowedFormatValues = []any{"csv"}

func BookmarkMain(db *sqlx.DB) error {
	arguments, err := docopt.ParseDoc(usageBookmark)
	if err != nil {
		return err
	}

	// ******************************************************************//
	if isSet, ok := arguments["--import"]; ok && isSet.(bool) {
		source, err := arguments.String("FILE")
		if err != nil {
			return ParameterConversionError{"FILE", arguments["FILE"], "string"}
		}

		serializedBookmarks, err := os.ReadFile(source)

		deserializedBookmarks, err := libbookmarks.DeserializeBookmarks(string(serializedBookmarks))
		if err != nil {
			return err
		}

		err = libbookmarks.ImportBookmarks(db, deserializedBookmarks)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--export"]; ok && isSet.(bool) {
		destination, err := arguments.String("FILE")
		if err != nil {
			return ParameterConversionError{"FILE", arguments["FILE"], "string"}
		}

		filter := libbookmarks.BookmarkFilter{}
		if value, ok := arguments["--filter"]; ok && value != nil {
			filterRaw, err := arguments.String("--filter")
			if err != nil {
				return ParameterConversionError{"--filter", arguments["--filter"], "string"}
			}

			err = json.Unmarshal([]byte(filterRaw), &filter)
			if err != nil {
				return err
			}
		}

		bookmarks, err := libbookmarks.ExportBookmarks(db)
		if err != nil {
			return err
		}

		serializedBookmarks, err := libbookmarks.SerializeBookmarks(bookmarks)
		if err != nil {
			return err
		}

		os.WriteFile(destination, []byte(serializedBookmarks), 0644)

		// ******************************************************************//
	} else if isSet, ok := arguments["--list"]; ok && isSet.(bool) {
		filter := libbookmarks.BookmarkFilter{}
		if value, ok := arguments["--filter"]; ok && value != nil {
			filterRaw, err := arguments.String("--filter")
			if err != nil {
				return ParameterConversionError{"--filter", arguments["--filter"], "string"}
			}

			err = json.Unmarshal([]byte(filterRaw), &filter)
			if err != nil {
				return err
			}
		}

		bookmarks, err := libbookmarks.GetBookmarks(db, filter)
		if err != nil {
			return err
		}

		var exportFormat string
		if _, ok := arguments["--format"]; ok {
			exportFormat, err = arguments.String("--format")
			if err != nil {
				return ParameterConversionError{"--format", arguments["--format"], "string"}
			}
		}

		switch strings.ToLower(exportFormat) {
		case "csv":
			serializedBookmarks, err := libbookmarks.SerializeBookmarks(bookmarks)
			if err != nil {
				return err
			}

			fmt.Println(serializedBookmarks)
		default:
			return InvalidParameterValueError{arguments["FORMAT"], arguments["FORMAT"], AllowedFormatValues}
		}

		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-type"]; ok && isSet.(bool) {
		type_, err := arguments.String("TYPE")
		if err != nil {
			return ParameterConversionError{"TYPE", arguments["TYPE"], "string"}
		}

		err = libbookmarks.AddType(db, nil, type_)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-type"]; ok && isSet.(bool) {
		type_, err := arguments.String("TYPE")
		if err != nil {
			return ParameterConversionError{"TYPE", arguments["TYPE"], "string"}
		}

		err = libbookmarks.RemoveType(db, nil, type_)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--list-types"]; ok && isSet.(bool) {
		types, err := libbookmarks.ListTypes(db)
		if err != nil {
			return err
		}

		for _, type_ := range types {
			fmt.Println(type_)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add"]; ok && isSet.(bool) {
		var newBookmark libbookmarks.Bookmark
		dataRaw, err := arguments.String("DATA")
		if err != nil {
			return ParameterConversionError{"DATA", arguments["DATA"], "string"}
		}

		err = json.Unmarshal([]byte(dataRaw), &newBookmark)
		if err != nil {
			return err
		}

		if newBookmark.Url == "" {
			return IncompleteCompoundParameterError{MissingFields: []string{"Url"}}
		}

		err = libbookmarks.AddBookmark(db, nil, newBookmark)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit"]; ok && isSet.(bool) {
		var data libbookmarks.Bookmark
		dataRaw, err := arguments.String("NEW_DATA")
		if err != nil {
			return ParameterConversionError{"NEW_DATA", arguments["NEW_DATA"], "string"}
		}

		err = json.Unmarshal([]byte(dataRaw), &data)
		if err != nil {
			return err
		}

		err = libbookmarks.EditBookmark(db, nil, data)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-is-read"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		if err != nil {
			return ParameterConversionError{"BOOKMARK_ID", arguments["BOOKMARK_ID"], "string"}
		}

		isReadRaw, err := arguments.String("IS_READ")
		if err != nil {
			return ParameterConversionError{"IS_READ", arguments["IS_READ"], "string"}
		}

		ID, err := strconv.Atoi(IDRaw)
		if err != nil {
			return err
		}

		isRead, err := strconv.ParseBool(isReadRaw)
		if err != nil {
			return err
		}

		err = libbookmarks.EditIsRead(db, nil, ID, isRead)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-title"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		if err != nil {
			return ParameterConversionError{"BOOKMARK_ID", arguments["BOOKMARK_ID"], "string"}
		}

		title, err := arguments.String("TITLE")
		if err != nil {
			return ParameterConversionError{"TITLE", arguments["BOOKMARK_ID"], "string"}
		}

		ID, err := strconv.Atoi(IDRaw)
		if err != nil {
			return err
		}

		err = libbookmarks.EditTitle(db, nil, ID, title)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-url"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		if err != nil {
			return ParameterConversionError{"BOOKMARK_ID", arguments["BOOKMARK_ID"], "string"}
		}

		url, err := arguments.String("URL")
		if err != nil {
			return ParameterConversionError{"URL", arguments["URL"], "string"}
		}

		ID, err := strconv.Atoi(IDRaw)
		if err != nil {
			return err
		}

		err = libbookmarks.EditUrl(db, nil, ID, url)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-type"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		if err != nil {
			return ParameterConversionError{"BOOKMARK_ID", arguments["BOOKMARK_ID"], "string"}
		}

		type_, err := arguments.String("TYPE")
		if err != nil {
			return ParameterConversionError{"TYPE", arguments["TYPE"], "string"}
		}

		ID, err := strconv.Atoi(IDRaw)
		if err != nil {
			return err
		}

		err = libbookmarks.EditType(db, nil, ID, type_)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--edit-is-collection"]; ok && isSet.(bool) {
		IDRaw, err := arguments.String("BOOKMARK_ID")
		if err != nil {
			return ParameterConversionError{"BOOKMARK_ID", arguments["BOOKMARK_ID"], "string"}
		}

		isCollectionRaw, err := arguments.String("IS_COLLECTION")
		if err != nil {
			return ParameterConversionError{"IS_COLLECTION", arguments["IS_COLLECTION"], "string"}
		}

		ID, err := strconv.Atoi(IDRaw)
		if err != nil {
			return err
		}

		isCollection, err := strconv.ParseBool(isCollectionRaw)
		if err != nil {
			return err
		}

		err = libbookmarks.EditIsCollection(db, nil, ID, isCollection)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-tag"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		if err != nil {
			return ParameterConversionError{"TAG", arguments["TAG"], "string"}
		}

		IDRaw, err := arguments.String("BOOKMARK_ID")
		if err != nil {
			return ParameterConversionError{"BOOKMARK_ID", arguments["BOOKMARK_ID"], "string"}
		}

		ID, err := strconv.Atoi(IDRaw)
		if err != nil {
			return err
		}

		err = libbookmarks.AddTag(db, nil, ID, tag)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-tag"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		if err != nil {
			return ParameterConversionError{"TAG", arguments["TAG"], "string"}
		}

		IDRaw, err := arguments.String("BOOKMARK_ID")
		if err != nil {
			return ParameterConversionError{"BOOKMARK_ID", arguments["BOOKMARK_ID"], "string"}
		}

		ID, err := strconv.Atoi(IDRaw)
		if err != nil {
			return err
		}

		err = libbookmarks.RemoveTag(db, nil, ID, tag)
		if err != nil {
			return err
		}
	}

	return nil
}
