package subcommands

import "github.com/docopt/docopt-go"

var usageLink string = `bntp link - Interact with links between documents.

Usage:
    bntp link -h | --help

Options:
    -h --help     Show this screen.
`

func LinkMain() {
	_, _ = docopt.ParseDoc(usageLink)
}
