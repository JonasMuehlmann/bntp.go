package main

import (
	"log"
	"os"

	"github.com/JonasMuehlmann/bntp.go/cmd/bntp/subcommands"
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

	log.Println(subcommand)

	switch subcommand {
	case "bookmark":
		subcommands.BookmarkMain()
	case "link":
		subcommands.LinkMain()
	case "tag":
		subcommands.TagMain()
	case "document":
		subcommands.DocumentMain()
	default:
		log.Fatal("Invalid subcommand")
	}
}
