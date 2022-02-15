package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	usage := `bntp - Personal all-in-one productivity system including components like bookmarks, notes, todos, projects, etc. .

Usage:
  bntp -h | --help
  bntp --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.ParseDoc(usage)

	fmt.Println(arguments)
}
