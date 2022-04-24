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

	sqlite_repo "github.com/JonasMuehlmann/bntp.go/pkg/libbookmarks/repository/sqlite3"
)

type BookmarkManager interface {
	New(...any) (BookmarkManager, error)

	Add(context.Context, []sqlite_repo.Bookmark) (numAffectedRecords int, newID int, err error)
	Replace(context.Context, []sqlite_repo.Bookmark) error
	UpdateWhere(context.Context, BookmarkFilter, map[BookmarkField]BookmarkUpdateOperation) (numAffectedRecords int, err error)
	Delete(context.Context, []sqlite_repo.Bookmark) error
	DeleteWhere(context.Context, BookmarkFilter) (numAffectedRecords int, err error)
	CountWhere(context.Context, BookmarkFilter) int
	CountAll(context.Context) int
	DoesExist(context.Context, sqlite_repo.Bookmark) bool
	DoesExistWhere(context.Context, BookmarkFilter) bool
	GetWhere(context.Context, BookmarkFilter) []sqlite_repo.Bookmark
	GetFirstWhere(context.Context, BookmarkFilter) sqlite_repo.Bookmark
	GetAll(context.Context) []sqlite_repo.Bookmark

	AddType(context.Context, string) error
	DeleteType(context.Context, string) error
	UpdateType(context.Context, string, string) error
}

type BookmarkTagManager interface {
	New(...any) (BookmarkTagManager, error)

	AddTag(context.Context, string) error
	DeleteTag(context.Context, string) error
	UpdateTag(context.Context, string, string) error
	CountAllTags(context.Context) int
	DoesExistTag(context.Context, string) bool
	GetAllTags(context.Context, bool) ([]string, error)

	ApplyTagWhere(context.Context, string, BookmarkFilter) error
	RemoveTagWhere(context.Context, string, BookmarkFilter) error
}

type BookmarkField int

const (
	ID BookmarkField = iota
	IsRead
	Title
	URL
	BookmarkTypeID
	IsCollection
	CreatedAt
	UpdatedAt
	DeletedAt
)

type BookmarkUpdateOperation func(any) any

type BookmarkHook func(context.Context, sqlite_repo.Bookmark) error

type HookPoint int

// TODO: Generate from constants
type Hooks[TEntityHook any] struct {
	BeforeAddHook TEntityHook
	AfterAddHook  TEntityHook

	BeforeSelectHook TEntityHook
	AfterSelectHook  TEntityHook

	BeforeUpdateHook TEntityHook
	AfterUpdateHook  TEntityHook

	BeforeDeleteHook TEntityHook
	AfterDeleteHook  TEntityHook

	BeforeUpsertHook TEntityHook
	AfterUpsertHook  TEntityHook

	BeforeAnyHook TEntityHook
	AfterAnyHook  TEntityHook
}

// TODO: Do I need to do anything to support caching?

// TODO: Allow composing through e.g. BeforeAddHook | AfterAddHook
const (
	BeforeAddHook HookPoint = iota
	AfterAddHook

	BeforeSelectHook
	AfterSelectHook

	BeforeUpdateHook
	AfterUpdateHook

	BeforeDeleteHook
	AfterDeleteHook

	BeforeUpsertHook
	AfterUpsertHook

	BeforeAnyHook
	AfterAnyHook

	AfterError
	AfterDeadline
	AfterTimeout
	AfterCancel
)
