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

var usageBookmark string = `bntp bookmark - Interact with bookmarks.

Usage:
    bntp bookmark -h | --help
    bntp bookmark -l [--filter]

Options:
    -h --help               Show this screen.
    -i --import             Import bookmarks.
    -e --export             Export bookmarks.
    -l --list               List bookmarks.
    --filter                A filter to apply when searching for bookmarks.
    --add-type              Add a bookmark type.
    --remove-type           Remove a bookmark type.
    --list-types            List the bookmark types.
    -a --add                Add a bookmark.
    -r --rmove              Remove a bookmark.
    -e --edit               Edit a bookmark.
    --edit-is-read          Change the is read status of a bookmark.
    --edit-title            Change the title of a bookmark.
    --edit-url              Change the url of a bookmark.
    --edit-type             Change the type of a bookmark.
    --edit-is-collection    Change the is collection status of a bookmark.
    --add-tag               Add a tag to a bookmark.
    --remove-tag            Remove a tag from a bookmark.
`

func BookmarkMain() {
	_, _ = docopt.ParseDoc(usageBookmark)
}
