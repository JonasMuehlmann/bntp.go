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
	Tag        string `json:"tag" toml:"tag" yaml:"tag"`
	ParentPath []*Tag `json:"parentPath" toml:"parentPath" yaml:"parentPath"`
	Subtags    []*Tag `json:"subtags" toml:"subtags" yaml:"subtags"`
	ID         int64  `json:"id" toml:"id" yaml:"id"`
}

type TagField string

var TagFields = struct {
	ID         TagField
	ParentPath TagField
	Tag        TagField
	Subtags    TagField
}{
	ID:         "id",
	ParentPath: "parentPath",
	Tag:        "tag",
	Subtags:    "subtags",
}

type TagFilter struct {
	ID         optional.Optional[model.FilterOperation[int64]]
	ParentPath optional.Optional[model.FilterOperation[*Tag]]
	Tag        optional.Optional[model.FilterOperation[string]]
	Subtags    optional.Optional[model.FilterOperation[*Tag]]
}

type TagUpdater struct {
	Tag        optional.Optional[model.UpdateOperation[string]]
	ParentPath optional.Optional[model.UpdateOperation[[]*Tag]]
	Subtags    optional.Optional[model.UpdateOperation[[]*Tag]]
	ID         optional.Optional[model.UpdateOperation[int64]]
}

const (
	TagFilterLeaf = "TagFilterLeaf"
	TagFilterRoot = "TagFilterRoot"
)

var PredefinedTagFilters = map[string]TagFilter{
	TagFilterLeaf: {ParentPath: optional.Make(model.FilterOperation[*Tag]{
		Operand: model.ScalarOperand[*Tag]{
			Operand: nil,
		},
		Operator: model.FilterEqual,
	})},
	TagFilterRoot: {Subtags: optional.Make(model.FilterOperation[*Tag]{
		Operand: model.ScalarOperand[*Tag]{
			Operand: nil,
		},
		Operator: model.FilterEqual,
	})},
}
