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

package libtags

import (
	"context"

	domain "github.com/JonasMuehlmann/bntp.go/model/domain"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository/sql"
	"github.com/JonasMuehlmann/goaoi"
	log "github.com/sirupsen/logrus"

	bntp "github.com/JonasMuehlmann/bntp.go/bntp"
)

type TagManager struct {
	Repository repository.TagRepository
	Hooks      *bntp.Hooks[domain.Tag]
}

func New(hooks *bntp.Hooks[domain.Tag], repository repository.TagRepository) (TagManager, error) {
	m := TagManager{}
	m.Repository = repository
	m.Hooks = hooks

	return m, nil
}

// TODO: Allow skipping certain hooks.
func (m *TagManager) Add(ctx context.Context, tags []*domain.Tag) error {
	err := goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Add(ctx, tags)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *TagManager) Replace(ctx context.Context, tags []*domain.Tag) error {
	err := goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Replace(ctx, tags)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *TagManager) UpdateWhere(ctx context.Context, tagFilter *domain.TagFilter, tagUpdater *domain.TagUpdater) (numAffectedRecords int64, err error) {
	tags := []*domain.Tag{}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	numAffectedRecords, err = m.Repository.UpdateWhere(ctx, tagFilter, tagUpdater)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *TagManager) Delete(ctx context.Context, tags []*domain.Tag) error {
	err := goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Delete(ctx, tags)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *TagManager) DeleteWhere(ctx context.Context, tagFilter *domain.TagFilter) (numAffectedRecords int64, err error) {
	tags := []*domain.Tag{}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	numAffectedRecords, err = m.Repository.DeleteWhere(ctx, tagFilter)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *TagManager) CountWhere(ctx context.Context, tagFilter *domain.TagFilter) (numRecords int64, err error) {
	tags := []*domain.Tag{}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	numRecords, err = m.Repository.CountWhere(ctx, tagFilter)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *TagManager) CountAll(ctx context.Context) (numRecords int64, err error) {
	tags := []*domain.Tag{}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	numRecords, err = m.Repository.CountAll(ctx)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *TagManager) DoesExist(ctx context.Context, tag *domain.Tag) (doesExist bool, err error) {
	tags := []*domain.Tag{tag}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	doesExist, err = m.Repository.DoesExist(ctx, tag)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *TagManager) DoesExistWhere(ctx context.Context, tagFilter *domain.TagFilter) (doesExist bool, err error) {
	tags := []*domain.Tag{}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	doesExist, err = m.Repository.DoesExistWhere(ctx, tagFilter)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *TagManager) GetWhere(ctx context.Context, tagFilter *domain.TagFilter) (records []*domain.Tag, err error) {
	tags := []*domain.Tag{}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	records, err = m.Repository.GetWhere(ctx, tagFilter)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *TagManager) GetFirstWhere(ctx context.Context, tagFilter *domain.TagFilter) (record *domain.Tag, err error) {
	tags := []*domain.Tag{}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	record, err = m.Repository.GetFirstWhere(ctx, tagFilter)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *TagManager) GetAll(ctx context.Context) (records []*domain.Tag, err error) {
	tags := []*domain.Tag{}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	records, err = m.Repository.GetAll(ctx)
	if err != nil {
		log.Error(err)
	}

	err = goaoi.ForeachSlice(tags, m.Hooks.PartiallySpecializeExecuteHooks(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}
