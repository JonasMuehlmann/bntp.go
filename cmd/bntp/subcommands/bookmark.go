package subcommands

import "github.com/docopt/docopt-go"

var usageBookmark string = `bntp bookmark - Interact with bookmarks.

Usage:
    bntp bookmark -h | --help

Options:
    -h --help     Show this screen.
`

func BookmarkMain() {
	_, _ = docopt.ParseDoc(usageBookmark)
}
