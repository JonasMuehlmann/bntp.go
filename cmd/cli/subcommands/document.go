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
	"github.com/jmoiron/sqlx"
)

var usageDocument string = `bntp document - Interact with documents.

Usage:
    bntp document -h | --help
    bntp document (--add-tag | --remove-tag) DOCUMENT_PATH TAG
    bntp document --rename-tag DOCUMENT_PATH OLD_TAG NEW_TAG
    bntp document (--get-tags | --find-tags-line | --find-links-lines | --find-backlinks-lines) DOCUMENT_PATH
    bntp document --has-tags DOCUMENT_PATH TAGS...
    bntp document --find-docs-with-tags TAGS...
    bntp document (--add-link | --remove-link) DOCUMENT_PATH LINK
    bntp document (--add-backlink | --remove-backlink) DOCUMENT_PATH BACKLINK
    bntp document (--add-doc| --remove-doc) DOCUMENT_PATH TYPE
    bntp document --rename-doc OLD_PATH NEW_PATH
    bntp document --change-doc-type DOCUMENT_PATH TYPE
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

func DocumentMain(db *sqlx.DB, exiter func(int)) {
	arguments, err := docopt.ParseDoc(usageDocument)
	helpers.OnError(err, helpers.MakeFatalLogger(exiter))

	// ******************************************************************//
	if isSet, ok := arguments["--add-tag"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.AddTagToFile(documentPath, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// FIX: This returns an unclear error message: no rows in result set
		err = libdocuments.AddTag(db, nil, documentPath, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-tag"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.RemoveTagFromFile(documentPath, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.RemoveTag(db, nil, documentPath, tag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--rename-tag"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for tag
		oldTag, err := arguments.String("OLD_TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for tag
		newTag, err := arguments.String("NEW_TAG")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.RenameTagInFile(documentPath, oldTag, newTag)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--get-tags"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		tags, err := libdocuments.GetTags(documentPath)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		for _, tag := range tags {
			fmt.Println(tag)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--find-tags-line"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		index, line, err := libdocuments.FindTagsLine(documentPath)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		fmt.Printf("%v %v\n", index, line)
		// ******************************************************************//
	} else if isSet, ok := arguments["--has-tags"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for tag
		tagsRaw, ok := arguments["TAGS"].([]string)
		if !ok {
			log.Println("Missing parameter TAGS")
			exiter(1)
		}

		hasTag, err := libdocuments.HasTags(documentPath, tagsRaw)
		// REFACTOR: Use regular if err != nil {return err}
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		fmt.Println(hasTag)
		// ******************************************************************//
	} else if isSet, ok := arguments["--find-docs-with-tags"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for tag
		tagsRaw, ok := arguments["TAGS"].([]string)
		if !ok {
			log.Println("Missing parameter TAGS")
			exiter(1)
		}

		documents, err := libdocuments.FindDocumentsWithTags(db, tagsRaw)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		for _, document := range documents {
			fmt.Println(document)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--find-links-lines"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		start, end, links, err := libdocuments.FindLinksLines(documentPath)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		fmt.Printf("%v %v\n", start, end)
		for _, link := range links {
			fmt.Println(link)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--find-backlinks-lines"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		start, end, backlinks, err := libdocuments.FindBacklinksLines(documentPath)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		fmt.Printf("%v %v\n", start, end)
		for _, link := range backlinks {
			fmt.Println(link)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-link"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		link, err := arguments.String("LINK")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.AddLink(documentPath, link)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-link"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for document
		link, err := arguments.String("LINK")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.RemoveLink(documentPath, link)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-backlink"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for document
		backlink, err := arguments.String("BACKLINK")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.AddBacklink(documentPath, backlink)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-backlink"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for document
		backlink, err := arguments.String("BACKLINK")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.RemoveBacklink(documentPath, backlink)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-doc"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for Type
		type_, err := arguments.String("TYPE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.AddDocument(db, nil, documentPath, type_)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-doc"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.RemoveDocument(db, nil, documentPath)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--rename-doc"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		oldPath, err := arguments.String("OLD_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		newPath, err := arguments.String("NEW_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.RenameDocument(db, nil, oldPath, newPath)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--change-doc-type"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		// TODO: Allow passing string and id for Type
		type_, err := arguments.String("TYPE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.ChangeDocumentType(db, nil, documentPath, type_)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-doc-type"]; ok && isSet.(bool) {
		type_, err := arguments.String("TYPE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.AddType(db, nil, type_)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-doc-type"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for Type
		type_, err := arguments.String("TYPE")
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))

		err = libdocuments.RemoveType(db, nil, type_)
		helpers.OnError(err, helpers.MakeFatalLogger(exiter))
	}
}
