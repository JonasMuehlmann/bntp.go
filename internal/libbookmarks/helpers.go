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

import (
	"strconv"
	"strings"
)

// REFACTOR: This might need an overhaul

// ApplyBookmarkFilters takes an SQL query and adds JOIN and WHERE clauses for filtering.
func ApplyBookmarkFilters(query string, filter BookmarkFilter) string {
	joinFragments := make([]string, 0, 10)
	whereFragments := make([]string, 0, 10)

	if filter.Title.HasValue {
		whereFragments = append(whereFragments, "Title LIKE '"+*&filter.Title.Wrappee+"'")
	}

	if filter.Url.HasValue {
		whereFragments = append(whereFragments, "Url LIKE '"+filter.Url.Wrappee+"'")
	}

	if filter.IsCollection.HasValue {
		var valConverted string

		if filter.IsCollection.Wrappee {
			valConverted = "1"
		} else {
			valConverted = "0"
		}

		whereFragments = append(whereFragments, "IsCollection = "+valConverted)
	}

	if filter.IsRead.HasValue {
		var valConverted string

		if filter.IsRead.Wrappee {
			valConverted = "1"
		} else {
			valConverted = "0"
		}

		whereFragments = append(whereFragments, "IsRead = "+valConverted)
	}

	if filter.MaxAge.HasValue {
		whereFragments = append(whereFragments, "timeAdded BETWEEN DATE('now') AND datetime(DATE('now'),'-'"+strconv.Itoa(filter.MaxAge.Wrappee)+" days')")
	}

	if filter.Tags.HasValue {
		joinFragments = append(joinFragments, "INNER JOIN Context ON Context.BookmarkId = Bookmark.Id INNER JOIN Tag ON Tag.Id = Context.TagId")

		var tags []string

		if len(filter.Tags.Wrappee) == 0 {
			tags = []string{"NULL"}
		} else {
			tags = filter.Tags.Wrappee
		}

		whereFragments = append(whereFragments, "Tag IN ('"+strings.Join(tags, "', '")+"')")
	}

	if filter.Types.HasValue {
		var types []string

		if len(filter.Types.Wrappee) == 0 {
			types = []string{"NULL"}
		} else {
			types = filter.Types.Wrappee
		}
		whereFragments = append(whereFragments, "BookmarkType.Type IN ('"+strings.Join(types, "', '")+"')")
	}

	joinFragment := strings.Join(joinFragments, " AND ")
	whereFragment := strings.Join(whereFragments, " AND ")

	filteredQuery := query + " " + joinFragment

	if len(whereFragments) > 0 {
		filteredQuery += " WHERE " + whereFragment
	}

	filteredQuery += ";"

	return filteredQuery
}
