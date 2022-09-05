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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/document.go.tpl


package domain

import (
	"time"

	"github.com/JonasMuehlmann/optional.go"
	"github.com/JonasMuehlmann/bntp.go/model"
)


type Document struct {
    
    CreatedAt time.Time `json:"created_at" toml:"created_at" yaml:"created_at"`
    UpdatedAt time.Time `json:"updated_at" toml:"updated_at" yaml:"updated_at"`
    DeletedAt optional.Optional[time.Time] `json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
    Path string `json:"path" toml:"path" yaml:"path"`
    DocumentType optional.Optional[string] `json:"document_type" toml:"document_type" yaml:"document_type"`
    TagIDs []int64 `json:"tagIDs" toml:"tagIDs" yaml:"tagIDs"`
    LinkedDocumentIDs []int64 `json:"linked_documentIDs" toml:"linked_documentIDs" yaml:"linked_documentIDs"`
    BacklinkedDocumentsIDs []int64 `json:"backlinked_documentIDs" toml:"backlinked_documentIDs" yaml:"backlinked_documentIDs"`
    ID int64 `json:"id" toml:"id" yaml:"id"`
}


func (t *Document) IsDefault() bool {
    
    var CreatedAtZero time.Time
    if t.CreatedAt != CreatedAtZero {
    
        return false
    }
    
    var UpdatedAtZero time.Time
    if t.UpdatedAt != UpdatedAtZero {
    
        return false
    }
    
    var DeletedAtZero optional.Optional[time.Time]
    if t.DeletedAt != DeletedAtZero {
    
        return false
    }
    
    var PathZero string
    if t.Path != PathZero {
    
        return false
    }
    
    var DocumentTypeZero optional.Optional[string]
    if t.DocumentType != DocumentTypeZero {
    
        return false
    }
    
    if t.TagIDs != nil {
    
        return false
    }
    
    if t.LinkedDocumentIDs != nil {
    
        return false
    }
    
    if t.BacklinkedDocumentsIDs != nil {
    
        return false
    }
    
    var IDZero int64
    if t.ID != IDZero {
    
        return false
    }
    

    return true
}

type DocumentField string

var DocumentFields = struct {
    CreatedAt  DocumentField
    UpdatedAt  DocumentField
    DeletedAt  DocumentField
    Path  DocumentField
    DocumentType  DocumentField
    TagIDs  DocumentField
    LinkedDocumentIDs  DocumentField
    BacklinkedDocumentsIDs  DocumentField
    ID  DocumentField
    
}{
    CreatedAt: "created_at",
    UpdatedAt: "updated_at",
    DeletedAt: "deleted_at",
    Path: "path",
    DocumentType: "document_type",
    TagIDs: "tagIDs",
    LinkedDocumentIDs: "linked_documentIDs",
    BacklinkedDocumentsIDs: "backlinked_documentIDs",
    ID: "id",
    
}

func (document *Document) GetCreatedAt() time.Time {
        return document.CreatedAt
}
func (document *Document) GetUpdatedAt() time.Time {
        return document.UpdatedAt
}
func (document *Document) GetDeletedAt() optional.Optional[time.Time] {
        return document.DeletedAt
}
func (document *Document) GetPath() string {
        return document.Path
}
func (document *Document) GetDocumentType() optional.Optional[string] {
        return document.DocumentType
}
func (document *Document) GetTagIDs() []int64 {
        return document.TagIDs
}
func (document *Document) GetLinkedDocumentIDs() []int64 {
        return document.LinkedDocumentIDs
}
func (document *Document) GetBacklinkedDocumentsIDs() []int64 {
        return document.BacklinkedDocumentsIDs
}
func (document *Document) GetID() int64 {
        return document.ID
}

func (document *Document) GetCreatedAtRef() *time.Time {
        return &document.CreatedAt
}
func (document *Document) GetUpdatedAtRef() *time.Time {
        return &document.UpdatedAt
}
func (document *Document) GetDeletedAtRef() *optional.Optional[time.Time] {
        return &document.DeletedAt
}
func (document *Document) GetPathRef() *string {
        return &document.Path
}
func (document *Document) GetDocumentTypeRef() *optional.Optional[string] {
        return &document.DocumentType
}
func (document *Document) GetTagIDsRef() *[]int64 {
        return &document.TagIDs
}
func (document *Document) GetLinkedDocumentIDsRef() *[]int64 {
        return &document.LinkedDocumentIDs
}
func (document *Document) GetBacklinkedDocumentsIDsRef() *[]int64 {
        return &document.BacklinkedDocumentsIDs
}
func (document *Document) GetIDRef() *int64 {
        return &document.ID
}


type DocumentFilter struct {
    CreatedAt optional.Optional[model.FilterOperation[time.Time]]`json:"createdAt,omitempty" toml:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    UpdatedAt optional.Optional[model.FilterOperation[time.Time]]`json:"updatedAt,omitempty" toml:"updatedAt,omitempty" yaml:"updatedAt,omitempty"`
    DeletedAt optional.Optional[model.FilterOperation[optional.Optional[time.Time]]]`json:"deletedAt,omitempty" toml:"deletedAt,omitempty" yaml:"deletedAt,omitempty"`
    Path optional.Optional[model.FilterOperation[string]]`json:"path,omitempty" toml:"path,omitempty" yaml:"path,omitempty"`
    DocumentType optional.Optional[model.FilterOperation[optional.Optional[string]]]`json:"documentType,omitempty" toml:"documentType,omitempty" yaml:"documentType,omitempty"`
    TagIDs optional.Optional[model.FilterOperation[int64]]`json:"tagIDs,omitempty" toml:"tagIDs,omitempty" yaml:"tagIDs,omitempty"`
    LinkedDocumentIDs optional.Optional[model.FilterOperation[int64]]`json:"linkedDocumentIDs,omitempty" toml:"linkedDocumentIDs,omitempty" yaml:"linkedDocumentIDs,omitempty"`
    BacklinkedDocumentsIDs optional.Optional[model.FilterOperation[int64]]`json:"backlinkedDocumentsIDs,omitempty" toml:"backlinkedDocumentsIDs,omitempty" yaml:"backlinkedDocumentsIDs,omitempty"`
    ID optional.Optional[model.FilterOperation[int64]]`json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty"`
    
}

