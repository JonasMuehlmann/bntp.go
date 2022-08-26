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
	"errors"

	domain "github.com/JonasMuehlmann/bntp.go/model/domain"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/goaoi"
	log "github.com/sirupsen/logrus"

	bntp "github.com/JonasMuehlmann/bntp.go/bntp"
)

type DocumentManager struct {
	Repository repository.DocumentRepository
	Hooks      *bntp.Hooks[domain.Document]
	Logger     *log.Logger
}

func NewDocumentManager(logger *log.Logger, hooks *bntp.Hooks[domain.Document], repository repository.DocumentRepository) (DocumentManager, error) {
	m := DocumentManager{}
	m.Repository = repository
	m.Hooks = hooks
	m.Logger = logger

	return m, nil
}

// TODO: Allow skipping certain hooks.
func (m *DocumentManager) Add(ctx context.Context, documents []*domain.Document) error {
	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	err := m.Repository.Add(ctx, documents)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return err
}

func (m *DocumentManager) Replace(ctx context.Context, documents []*domain.Document) error {
	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	err := m.Repository.Replace(ctx, documents)
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

func (m *DocumentManager) Upsert(ctx context.Context, documents []*domain.Document) error {
	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	err := m.Repository.Upsert(ctx, documents)
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

func (m *DocumentManager) Update(ctx context.Context, documents []*domain.Document, documentUpdater *domain.DocumentUpdater) error {
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

func (m *DocumentManager) UpdateWhere(ctx context.Context, documentFilter *domain.DocumentFilter, documentUpdater *domain.DocumentUpdater) (numAffectedRecords int64, err error) {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	numAffectedRecords, err = m.Repository.UpdateWhere(ctx, documentFilter, documentUpdater)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return
}

func (m *DocumentManager) Delete(ctx context.Context, documents []*domain.Document) error {
	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	err := m.Repository.Delete(ctx, documents)
	if err != nil {
		m.Logger.Error(err)
	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return err
}

func (m *DocumentManager) DeleteWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (numAffectedRecords int64, err error) {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	numAffectedRecords, err = m.Repository.DeleteWhere(ctx, documentFilter)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return
}

func (m *DocumentManager) CountWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (numRecords int64, err error) {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	numRecords, err = m.Repository.CountWhere(ctx, documentFilter)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return
}

func (m *DocumentManager) CountAll(ctx context.Context) (numRecords int64, err error) {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	numRecords, err = m.Repository.CountAll(ctx)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return
}

func (m *DocumentManager) DoesExist(ctx context.Context, document *domain.Document) (doesExist bool, err error) {
	documents := []*domain.Document{document}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	doesExist, err = m.Repository.DoesExist(ctx, document)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return
}

func (m *DocumentManager) DoesExistWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (doesExist bool, err error) {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	doesExist, err = m.Repository.DoesExistWhere(ctx, documentFilter)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return
}

func (m *DocumentManager) GetWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (records []*domain.Document, err error) {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	records, err = m.Repository.GetWhere(ctx, documentFilter)
	if err != nil {
		m.Logger.Error(err)

	}
	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return
}

func (m *DocumentManager) GetFirstWhere(ctx context.Context, documentFilter *domain.DocumentFilter) (record *domain.Document, err error) {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	record, err = m.Repository.GetFirstWhere(ctx, documentFilter)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return
}

func (m *DocumentManager) GetAll(ctx context.Context) (records []*domain.Document, err error) {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	records, err = m.Repository.GetAll(ctx)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return
}

func (m *DocumentManager) AddType(ctx context.Context, types []string) error {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	err := m.Repository.AddType(ctx, types)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return err
}

func (m *DocumentManager) DeleteType(ctx context.Context, types []string) error {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	err := m.Repository.DeleteType(ctx, types)
	if err != nil {
		m.Logger.Error(err)

	}

	hookErr = goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	return err
}

func (m *DocumentManager) UpdateType(ctx context.Context, oldType string, newType string) error {
	documents := []*domain.Document{}

	hookErr := goaoi.ForeachSlice(documents, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if hookErr != nil && !errors.Is(hookErr, goaoi.EmptyIterableError{}) {
		hookErr = bntp.HookExecutionError{Inner: hookErr}
		m.Logger.Error(hookErr)

	}

	err := m.Repository.UpdateType(ctx, oldType, newType)
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
