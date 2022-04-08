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

func TagMain(db *sqlx.DB) error {
	arguments, err := docopt.ParseDoc(usageTag)
	if err != nil {
		return err
	}

	// ******************************************************************//
	if isSet, ok := arguments["--import"]; ok && isSet.(bool) {
		path, err := arguments.String("PATH")
		if err != nil {
			return ParameterConversionError{"PATH", arguments["PATH"], "string"}
		}

		err = libtags.ImportYML(db, path)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--export"]; ok && isSet.(bool) {
		path, err := arguments.String("PATH")
		if err != nil {
			return ParameterConversionError{"PATH", arguments["PATH"], "string"}
		}

		err = libtags.ExportYML(db, path)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--ambiguous"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		if err != nil {
			return ParameterConversionError{"TAG", arguments["TAG"], "string"}
		}

		isAmbiguous, err := libtags.IsLeafAmbiguous(db, tag)
		if err != nil {
			return err
		}

		fmt.Println(strconv.FormatBool(isAmbiguous))
		// ******************************************************************//
	} else if isSet, ok := arguments["--component"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		if err != nil {
			return ParameterConversionError{"TAG", arguments["TAG"], "string"}
		}

		index, component, err := libtags.FindAmbiguousTagComponent(db, tag)
		if err != nil {
			return err
		}

		fmt.Println(index, component)
		// ******************************************************************//
	} else if isSet, ok := arguments["--add"]; ok && isSet.(bool) {
		tag, err := arguments.String("TAG")
		if err != nil {
			return ParameterConversionError{"TAG", arguments["TAG"], "string"}
		}

		err = libtags.AddTag(db, nil, tag)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		if err != nil {
			return ParameterConversionError{"TAG", arguments["TAG"], "string"}
		}

		err = libtags.DeleteTag(db, nil, tag)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--rename"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for tag
		oldName, err := arguments.String("OLD")
		if err != nil {
			return ParameterConversionError{"OLD", arguments["OLD"], "string"}
		}

		newName, err := arguments.String("NEW")
		if err != nil {
			return ParameterConversionError{"NEW", arguments["NEW"], "string"}
		}

		err = libtags.RenameTag(db, nil, oldName, newName)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--shorten"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		if err != nil {
			return ParameterConversionError{"TAG", arguments["TAG"], "string"}
		}

		shortened, err := libtags.TryShortenTag(db, tag)
		if err != nil {
			return err
		}

		fmt.Println(shortened)
		// ******************************************************************//
	} else if isSet, ok := arguments["--list"]; ok && isSet.(bool) {
		tags, err := libtags.ListTags(db)
		if err != nil {
			return err
		}

		for _, tag := range tags {
			fmt.Println(tag)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--list-short"]; ok && isSet.(bool) {
		tags, err := libtags.ListTagsShortened(db)
		if err != nil {
			return err
		}

		for _, tag := range tags {
			fmt.Println(tag)
		}
	}

	return nil
}
