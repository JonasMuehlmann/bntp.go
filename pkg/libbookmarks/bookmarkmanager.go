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

	domain "github.com/JonasMuehlmann/bntp.go/model/domain"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository"

	bntp "github.com/JonasMuehlmann/bntp.go/pkg"
)

// TODO: Allow skipping certain hooks
type BookmarkManager struct {
	hooks      bntp.Hooks[Bookmark]
	repository repository.BookmarkRepository
}

func (m *BookmarkManager) New(...any) (BookmarkManager, error) {
}

func (m *BookmarkManager) Add(context.Context, []domain.Bookmark) (numAffectedRecords int, newID int, err error) {
}
func (m *BookmarkManager) Replace(context.Context, []domain.Bookmark) error {
}
func (m *BookmarkManager) UpdateWhere(context.Context, BookmarkFilter, map[domain.BookmarkField]domain.BookmarkUpdater) (numAffectedRecords int, err error) {
}
func (m *BookmarkManager) Delete(context.Context, []domain.Bookmark) error {
}
func (m *BookmarkManager) DeleteWhere(context.Context, BookmarkFilter) (numAffectedRecords int, err error) {
}
func (m *BookmarkManager) CountWhere(context.Context, BookmarkFilter) int {
}
func (m *BookmarkManager) CountAll(context.Context) int {
}
func (m *BookmarkManager) DoesExist(context.Context, domain.Bookmark) bool {
}
func (m *BookmarkManager) DoesExistWhere(context.Context, BookmarkFilter) bool {
}
func (m *BookmarkManager) GetWhere(context.Context, BookmarkFilter) []domain.Bookmark {
}
func (m *BookmarkManager) GetFirstWhere(context.Context, BookmarkFilter) domain.Bookmark {
}
func (m *BookmarkManager) GetAll(context.Context) []domain.Bookmark {
}

func (m *BookmarkManager) AddType(context.Context, string) error {
}
func (m *BookmarkManager) DeleteType(context.Context, string) error {
}
func (m *BookmarkManager) UpdateType(context.Context, string, string) error {
}
