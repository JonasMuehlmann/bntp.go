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

package libbookmarks

import "github.com/JonasMuehlmann/optional.go"

// BookmarkFilter is used to filter bookmark retrieval.
// Assigning nil to a field does not trigger filtering for that field.
// Empty values("" or []T{}) filter for unset data (e.g. WHERE Column = NULL).
type BookmarkFilter struct {
	Title        optional.Optional[string]
	Url          optional.Optional[string]
	IsRead       optional.Optional[bool]
	IsCollection optional.Optional[bool]
	MaxAge       optional.Optional[int]
	Types        optional.Optional[[]string]
	Tags         optional.Optional[[]string]
}

// BookmarkFilterInboxed is a BookmarkFilter configured to filter "Inboxed" bookmarks.
// "Inboxed" bookmarks don't have a Type and are not tagged.
var BookmarkFilterInboxed BookmarkFilter = BookmarkFilter{Types: optional.Optional[[]string]{}, Tags: optional.Optional[[]string]{}}
