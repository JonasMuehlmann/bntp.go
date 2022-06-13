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
	"github.com/JonasMuehlmann/goaoi"
	log "github.com/sirupsen/logrus"

	bntp "github.com/JonasMuehlmann/bntp.go/bntp"
)

type DocumentManager struct {
	Repository repository.DocumentRepository
	Hooks      *bntp.Hooks[domain.Document]
}

func NewDocumentManager(hooks *bntp.Hooks[domain.Document], repository repository.DocumentRepository) (DocumentManager, error) {
	m := DocumentManager{}
	m.Repository = repository
	m.Hooks = hooks

	return m, nil
}

// TODO: Allow skipping certain hooks.
func (m *DocumentManager) Add(ctx context.Context, documents []*domain.Document) error {
	err := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Add(ctx, documents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentManager) Replace(ctx context.Context, documents []*domain.Document) error {
	err := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Replace(ctx, documents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentManager) UpdateWhere(ctx context.Context, documentFilter *domain.DocumentFilter, documentUpdater *domain.DocumentUpdater) (numAffectedRecords int64, err error) {
	documents := []*domain.Document{}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	numAffectedRecords, err = m.Repository.UpdateWhere(ctx, documentFilter, documentUpdater)
	if err != nil {
		log.Error(err)

		return
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentManager) Delete(ctx context.Context, documents []*domain.Document) error {
	err := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Delete(ctx, documents)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentManager) DeleteWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (numAffectedRecords int64, err error) {
	documents := []*domain.Document{}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	numAffectedRecords, err = m.Repository.DeleteWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)

		return
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentManager) CountWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (numRecords int64, err error) {
	documents := []*domain.Document{}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	numRecords, err = m.Repository.CountWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)

		return
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentManager) CountAll(ctx context.Context) (numRecords int64, err error) {
	documents := []*domain.Document{}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	numRecords, err = m.Repository.CountAll(ctx)
	if err != nil {
		log.Error(err)

		return
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentManager) DoesExist(ctx context.Context, document *domain.Document) (doesExist bool, err error) {
	documents := []*domain.Document{document}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	doesExist, err = m.Repository.DoesExist(ctx, document)
	if err != nil {
		log.Error(err)

		return
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentManager) DoesExistWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (doesExist bool, err error) {
	documents := []*domain.Document{}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	doesExist, err = m.Repository.DoesExistWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)

		return
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentManager) GetWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (records []*domain.Document, err error) {
	documents := []*domain.Document{}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	records, err = m.Repository.GetWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)

		return
	}
	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentManager) GetFirstWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (record *domain.Document, err error) {
	documents := []*domain.Document{}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	record, err = m.Repository.GetFirstWhere(ctx, documentFilter)
	if err != nil {
		log.Error(err)

		return
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentManager) GetAll(ctx context.Context) (records []*domain.Document, err error) {
	documents := []*domain.Document{}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	records, err = m.Repository.GetAll(ctx)
	if err != nil {
		log.Error(err)

		return
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentManager) AddType(ctx context.Context, types []string) error {
	documents := []*domain.Document{}

	err := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.AddType(ctx, types)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentManager) DeleteType(ctx context.Context, types []string) error {
	documents := []*domain.Document{}

	err := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.DeleteType(ctx, types)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentManager) UpdateType(ctx context.Context, oldType string, newType string) error {
	documents := []*domain.Document{}

	err := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.UpdateType(ctx, oldType, newType)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}
