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
	log "github.com/sirupsen/logrus"
)

// TODO: Allow skipping certain hooks.
type BookmarkManager struct {
	hooks      bntp.Hooks[domain.Bookmark]
	repository repository.BookmarkRepository
}

func (m *BookmarkManager) New(...any) (BookmarkManager, error) {
	panic("Not implemented")
}

// TODO: Execute hooks
func (m *BookmarkManager) Add(ctx context.Context, bookmarks []*domain.Bookmark) error {
	err := m.repository.Add(ctx, bookmarks)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *BookmarkManager) Replace(ctx context.Context, bookmarks []*domain.Bookmark) error {
	err := m.repository.Replace(ctx, bookmarks)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *BookmarkManager) UpdateWhere(ctx context.Context, bookmarkFilter domain.BookmarkFilter, bookmarkUpdater domain.BookmarkUpdater) (numAffectedRecords int64, err error) {
	numAffectedRecords, err = m.repository.UpdateWhere(ctx, bookmarkFilter, bookmarkUpdater)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *BookmarkManager) Delete(ctx context.Context, bookmarks []*domain.Bookmark) error {
	err := m.repository.Delete(ctx, bookmarks)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *BookmarkManager) DeleteWhere(ctx context.Context, bookmarkFilter domain.BookmarkFilter) (numAffectedRecords int64, err error) {
	numAffectedRecords, err = m.repository.DeleteWhere(ctx, bookmarkFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *BookmarkManager) CountWhere(ctx context.Context, bookmarkFilter domain.BookmarkFilter) (numRecords int64, err error) {
	numRecords, err = m.repository.CountWhere(ctx, bookmarkFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *BookmarkManager) CountAll(ctx context.Context) (numRecords int64, err error) {
	numRecords, err = m.repository.CountAll(ctx)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *BookmarkManager) DoesExist(ctx context.Context, bookmark *domain.Bookmark) (doesExist bool, err error) {
	doesExist, err = m.repository.DoesExist(ctx, bookmark)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *BookmarkManager) DoesExistWhere(ctx context.Context, bookmarkFilter domain.BookmarkFilter) (doesExist bool, err error) {
	doesExist, err = m.repository.DoesExistWhere(ctx, bookmarkFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *BookmarkManager) GetWhere(ctx context.Context, bookmarkFilter domain.BookmarkFilter) (records []*domain.Bookmark, err error) {
	records, err = m.repository.GetWhere(ctx, bookmarkFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *BookmarkManager) GetFirstWhere(ctx context.Context, bookmarkFilter domain.BookmarkFilter) (record *domain.Bookmark, err error) {
	record, err = m.repository.GetFirstWhere(ctx, bookmarkFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *BookmarkManager) GetAll(ctx context.Context) (records []*domain.Bookmark, err error) {
	records, err = m.repository.GetAll(ctx)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *BookmarkManager) AddType(ctx context.Context, type_ string) error {
	err := m.repository.AddType(ctx, type_)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *BookmarkManager) DeleteType(ctx context.Context, type_ string) error {
	err := m.repository.DeleteType(ctx, type_)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *BookmarkManager) UpdateType(ctx context.Context, oldType string, newType string) error {
	err := m.repository.UpdateType(ctx, oldType, newType)
	if err != nil {
		log.Error(err)
	}

	return err
}
