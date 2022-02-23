package subcommands

import (
	"log"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/liblinks"
	"github.com/docopt/docopt-go"
)

var usageLink string = `bntp link - Interact with links between documents.

Usage:
    bntp link -h | --help
    bntp link (-a | -r) SRC DEST
    bntp link -l SRC
    bntp link -L DEST

Options:
    -h --help       Show this screen.
    -a --add        Add a link
    -r --remove     Remove a link
    -l --list       List all links
    -L --list-back  List backlinks
`

func LinkMain() {
	arguments, err := docopt.ParseDoc(usageLink)
	if err != nil {
		log.Fatal(err)
	}

	db, err := helpers.GetDefaultDB()
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := arguments["--add"]; ok {
		source, err := arguments.String("SRC")
		if err != nil {
			log.Fatal(err)
		}

		destination, err := arguments.String("DEST")
		if err != nil {
			log.Fatal(err)
		}

		log.Fatal(liblinks.AddLink(db, nil, source, destination))
	} else if _, ok := arguments["--remove"]; ok {
		source, err := arguments.String("SRC")
		if err != nil {
			log.Fatal(err)
		}

		destination, err := arguments.String("DEST")
		if err != nil {
			log.Fatal(err)
		}

		log.Fatal(liblinks.RemoveLink(db, nil, source, destination))
	} else if _, ok := arguments["--list"]; ok {
		source, err := arguments.String("SRC")
		if err != nil {
			log.Fatal(err)
		}

		log.Fatal(liblinks.ListLinks(db, source))
	} else if _, ok := arguments["--list-back"]; ok {
		destination, err := arguments.String("DEST")
		if err != nil {
			log.Fatal(err)
		}

		log.Fatal(liblinks.ListBacklinks(db, destination))
	}
}
