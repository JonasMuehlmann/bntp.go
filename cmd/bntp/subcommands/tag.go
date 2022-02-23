package subcommands

import (
	"log"
	"strconv"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libtags"
	"github.com/docopt/docopt-go"
)

var usageTag string = `bntp tag - Interact with bookmark's or document's tags.

Usage:
    bntp tag -h | --help
    bntp tag (-i | -e) PATH
    bntp tag (-A | -c | -a | -r | -s) TAG
    bntp tag -R OLD NEW
    bntp (-l | -L)

Options:
    -h --help           Show this screen.
    -i --import         Import a tag structure.
    -e --export         Export a tag structure.
    -A --ambiguous      Check if a tag is ambiguous.
    -c --component      Find ambiguous tag component.
    -a --add            Add a tag.
    -r --remove         Remove a tag.
    -R --rename         Rename a tag.
    -s --shorten        Try to shorten a tag.
    -l --list           List all tags.
    -L --list-short     List all tags, shortened.
`

func TagMain() {
	arguments, err := docopt.ParseDoc(usageTag)
	if err != nil {
		log.Fatal(err)
	}

	db, err := helpers.GetDefaultDB()
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := arguments["--import"]; ok {
		path, err := arguments.String("PATH")
		if err != nil {
			log.Fatal(err)
		}

		log.Fatal(libtags.ImportYML(db, path))
	} else if _, ok := arguments["--export"]; ok {
		path, err := arguments.String("PATH")
		if err != nil {
			log.Fatal(err)
		}

		log.Fatal(libtags.ExportYML(db, path))
	} else if _, ok := arguments["--ambiguous"]; ok {
		tag, err := arguments.String("TAG")
		if err != nil {
			log.Fatal(err)
		}

		isAmbiguous, err := libtags.IsLeafAmbiguous(db, tag)
		if err != nil {
			log.Fatal(err)
		}

		println(strconv.FormatBool(isAmbiguous))
	} else if _, ok := arguments["--component"]; ok {
		tag, err := arguments.String("TAG")
		if err != nil {
			log.Fatal(err)
		}

		index, err := libtags.FindAmbiguousTagComponent(db, tag)
		if err != nil {
			return
		}

		println(strconv.FormatInt(int64(index), 10))
	} else if _, ok := arguments["--add"]; ok {
		tag, err := arguments.String("TAG")
		if err != nil {
			log.Fatal(err)
		}

		err = libtags.AddTag(db, nil, tag)
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--remove"]; ok {
		tag, err := arguments.String("TAG")
		if err != nil {
			log.Fatal(err)
		}

		err = libtags.DeleteTag(db, nil, tag)
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--rename"]; ok {
		oldName, err := arguments.String("OLD")
		if err != nil {
			log.Fatal(err)
		}

		newName, err := arguments.String("NEW")
		if err != nil {
			log.Fatal(err)
		}

		err = libtags.RenameTag(db, nil, oldName, newName)
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--shorten"]; ok {
		tag, err := arguments.String("TAG")
		if err != nil {
			log.Fatal(err)
		}

		shortened, err := libtags.TryShortenTag(db, tag)
		if err != nil {
			return
		}

		println(shortened)
	} else if _, ok := arguments["--list"]; ok {
		tags, err := libtags.ListTags(db)
		if err != nil {
			log.Fatal(err)
		}

		for _, tag := range tags {
			println(tag)
		}
	} else if _, ok := arguments["--list-short"]; ok {
		tags, err := libtags.ListTags(db)
		if err != nil {
			log.Fatal(err)
		}

		for _, tag := range tags {
			println(tag)
		}
	}
}
