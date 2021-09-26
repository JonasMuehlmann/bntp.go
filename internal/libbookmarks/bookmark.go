package libbookmarks

// Bookmark is a code side representation of DB bookmarks.
type Bookmark struct {
	Title        string
	Url          string
	TimeAdded    string
	Type         string
	Tags         []string
	Id           int
	IsRead       bool
	IsCollection bool
}
