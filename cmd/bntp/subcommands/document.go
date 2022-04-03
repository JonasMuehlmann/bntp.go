// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the"Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED"AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package subcommands

import (
	"fmt"
	"log"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/internal/libdocuments"
	"github.com/docopt/docopt-go"
)

var usageDocument string = `bntp document - Interact with documents.

Usage:
    bntp document -h | --help
    bntp document (--add-tag | --remove-tag) DOCUMENT_PATH TAG
    bntp document --rename-tag DOCUMENT_PATH OLD_TAG NEW_TAG
    bntp document (--get-tags | --find-tags-line | --find-links-line | --find-backlinks-line) DOCUMENT_PATH
    bntp document --has-tag DOCUMENT_PATH TAG
    bntp document --find-documents-with-tags TAGS
    bntp document (--add-link | --remove-link) DOCUMENT_PATH LINK
    bntp document (--add-backlink | --remove-backlink) DOCUMENT_PATH BACKLINK
    bntp document (--add-doc| --remove-doc) DOCUMENT_PATH
    bntp document --rename-doc OLD_PATH NEW_PATH
    bntp document --change-doc-type DOCUMENT_PATH NEW_TYPE
    bntp document (--add-doc-type | --remove-doc-type) TYPE

Options:
    -h --help                   Show this screen.
    --add-tag                   Add a tag to a document.
    --remove-tag                Remove a tag from a document.
    --rename-tag                Rename a tag in a document.
    --get-tags                  Get tags in a document.
    --find-tags-line            Find the line of the document, which contains tags.
    --has-tags                  Check if the  documentcontains all tags.
    --find-docs-with-tags       Find all documents, which contain the given tags.
    --find-links-lines          Find the lines of the document containing links.
    --find-backlinks-lines      Find the lines of the document containing backlinks.
    --add-link                  Add a link to a document.
    --remove-link               Remove a link from a document.
    --add-backlink              Add a backlink to a document.
    --remove-backlink           Remove a backlink from a document.
    -a --add-doc                Add a document.
    -r --remove-doc             Remove a document.
    -R --rename-doc             Rename a document.
    --change-doc-type           Change a document's type.
    --add-doc-type              Add a new type to give documents.
    --remove-doc-type           Remove a type to give documents.
`

func DocumentMain() {
	arguments, err := docopt.ParseDoc(usageDocument)
	if err != nil {
		log.Fatal(err)
	}

	db, err := helpers.GetDefaultDB()
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := arguments["--add-tag"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		tag, err := arguments.String("TAG")
		if err != nil {
			log.Fatal(err)
		}

		err = libdocuments.AddTagToFile(documentPath, tag)
		if err != nil {
			log.Fatal(err)
		}

		err = libdocuments.AddTag(db, nil, documentPath, tag)
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--remove-tag"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		tag, err := arguments.String("TAG")
		if err != nil {
			log.Fatal(err)
		}

		err = libdocuments.RemoveTagFromFile(documentPath, tag)
		if err != nil {
			log.Fatal(err)
		}

		err = libdocuments.RemoveTag(db, nil, documentPath, tag)
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--rename-tag"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		oldTag, err := arguments.String("OLD_TAG")
		if err != nil {
			log.Fatal(err)
		}

		newTag, err := arguments.String("NEW_TAG")
		if err != nil {
			log.Fatal(err)
		}

		err = libdocuments.RenameTagInFile(documentPath, oldTag, newTag)
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--get-tags"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		tags, err := libdocuments.GetTags(documentPath)
		if err != nil {
			log.Fatal(err)
		}

		for tag := range tags {
			println(tag)
		}
	} else if _, ok := arguments["--find-tags-line"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		index, line, err := libdocuments.FindTagsLine(documentPath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("%v %v", index, line)
	} else if _, ok := arguments["--has-tags"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		tag, err := arguments.String("TAG")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--find-docs-with-tags"]; ok {
		tagsRaw, err := arguments.String("TAGS")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--find-links-lines"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--find-backlinks-lines"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--add-link"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		link, err := arguments.String("LINK")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--remove-link"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		link, err := arguments.String("LINK")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--add-backlink"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		backlink, err := arguments.String("BACKLINK")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--remove-backlink"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		backlink, err := arguments.String("BACKLINK")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--add-doc"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--remove-doc"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--rename-doc"]; ok {
		oldPath, err := arguments.String("OLD_PATH")
		if err != nil {
			log.Fatal(err)
		}

		newPath, err := arguments.String("NEW_PATH")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--change-doc-type"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--add-doc-type"]; ok {
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			log.Fatal(err)
		}

		type_, err := arguments.String("TYPE")
		if err != nil {
			log.Fatal(err)
		}
	} else if _, ok := arguments["--remove-doc-type"]; ok {
		type_, err := arguments.String("TYPE")
		if err != nil {
			log.Fatal(err)
		}
	}
}
