package libbookmarks

// BookmarkFilter is used to filter bookmark retrieval.
// Assigning nil to a field does not trigger filtering for that field.
// Empty values("" or []T{}) filter for unset data (e.g. WHERE Column = NULL).
type BookmarkFilter struct {
	Title        *string
	Url          *string
	IsRead       *bool
	IsCollection *bool
	MaxAge       *int
	Types        []string
	Tags         []string
}

// BookmarkFilterInboxed is a BookmarkFilter configured to filter "Inboxed" bookmarks.
// "Inboxed" bookmarks don't have a Type and are not tagged.
var BookmarkFilterInboxed BookmarkFilter = BookmarkFilter{Types: []string{}, Tags: []string{}}
