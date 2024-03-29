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

package main

import (
	"os"
	"text/template"
	"time"

	"github.com/JonasMuehlmann/bntp.go/tools"
	"github.com/JonasMuehlmann/optional.go"
)

type Tag struct {
	ID            int64   `json:"id" toml:"id" yaml:"id"`
	ParentPathIDs []int64 `json:"parentPathIDs" toml:"parentPathIDs" yaml:"parentPathIDs"`
	Tag           string  `json:"tag" toml:"tag" yaml:"tag"`
	SubtagIDs     []int64 `json:"subtagsIDs" toml:"subtagsIDs" yaml:"subtagsIDs"`
}

type Bookmark struct {
	CreatedAt    time.Time                    `json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time                    `json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt    optional.Optional[time.Time] `json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	URL          string                       `json:"url" toml:"url" yaml:"url"`
	Title        optional.Optional[string]    `json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	TagIDs       []int64                      `json:"tagIDs" toml:"tagIDs" yaml:"tagIDs"`
	ID           int64                        `json:"id" toml:"id" yaml:"id"`
	IsCollection bool                         `json:"is_collection,omitempty" toml:"is_collection" yaml:"is_collection,omitempty"`
	IsRead       bool                         `json:"is_read,omitempty" toml:"is_read" yaml:"is_read,omitempty"`
	BookmarkType optional.Optional[string]    `json:"bookmark_type,omitempty" toml:"bookmark_type" yaml:"bookmark_type,omitempty"`
}

type Document struct {
	CreatedAt              time.Time                    `json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt              time.Time                    `json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt              optional.Optional[time.Time] `json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Path                   string                       `json:"path" toml:"path" yaml:"path"`
	DocumentType           optional.Optional[string]    `json:"document_type" toml:"document_type" yaml:"document_type"`
	TagIDs                 []int64                      `json:"tagIDs" toml:"tagIDs" yaml:"tagIDs"`
	LinkedDocumentIDs      []int64                      `json:"linked_documentIDs" toml:"linked_documentIDs" yaml:"linked_documentIDs"`
	BacklinkedDocumentsIDs []int64                      `json:"backlinked_documentIDs" toml:"backlinked_documentIDs" yaml:"backlinked_documentIDs"`
	ID                     int64                        `json:"id" toml:"id" yaml:"id"`
}

var entities = []any{Document{}, Tag{}, Bookmark{}}

// TODO: Add this to template
// func (updater *DocumentUpdater) ApplyToModel(documentModel *Document) {
// 	if updater.ID.HasValue {
// 		model.ApplyUpdater(&(*documentModel).ID, updater.ID.Wrappee)
// 	}
// 	if updater.Path.HasValue {
// 		model.ApplyUpdater(&(*documentModel).Path, updater.Path.Wrappee)
// 	}
// 	if updater.DocumentTypeID.HasValue {
// 		model.ApplyUpdater(&(*documentModel).DocumentTypeID, updater.DocumentTypeID.Wrappee)
// 	}
// 	if updater.CreatedAt.HasValue {
// 		model.ApplyUpdater(&(*documentModel).CreatedAt, updater.CreatedAt.Wrappee)
// 	}
// 	if updater.UpdatedAt.HasValue {
// 		model.ApplyUpdater(&(*documentModel).UpdatedAt, updater.UpdatedAt.Wrappee)
// 	}
// 	if updater.DeletedAt.HasValue {
// 		model.ApplyUpdater(&(*documentModel).DeletedAt, updater.DeletedAt.Wrappee)
// 	}

// }

func main() {
	for _, entity := range entities {
		entityStruct := tools.NewStructModel(entity)

		tmplRaw, err := os.ReadFile("templates/" + tools.LowercaseBeginning(entityStruct.StructName) + ".go.tpl")
		if err != nil {
			panic(err)
		}

		tmpl, err := template.New("domain").Funcs(tools.FullFuncMap).Parse(string(tmplRaw))
		if err != nil {
			panic(err)
		}

		outFile, err := os.Create("model/domain/" + tools.LowercaseBeginning(entityStruct.StructName) + ".go")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(outFile, entityStruct)
		if err != nil {
			panic(err)
		}
	}
}
