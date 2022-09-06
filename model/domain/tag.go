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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/tag.go.tpl

package domain

import (
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/optional.go"
)

type Tag struct {
	Tag           string  `json:"tag" toml:"tag" yaml:"tag"`
	ParentPathIDs []int64 `json:"parentPathIDs" toml:"parentPathIDs" yaml:"parentPathIDs"`
	SubtagIDs     []int64 `json:"subtagsIDs" toml:"subtagsIDs" yaml:"subtagsIDs"`
	ID            int64   `json:"id" toml:"id" yaml:"id"`
}

func (t *Tag) IsDefault() bool {

	var IDZero int64
	if t.ID != IDZero {

		return false
	}

	if t.ParentPathIDs != nil {

		return false
	}

	var TagZero string
	if t.Tag != TagZero {

		return false
	}

	if t.SubtagIDs != nil {

		return false
	}

	return true
}

type TagField string

var TagFields = struct {
	ID            TagField
	ParentPathIDs TagField
	Tag           TagField
	SubtagIDs     TagField
}{
	ID:            "id",
	ParentPathIDs: "parentPathIDs",
	Tag:           "tag",
	SubtagIDs:     "subtagsIDs",
}

func (tag *Tag) GetID() int64 {
	return tag.ID
}
func (tag *Tag) GetParentPathIDs() []int64 {
	return tag.ParentPathIDs
}
func (tag *Tag) GetTag() string {
	return tag.Tag
}
func (tag *Tag) GetSubtagIDs() []int64 {
	return tag.SubtagIDs
}

func (tag *Tag) GetIDRef() *int64 {
	return &tag.ID
}
func (tag *Tag) GetParentPathIDsRef() *[]int64 {
	return &tag.ParentPathIDs
}
func (tag *Tag) GetTagRef() *string {
	return &tag.Tag
}
func (tag *Tag) GetSubtagIDsRef() *[]int64 {
	return &tag.SubtagIDs
}

type TagFilter struct {
	ID            optional.Optional[model.FilterOperation[int64]]  `json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty"`
	ParentPathIDs optional.Optional[model.FilterOperation[int64]]  `json:"parentPathIDs,omitempty" toml:"parentPathIDs,omitempty" yaml:"parentPathIDs,omitempty"`
	Tag           optional.Optional[model.FilterOperation[string]] `json:"tag,omitempty" toml:"tag,omitempty" yaml:"tag,omitempty"`
	SubtagIDs     optional.Optional[model.FilterOperation[int64]]  `json:"subtagIDs,omitempty" toml:"subtagIDs,omitempty" yaml:"subtagIDs,omitempty"`
}

func (filter *TagFilter) IsDefault() bool {
	if filter.ID.HasValue {
		return false
	}
	if filter.ParentPathIDs.HasValue {
		return false
	}
	if filter.Tag.HasValue {
		return false
	}
	if filter.SubtagIDs.HasValue {
		return false
	}

	return true
}

type TagUpdater struct {
	Tag           optional.Optional[model.UpdateOperation[string]]  `json:"tag,omitempty" toml:"tag,omitempty" yaml:"tag,omitempty"`
	ParentPathIDs optional.Optional[model.UpdateOperation[[]int64]] `json:"parentPathIDs,omitempty" toml:"parentPathIDs,omitempty" yaml:"parentPathIDs,omitempty"`
	SubtagIDs     optional.Optional[model.UpdateOperation[[]int64]] `json:"subtagIDs,omitempty" toml:"subtagIDs,omitempty" yaml:"subtagIDs,omitempty"`
	ID            optional.Optional[model.UpdateOperation[int64]]   `json:"id,omitempty" toml:"id,omitempty" yaml:"id,omitempty"`
}

func (updater *TagUpdater) IsDefault() bool {
	if updater.ID.HasValue {
		return false
	}
	if updater.ParentPathIDs.HasValue {
		return false
	}
	if updater.Tag.HasValue {
		return false
	}
	if updater.SubtagIDs.HasValue {
		return false
	}

	return true
}

const (
	TagFilterLeaf = "TagFilterLeaf"
	TagFilterRoot = "TagFilterRoot"
)

var PredefinedTagFilters = map[string]*TagFilter{
	// FIX: This operating on int64s instead of the slice is nonsense, right?
	TagFilterLeaf: {ParentPathIDs: optional.Make(model.FilterOperation[int64]{
		Operand: model.ScalarOperand[int64]{
			Operand: -1,
		},
		Operator: model.FilterEqual,
	})},
	// FIX: This operating on int64s instead of the slice is nonsense, right?
	TagFilterRoot: {SubtagIDs: optional.Make(model.FilterOperation[int64]{
		Operand: model.ScalarOperand[int64]{
			Operand: -1,
		},
		Operator: model.FilterEqual,
	})},
}
