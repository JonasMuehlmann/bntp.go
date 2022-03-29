// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package subcommands

import "github.com/docopt/docopt-go"

var usageDocument string = `bntp document - Interact with documents.

Usage:
    bntp document -h | --help

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
	_, _ = docopt.ParseDoc(usageDocument)
}
