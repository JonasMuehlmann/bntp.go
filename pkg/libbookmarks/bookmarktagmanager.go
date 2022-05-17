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
	"context"

	"github.com/JonasMuehlmann/bntp.go/model/domain"
	bntp "github.com/JonasMuehlmann/bntp.go/pkg"
)

type BookmarkTagManager struct {
	hooks bntp.Hooks[string]
}

func (m *BookmarkTagManager) New(...any) (BookmarkTagManager, error) {
	panic("Not implemented")
}

func (m *BookmarkTagManager) AddTag(context.Context, string) error {
	panic("Not implemented")
}

func (m *BookmarkTagManager) DeleteTag(context.Context, string) error {
	panic("Not implemented")
}

func (m *BookmarkTagManager) UpdateTag(context.Context, string, string) error {
	panic("Not implemented")
}

func (m *BookmarkTagManager) CountAllTags(context.Context) int {
	panic("Not implemented")
}

func (m *BookmarkTagManager) DoesExistTag(context.Context, string) bool {
	panic("Not implemented")
}

func (m *BookmarkTagManager) GetAllTags(context.Context, bool) ([]string, error) {
	panic("Not implemented")
}

func (m *BookmarkTagManager) ApplyTagWhere(context.Context, string, domain.BookmarkFilter) error {
	panic("Not implemented")
}

func (m *BookmarkTagManager) RemoveTagWhere(context.Context, string, domain.BookmarkFilter) error {
	panic("Not implemented")
}
