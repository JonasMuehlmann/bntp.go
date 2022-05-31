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

package libdocuments

import (
	"context"

	domain "github.com/JonasMuehlmann/bntp.go/model/domain"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository"
	log "github.com/sirupsen/logrus"

	bntp "github.com/JonasMuehlmann/bntp.go/pkg"
)

type DocumentManager struct {
	repository repository.DocumentRepository
	hooks      bntp.Hooks[domain.Document]
}

func (m *DocumentManager) New(...any) (DocumentManager, error) {
	panic("Not implemented")
}

// TODO: Allow skipping certain hooks.
// TODO: Execute hooks
func (m *DocumentManager) Add(ctx context.Context, documents []*domain.Document) error {
	err := m.repository.Add(ctx, documents)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *DocumentManager) Replace(ctx context.Context, documents []*domain.Document) error {
	err := m.repository.Replace(ctx, documents)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *DocumentManager) UpdateWhere(ctx context.Context, documentFilter domain.DocumentFilter, documentUpdater domain.DocumentUpdater) (numAffectedRecords int64, err error) {
	numAffectedRecords, err = m.repository.UpdateWhere(ctx, documentFilter, documentUpdater)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *DocumentManager) Delete(ctx context.Context, documents []*domain.Document) error {
	err := m.repository.Delete(ctx, documents)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *DocumentManager) DeleteWhere(ctx context.Context, documentFilter domain.DocumentFilter) (numAffectedRecords int64, err error) {
	numAffectedRecords, err = m.repository.DeleteWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *DocumentManager) CountWhere(ctx context.Context, documentFilter domain.DocumentFilter) (numRecords int64, err error) {
	numRecords, err = m.repository.CountWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *DocumentManager) CountAll(ctx context.Context) (numRecords int64, err error) {
	numRecords, err = m.repository.CountAll(ctx)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *DocumentManager) DoesExist(ctx context.Context, document *domain.Document) (doesExist bool, err error) {
	doesExist, err = m.repository.DoesExist(ctx, document)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *DocumentManager) DoesExistWhere(ctx context.Context, documentFilter domain.DocumentFilter) (doesExist bool, err error) {
	doesExist, err = m.repository.DoesExistWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *DocumentManager) GetWhere(ctx context.Context, documentFilter domain.DocumentFilter) (records []*domain.Document, err error) {
	records, err = m.repository.GetWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *DocumentManager) GetFirstWhere(ctx context.Context, documentFilter domain.DocumentFilter) (record *domain.Document, err error) {
	record, err = m.repository.GetFirstWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *DocumentManager) GetAll(ctx context.Context) (records []*domain.Document, err error) {
	records, err = m.repository.GetAll(ctx)
	if err != nil {
		log.Error(err)
	}

	return
}

func (m *DocumentManager) AddType(ctx context.Context, type_ string) error {
	err := m.repository.AddType(ctx, type_)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *DocumentManager) DeleteType(ctx context.Context, type_ string) error {
	err := m.repository.DeleteType(ctx, type_)
	if err != nil {
		log.Error(err)
	}

	return err
}

func (m *DocumentManager) UpdateType(ctx context.Context, oldType string, newType string) error {
	err := m.repository.UpdateType(ctx, oldType, newType)
	if err != nil {
		log.Error(err)
	}

	return err
}
