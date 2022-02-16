package subcommands

import "github.com/docopt/docopt-go"

var usageDocument string = `bntp document - Interact with documents.

Usage:
    bntp document -h | --help

Options:
    -h --help     Show this screen.
`

func DocumentMain() {
	_, _ = docopt.ParseDoc(usageDocument)
}
