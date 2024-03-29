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
	"errors"

	domain "github.com/JonasMuehlmann/bntp.go/model/domain"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/goaoi"

	bntp "github.com/JonasMuehlmann/bntp.go/bntp"
	log "github.com/sirupsen/logrus"
)

type BookmarkManager struct {
	Hooks      *bntp.Hooks[domain.Bookmark]
	Repository repository.BookmarkRepository
	Logger     *log.Logger
}

func NewBookmarkManager(logger *log.Logger, hooks *bntp.Hooks[domain.Bookmark], repository repository.BookmarkRepository) (BookmarkManager, error) {
	m := BookmarkManager{}
	m.Repository = repository
	m.Hooks = hooks
	m.Logger = logger

	return m, nil
}

// TODO: Allow skipping certain hooks.
func (m *BookmarkManager) Add(ctx context.Context, bookmarks []*domain.Bookmark) error {
	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	err := m.Repository.Add(ctx, bookmarks)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return err
}

func (m *BookmarkManager) Replace(ctx context.Context, bookmarks []*domain.Bookmark) error {
	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	err := m.Repository.Replace(ctx, bookmarks)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return err
}

func (m *BookmarkManager) Upsert(ctx context.Context, bookmarks []*domain.Bookmark) error {
	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	err := m.Repository.Upsert(ctx, bookmarks)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return err
}

func (m *BookmarkManager) Update(ctx context.Context, documents []*domain.Bookmark, documentUpdater *domain.BookmarkUpdater) error {
	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	err := m.Repository.Update(ctx, documents, documentUpdater)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return err
}

func (m *BookmarkManager) UpdateWhere(ctx context.Context, bookmarkFilter *domain.BookmarkFilter, bookmarkUpdater *domain.BookmarkUpdater) (numAffectedRecords int64, err error) {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	numAffectedRecords, err = m.Repository.UpdateWhere(ctx, bookmarkFilter, bookmarkUpdater)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}

func (m *BookmarkManager) Delete(ctx context.Context, bookmarks []*domain.Bookmark) error {
	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	err := m.Repository.Delete(ctx, bookmarks)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return err
}

func (m *BookmarkManager) DeleteWhere(ctx context.Context, bookmarkFilter *domain.BookmarkFilter) (numAffectedRecords int64, err error) {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	numAffectedRecords, err = m.Repository.DeleteWhere(ctx, bookmarkFilter)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}

func (m *BookmarkManager) CountWhere(ctx context.Context, bookmarkFilter *domain.BookmarkFilter) (numRecords int64, err error) {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	numRecords, err = m.Repository.CountWhere(ctx, bookmarkFilter)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}

func (m *BookmarkManager) CountAll(ctx context.Context) (numRecords int64, err error) {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	numRecords, err = m.Repository.CountAll(ctx)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}

func (m *BookmarkManager) DoesExist(ctx context.Context, bookmark *domain.Bookmark) (doesExist bool, err error) {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	doesExist, err = m.Repository.DoesExist(ctx, bookmark)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}

func (m *BookmarkManager) DoesExistWhere(ctx context.Context, bookmarkFilter *domain.BookmarkFilter) (doesExist bool, err error) {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	doesExist, err = m.Repository.DoesExistWhere(ctx, bookmarkFilter)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}

func (m *BookmarkManager) GetWhere(ctx context.Context, bookmarkFilter *domain.BookmarkFilter) (records []*domain.Bookmark, err error) {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	records, err = m.Repository.GetWhere(ctx, bookmarkFilter)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}

func (m *BookmarkManager) GetFirstWhere(ctx context.Context, bookmarkFilter *domain.BookmarkFilter) (record *domain.Bookmark, err error) {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	record, err = m.Repository.GetFirstWhere(ctx, bookmarkFilter)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}

func (m *BookmarkManager) GetAll(ctx context.Context) (records []*domain.Bookmark, err error) {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	records, err = m.Repository.GetAll(ctx)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}

func (m *BookmarkManager) AddType(ctx context.Context, types []string) error {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	err := m.Repository.AddType(ctx, types)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return err
}

func (m *BookmarkManager) DeleteType(ctx context.Context, types []string) error {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	err := m.Repository.DeleteType(ctx, types)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return err
}

func (m *BookmarkManager) UpdateType(ctx context.Context, oldType string, newType string) error {
	bookmarks := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	err := m.Repository.UpdateType(ctx, oldType, newType)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(bookmarks, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return err
}

func (m *BookmarkManager) GetAllTypes(ctx context.Context) ([]string, error) {
	documents := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	types, err := m.Repository.GetAllTypes(ctx)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return types, err
}

func (m *BookmarkManager) GetFromIDs(ctx context.Context, ids []int64) (records []*domain.Bookmark, err error) {
	tags := []*domain.Bookmark{}

	hookErr := goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	records, err = m.Repository.GetFromIDs(ctx, ids)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)
	}

	return
}
