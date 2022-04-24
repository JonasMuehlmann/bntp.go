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

package domain

import (
	"time"

	"github.com/JonasMuehlmann/optional.go"
)

type Bookmark struct {
	CreatedAt    time.Time                    `json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time                    `json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt    optional.Optional[time.Time] `json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	URL          string                       `json:"url" toml:"url" yaml:"url"`
	Title        optional.Optional[string]    `json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	Tags         []*Tag                       `json:"tags" toml:"tags" yaml:"tags"`
	ID           int64                        `json:"id" toml:"id" yaml:"id"`
	IsCollection bool                         `json:"is_collection,omitempty" toml:"is_collection" yaml:"is_collection,omitempty"`
	IsRead       bool                         `json:"is_read,omitempty" toml:"is_read" yaml:"is_read,omitempty"`
	BookmarkType bool                         `json:"bookmark_type,omitempty" toml:"bookmark_type" yaml:"bookmark_type,omitempty"`
}
