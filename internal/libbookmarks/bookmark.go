package libbookmarks

type Bookmark struct {
	Id           int
	IsRead       bool
	IsCollection bool
	Title        string
	Url          string
	TimeAdded    string
	Type         string
	Tags         []string
}
