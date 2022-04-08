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

func DocumentMain(db *sqlx.DB) error {
	arguments, err := docopt.ParseDoc(usageDocument)
	if err != nil {
		return err
	}

	// ******************************************************************//
	if isSet, ok := arguments["--add-tag"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		if err != nil {
			return ParameterConversionError{"TAG", arguments["TAG"], "string"}
		}

		err = libdocuments.AddTagToFile(documentPath, tag)
		if err != nil {
			return err
		}

		// FIX: This returns an unclear error message: no rows in result set
		err = libdocuments.AddTag(db, nil, documentPath, tag)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-tag"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		// TODO: Allow passing string and id for tag
		tag, err := arguments.String("TAG")
		if err != nil {
			return ParameterConversionError{"TAG", arguments["TAG"], "string"}
		}

		err = libdocuments.RemoveTagFromFile(documentPath, tag)
		if err != nil {
			return err
		}

		err = libdocuments.RemoveTag(db, nil, documentPath, tag)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--rename-tag"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		// TODO: Allow passing string and id for tag
		oldTag, err := arguments.String("OLD_TAG")
		if err != nil {
			return ParameterConversionError{"OLD_TAG", arguments["OLD_TAG"], "string"}
		}

		// TODO: Allow passing string and id for tag
		newTag, err := arguments.String("NEW_TAG")
		if err != nil {
			return ParameterConversionError{"NEW_TAG", arguments["NEW_TAG"], "string"}
		}

		err = libdocuments.RenameTagInFile(documentPath, oldTag, newTag)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--get-tags"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		tags, err := libdocuments.GetTags(documentPath)
		if err != nil {
			return err
		}

		for _, tag := range tags {
			fmt.Println(tag)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--find-tags-line"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		index, line, err := libdocuments.FindTagsLine(documentPath)
		if err != nil {
			return err
		}

		fmt.Printf("%v %v\n", index, line)
		// ******************************************************************//
	} else if isSet, ok := arguments["--has-tags"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		// TODO: Allow passing string and id for tag
		tagsRaw, _ := arguments["TAGS"].([]string)

		hasTag, err := libdocuments.HasTags(documentPath, tagsRaw)
		if err != nil {
			return err
		}

		fmt.Println(hasTag)
		// ******************************************************************//
	} else if isSet, ok := arguments["--find-docs-with-tags"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for tag
		tagsRaw, _ := arguments["TAGS"].([]string)

		documents, err := libdocuments.FindDocumentsWithTags(db, tagsRaw)
		if err != nil {
			return err
		}

		for _, document := range documents {
			fmt.Println(document)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--find-links-lines"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		start, end, links, err := libdocuments.FindLinksLines(documentPath)
		if err != nil {
			return err
		}

		fmt.Printf("%v %v\n", start, end)
		for _, link := range links {
			fmt.Println(link)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--find-backlinks-lines"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		start, end, backlinks, err := libdocuments.FindBacklinksLines(documentPath)
		if err != nil {
			return err
		}

		fmt.Printf("%v %v\n", start, end)
		for _, link := range backlinks {
			fmt.Println(link)
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-link"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		link, err := arguments.String("LINK")
		if err != nil {
			return ParameterConversionError{"LINK", arguments["LINK"], "string"}
		}

		err = libdocuments.AddLink(documentPath, link)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-link"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		// TODO: Allow passing string and id for document
		link, err := arguments.String("LINK")
		if err != nil {
			return ParameterConversionError{"LINK", arguments["LINK"], "string"}
		}

		err = libdocuments.RemoveLink(documentPath, link)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-backlink"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		// TODO: Allow passing string and id for document
		backlink, err := arguments.String("BACKLINK")
		if err != nil {
			return ParameterConversionError{"BACKLINK", arguments["BACKLINK"], "string"}
		}

		err = libdocuments.AddBacklink(documentPath, backlink)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-backlink"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		// TODO: Allow passing string and id for document
		backlink, err := arguments.String("BACKLINK")
		if err != nil {
			return ParameterConversionError{"BACKLINK", arguments["BACKLINK"], "string"}
		}

		err = libdocuments.RemoveBacklink(documentPath, backlink)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-doc"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		// TODO: Allow passing string and id for Type
		type_, err := arguments.String("TYPE")
		if err != nil {
			return ParameterConversionError{"TYPE", arguments["TYPE"], "string"}
		}

		err = libdocuments.AddDocument(db, nil, documentPath, type_)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-doc"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		err = libdocuments.RemoveDocument(db, nil, documentPath)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--rename-doc"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		oldPath, err := arguments.String("OLD_PATH")
		if err != nil {
			return ParameterConversionError{"OLD_PATH", arguments["OLD_PATH"], "string"}
		}

		newPath, err := arguments.String("NEW_PATH")
		if err != nil {
			return ParameterConversionError{"NEW_PATH", arguments["NEW_PATH"], "string"}
		}

		err = libdocuments.RenameDocument(db, nil, oldPath, newPath)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--change-doc-type"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for document
		documentPath, err := arguments.String("DOCUMENT_PATH")
		if err != nil {
			return ParameterConversionError{"DOCUMENT_PATH", arguments["DOCUMENT_PATH"], "string"}
		}

		// TODO: Allow passing string and id for Type
		type_, err := arguments.String("TYPE")
		if err != nil {
			return ParameterConversionError{"TYPE", arguments["TYPE"], "string"}
		}

		err = libdocuments.ChangeDocumentType(db, nil, documentPath, type_)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--add-doc-type"]; ok && isSet.(bool) {
		type_, err := arguments.String("TYPE")
		if err != nil {
			return ParameterConversionError{"TYPE", arguments["TYPE"], "string"}
		}

		err = libdocuments.AddType(db, nil, type_)
		if err != nil {
			return err
		}
		// ******************************************************************//
	} else if isSet, ok := arguments["--remove-doc-type"]; ok && isSet.(bool) {
		// TODO: Allow passing string and id for Type
		type_, err := arguments.String("TYPE")
		if err != nil {
			return ParameterConversionError{"TYPE", arguments["TYPE"], "string"}
		}

		err = libdocuments.RemoveType(db, nil, type_)
		if err != nil {
			return err
		}
	}

	return nil
}
