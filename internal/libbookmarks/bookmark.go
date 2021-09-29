package libbookmarks

import "database/sql"

// Bookmark is a code side representation of DB bookmarks.
type Bookmark struct {
	Title        sql.NullString `db:"Title"`
	Url          string         `db:"Url"`
	TimeAdded    string         `db:"TimeAdded"`
	Type         sql.NullString `db:"Type"`
	Tags         []string
	Id           int          `db:"Id"`
	IsRead       bool         `db:"IsRead"`
	IsCollection sql.NullBool `db:"IsCollection"`
}
