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
	"log"
	"strconv"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/docopt/docopt-go"
	"github.com/jmoiron/sqlx"
)

var usageTag string = `bntp tag - Interact with bookmark's or document's tags.

Usage:
    bntp tag -h | --help
    bntp tag (-i | -e) PATH
    bntp tag (-A | -c | -a | -r | -s) TAG
    bntp tag -R OLD NEW
    bntp (-l | -L)

Options:
    -h --help           Show this screen.
    -i --import         Import a tag structure.
    -e --export         Export a tag structure.
    -A --ambiguous      Check if a tag is ambiguous.
    -c --component      Find ambiguous tag component.
    -a --add            Add a tag.
    -r --remove         Remove a tag.
    -R --rename         Rename a tag.
    -s --shorten        Try to shorten a tag.
    -l --list           List all tags.
    -L --list-short     List all tags, shortened.
`

func TagMain(db *sqlx.DB) {
	arguments, err := docopt.ParseDoc(usageTag)
	helpers.OnError(err, log.Panic)

	// ******************************************************************//
	if _, ok := arguments["--import"]; ok {
		path, err := arguments.String("PATH")
		helpers.OnError(err, log.Panic)

		err = libtags.ImportYML(db, path)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--export"]; ok {
		path, err := arguments.String("PATH")
		helpers.OnError(err, log.Panic)

		err = libtags.ExportYML(db, path)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--ambiguous"]; ok {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, log.Panic)

		isAmbiguous, err := libtags.IsLeafAmbiguous(db, tag)
		helpers.OnError(err, log.Panic)

		println(strconv.FormatBool(isAmbiguous))
		// ******************************************************************//
	} else if _, ok := arguments["--component"]; ok {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, log.Panic)

		index, err := libtags.FindAmbiguousTagComponent(db, tag)
		helpers.OnError(err, log.Panic)

		println(index)
		// ******************************************************************//
	} else if _, ok := arguments["--remove"]; ok {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, log.Panic)

		err = libtags.DeleteTag(db, nil, tag)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--rename"]; ok {
		oldName, err := arguments.String("OLD")
		helpers.OnError(err, log.Panic)

		newName, err := arguments.String("NEW")
		helpers.OnError(err, log.Panic)

		err = libtags.RenameTag(db, nil, oldName, newName)
		helpers.OnError(err, log.Panic)
		// ******************************************************************//
	} else if _, ok := arguments["--shorten"]; ok {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, log.Panic)

		shortened, err := libtags.TryShortenTag(db, tag)
		helpers.OnError(err, log.Panic)

		println(shortened)
		// ******************************************************************//
	} else if _, ok := arguments["--list-short"]; ok {
		tags, err := libtags.ListTags(db)
		helpers.OnError(err, log.Panic)

		for _, tag := range tags {
			println(tag)
		}
	}
}
