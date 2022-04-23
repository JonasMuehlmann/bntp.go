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

type Document struct {
	CreatedAt           time.Time                    `json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt           time.Time                    `json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Path                string                       `json:"path" toml:"path" yaml:"path"`
	DeletedAt           optional.Optional[time.Time] `json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	DocumentType        optional.Optional[string]    `json:"document_type" toml:"document_type" yaml:"document_type"`
	ID                  int64                        `json:"id" toml:"id" yaml:"id"`
	Tags                []*Tag                       `json:"Tags" toml:"Tags" yaml:"Tags"`
	LinkedDocuments     []*Document                  `json:"linked_documents" toml:"linked_documents" yaml:"linked_documents"`
	BacklinkedDocuments []*Document                  `json:"backlinked_documents" toml:"backlinked_documents" yaml:"backlinked_documents"`
}