func (filter *DocumentFilter) IsDefault() bool {
    if filter.CreatedAt.HasValue {
        return false
    }
    if filter.UpdatedAt.HasValue {
        return false
    }
    if filter.DeletedAt.HasValue {
        return false
    }
    if filter.Path.HasValue {
        return false
    }
    if filter.DocumentType.HasValue {
        return false
    }
    if filter.TagIDs.HasValue {
        return false
    }
    if filter.LinkedDocumentIDs.HasValue {
        return false
    }
    if filter.BacklinkedDocumentsIDs.HasValue {
        return false
    }
    if filter.ID.HasValue {
        return false
    }
    

    return true
}


type DocumentUpdater struct {
    CreatedAt optional.Optional[model.UpdateOperation[time.Time]]`json:"createdAt,omitempty" toml:"createdAt,omitempty" yaml:"createdAt,omitempty"`
    UpdatedAt optional.Optional[model.UpdateOperation[time.Time]]`json:"updatedAt,omitempty" toml:"updatedAt,omitempty" yaml:"updatedAt,omitempty"`
    DeletedAt optional.Optional[model.UpdateOperation[optional.Optional[time.Time]]]`json:"deletedAt,omitempty" toml:"deletedAt,omitempty" yaml:"deletedAt,omitempty"`
    Path optional.Optional[model.UpdateOperation[string]]`json:"path,omitempty" toml:"path,omitempty" yaml:"path,omitempty"`
    DocumentType optional.Optional[model.UpdateOperation[optional.Optional[string]]]`json:"documentType,omitempty" toml:"documentType,omitempty" yaml:"documentType,omitempty"`
    TagIDs optional.Optional[model.UpdateOperation[[]int64]]`json:"tagIDs,omitempty" toml:"tagIDs,omitempty" yaml:"tagIDs,omitempty"`
    LinkedDocumentIDs optional.Optional[model.UpdateOperation[[]int64]]`json:"linkedDocumentIDs,omitempty" toml:"linkedDocumentIDs,omitempty" yaml:"linkedDocumentIDs,omitempty"`
    BacklinkedDocumentsIDs optional.Optional[model.UpdateOperation[[]int64]]`json:"backlinkedDocumentsIDs,omitempty" toml:"backlinkedDocumentsIDs,omitempty" yaml:"backlinkedDocumentsIDs,omitempty"`
    ID optional.Optional[model.UpdateOperation[int64]]`json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty"`
    
}

func (updater *DocumentUpdater) IsDefault() bool {
    if updater.CreatedAt.HasValue {
        return false
    }
    if updater.UpdatedAt.HasValue {
        return false
    }
    if updater.DeletedAt.HasValue {
        return false
    }
    if updater.Path.HasValue {
        return false
    }
    if updater.DocumentType.HasValue {
        return false
    }
    if updater.TagIDs.HasValue {
        return false
    }
    if updater.LinkedDocumentIDs.HasValue {
        return false
    }
    if updater.BacklinkedDocumentsIDs.HasValue {
        return false
    }
    if updater.ID.HasValue {
        return false
    }
    

    return true
}

const (
    DocumentFilterUntagged = "DocumentFilterUntagged"
    DocumentFilterDeleted = "DocumentFilterDeleted"
)

// FIX: This operating on int64s instead of the slice is nonsense, right?
var PredefinedDocumentFilters = map[string]*DocumentFilter {
    DocumentFilterUntagged: {TagIDs: optional.Make(model.FilterOperation[int64]{
        Operand: model.ScalarOperand[int64]{
            Operand: -1,
        },
        Operator: model.FilterEqual,
    })},
    DocumentFilterDeleted: {DeletedAt: optional.Make(model.FilterOperation[optional.Optional[time.Time]]{
        Operand: model.ScalarOperand[optional.Optional[time.Time]]{
            Operand: optional.Optional[time.Time]{},
        },
        Operator: model.FilterEqual,
    })},
}
