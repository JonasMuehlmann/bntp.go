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
	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
)

// Bookmark is a code side representation of DB bookmarks.
type Bookmark struct {
	Title        helpers.Optional[string] `json:"title" db:"Title"`
	Url          string                   `json:"url" db:"Url"`
	TimeAdded    string                   `json:"time_added" db:"TimeAdded"`
	Type         helpers.Optional[string] `json:"type" db:"Type"`
	Id           int                      `json:"id" db:"Id"`
	IsRead       bool                     `json:"is_read" db:"IsRead"`
	IsCollection helpers.Optional[bool]   `json:"is_collection" db:"IsCollection"`
	Tags         []string
}
