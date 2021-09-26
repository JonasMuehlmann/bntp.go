package libbookmarks

import (
	"strconv"
	"strings"
)

// ApplyBookmarkFilters takes an SQL query and adds JOIN and WHERE clauses for filtering.
func ApplyBookmarkFilters(query string, filter BookmarkFilter) string {
	joinFragments := make([]string, 0, 10)
	whereFragments := make([]string, 0, 10)

	if filter.Title != nil {
		whereFragments = append(whereFragments, "WHERE Title LIKE = "+*filter.Title)
	}

	if filter.Url != nil {
		whereFragments = append(whereFragments, "WHERE Url LIKE = "+*filter.Url)
	}

	if filter.IsCollection != nil {
		var valConverted string

		if *filter.IsCollection {
			valConverted = "1"
		} else {
			valConverted = "0"
		}

		whereFragments = append(whereFragments, "WHERE IsCollection = "+valConverted)
	}

	if filter.IsRead != nil {
		var valConverted string

		if *filter.IsRead {
			valConverted = "1"
		} else {
			valConverted = "0"
		}

		whereFragments = append(whereFragments, "WHERE IsRead = "+valConverted)
	}

	if filter.MaxAge != nil {
		whereFragments = append(whereFragments, "WHERE timeAdded BETWEEN DATE('now') AND datetime(DATE('now'),'-'"+strconv.Itoa(*filter.MaxAge)+" days')")
	}

	if filter.Tags != nil {
		joinFragments = append(joinFragments, "INNER JOIN Context ON Context.BookmarkId = Bookmark.Id INNER JOIN Tag ON Tag.Id = Context.TagId")

		var tags []string

		if len(filter.Tags) == 0 {
			tags = []string{"NULL"}
		} else {
			tags = filter.Tags
		}

		whereFragments = append(whereFragments, "WHERE Tag IN ('"+strings.Join(tags, "', '")+"')")
	}

	if filter.Types != nil {
		var types []string

		if len(filter.Types) == 0 {
			types = []string{"NULL"}
		} else {
			types = filter.Types
		}
		whereFragments = append(whereFragments, "WHERE Type IN ('"+strings.Join(types, "', '")+"')")
	}

	joinFragment := strings.Join(joinFragments, " AND ")
	whereFragment := strings.Join(whereFragments, " AND ")

	filteredQuery := query + " " + joinFragment + " " + whereFragment + ";"

	return filteredQuery
}
