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

	if filter.Title != nil {
		whereFragments = append(whereFragments, "Title LIKE '"+*filter.Title+"'")
	}

	if filter.Url != nil {
		whereFragments = append(whereFragments, "Url LIKE '"+*filter.Url+"'")
	}

	if filter.IsCollection != nil {
		var valConverted string

		if *filter.IsCollection {
			valConverted = "1"
		} else {
			valConverted = "0"
		}

		whereFragments = append(whereFragments, "IsCollection = "+valConverted)
	}

	if filter.IsRead != nil {
		var valConverted string

		if *filter.IsRead {
			valConverted = "1"
		} else {
			valConverted = "0"
		}

		whereFragments = append(whereFragments, "IsRead = "+valConverted)
	}

	if filter.MaxAge != nil {
		whereFragments = append(whereFragments, "timeAdded BETWEEN DATE('now') AND datetime(DATE('now'),'-'"+strconv.Itoa(*filter.MaxAge)+" days')")
	}

	if filter.Tags != nil {
		joinFragments = append(joinFragments, "INNER JOIN Context ON Context.BookmarkId = Bookmark.Id INNER JOIN Tag ON Tag.Id = Context.TagId")

		var tags []string

		if len(filter.Tags) == 0 {
			tags = []string{"NULL"}
		} else {
			tags = filter.Tags
		}

		whereFragments = append(whereFragments, "Tag IN ('"+strings.Join(tags, "', '")+"')")
	}

	if filter.Types != nil {
		var types []string

		if len(filter.Types) == 0 {
			types = []string{"NULL"}
		} else {
			types = filter.Types
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
