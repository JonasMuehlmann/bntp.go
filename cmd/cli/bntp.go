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

package main

import (
	"log"
	"os"

	"github.com/JonasMuehlmann/bntp.go/cmd/cli/subcommands"
	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/docopt/docopt-go"
	_ "github.com/mattn/go-sqlite3"
)

var usage string = `bntp - The all-in-one productivity system bookmarks, notes, todos, projects, etc.

Usage:
    bntp COMMAND ARG...
    bntp -h | --help
    bntp --version

Available subcommands:
    bookmark
    link
    tag
    document

Options:
    -h --help     Show this screen.
    --version     Show version.`

func main() {
	parser := &docopt.Parser{
		HelpHandler:  docopt.PrintHelpOnly,
		OptionsFirst: true,
	}

	arguments, _ := parser.ParseArgs(usage, os.Args[1:], "")

	options := []string{"--help", "--version"}

	for _, option := range options {
		if hasOption, _ := arguments.Bool(option); hasOption {
			return
		}
	}

	subcommand, err := arguments.String("COMMAND")
	if err != nil {
		// COMMAND not specified
		return
	}

	db, err := helpers.GetDefaultDB()
	helpers.OnError(err, log.Panic)

	switch subcommand {
	case "bookmark":
		subcommands.BookmarkMain(db)
	case "link":
		subcommands.LinkMain(db)
	case "tag":
		subcommands.TagMain(db)
	case "document":
		subcommands.DocumentMain(db)
	default:
		log.Panic("Invalid subcommand")
	}
}
