package libbookmarks

type BookmarkFilter struct {
	Title        string
	Url          string
	IsRead       *bool
	IsCollection *bool
	MaxAge       *int
	Types        []string
	Tags         []string
}

var BookmarkFilterInboxed BookmarkFilter = BookmarkFilter{Title: "", Types: []string{}, Tags: []string{}}
