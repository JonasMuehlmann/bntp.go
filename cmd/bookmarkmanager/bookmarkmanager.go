package main

import (
	"github.com/docopt/docopt-go"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	usage := `Bookmarkmanager.

Usage:
  bookmarkmanager -h | --help
  bookmarkmanager --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.ParseDoc(usage)
}
