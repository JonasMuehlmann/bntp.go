package libbookmarks

import "database/sql"

// Bookmark is a code side representation of DB bookmarks.
type Bookmark struct {
	Title        sql.NullString `json:"title" db:"Title"`
	Url          string         `json:"url" db:"Url"`
	TimeAdded    string         `json:"time_added" db:"TimeAdded"`
	Type         sql.NullString `json:"type" db:"Type"`
	Id           int            `json:"id" db:"Id"`
	IsRead       bool           `json:"is_read" db:"IsRead"`
	IsCollection sql.NullBool   `json:"is_collection" db:"IsCollection"`
	Tags         []string
}
