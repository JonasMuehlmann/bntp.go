package subcommands

import "github.com/docopt/docopt-go"

var usageTag string = `bntp tag - Interact with bookmark's or document's tags.

Usage:
    bntp tag -h | --help

Options:
    -h --help     Show this screen.
`

func TagMain() {
	_, _ = docopt.ParseDoc(usageTag)
}
