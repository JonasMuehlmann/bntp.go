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
	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/liblinks"
	"github.com/docopt/docopt-go"
	"github.com/jmoiron/sqlx"
)

var usageLink string = `bntp link - Interact with links between documents.

Usage:
    bntp link -h | --help
    bntp link (-a | -r) SRC DEST
    bntp link -l SRC
    bntp link -L DEST

Options:
    -h --help       Show this screen.
    -a --add        Add a link.
    -r --remove     Remove a link.
    -l --list       List all links.
    -L --list-back  List backlinks.
`

func LinkMain(db *sqlx.DB, exiter func(int)) {
	arguments, err := docopt.ParseDoc(usageLink)
	helpers.OnError(err, helpers.MakeFatalLogger(exiter))

	// ******************************************************************//
	if _, ok := arguments["--add"]; ok {
		source, err := arguments.String("SRC")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		destination, err := arguments.String("DEST")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = liblinks.AddLink(db, nil, source, destination)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if _, ok := arguments["--remove"]; ok {
		source, err := arguments.String("SRC")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		destination, err := arguments.String("DEST")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = liblinks.RemoveLink(db, nil, source, destination)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if _, ok := arguments["--list"]; ok {
		source, err := arguments.String("SRC")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		links, err := liblinks.ListLinks(db, source)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		for _, link := range links {
			println(link)
		}
		// ******************************************************************//
	} else if _, ok := arguments["--list-back"]; ok {
		destination, err := arguments.String("DEST")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		backlinks, err := liblinks.ListBacklinks(db, destination)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		for _, backlink := range backlinks {
			println(backlink)
		}
	}
}
