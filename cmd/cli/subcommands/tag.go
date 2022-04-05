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
	"fmt"
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
    bntp tag (-l | -L)

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

func TagMain(db *sqlx.DB, exiter func(int)) {
	arguments, err := docopt.ParseDoc(usageTag)
	helpers.OnError(err, helpers.MakeFatalLogger(exiter))

	// ******************************************************************//
	if isSet, ok := arguments["--import"]; ok && isSet.(bool) {
		path, err := arguments.String("PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libtags.ImportYML(db, path)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--export"]; ok && isSet.(bool) {
		path, err := arguments.String("PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libtags.ExportYML(db, path)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--ambiguous"]; ok && isSet.(bool) {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		isAmbiguous, err := libtags.IsLeafAmbiguous(db, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		fmt.Println(strconv.FormatBool(isAmbiguous))
		// ******************************************************************//
	} else if isSet, ok := arguments["--component"]; ok && isSet.(bool) {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		index, component, err := libtags.FindAmbiguousTagComponent(db, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		fmt.Println(index, component)
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove"]; ok && isSet.(bool) {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libtags.DeleteTag(db, nil, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--rename"]; ok && isSet.(bool) {
		oldName, err := arguments.String("OLD")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		newName, err := arguments.String("NEW")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libtags.RenameTag(db, nil, oldName, newName)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--shorten"]; ok && isSet.(bool) {
		tag, err := arguments.String("TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		shortened, err := libtags.TryShortenTag(db, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		fmt.Println(shortened)
		// ******************************************************************//
	} else if isSet, ok := arguments["--list"]; ok && isSet.(bool) {
		tags, err := libtags.ListTags(db)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		for _, tag := range tags {
			fmt.Println(tag)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--list-short"]; ok && isSet.(bool) {
		tags, err := libtags.ListTagsShortened(db)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		for _, tag := range tags {
			fmt.Println(tag)
		}
	}
}
