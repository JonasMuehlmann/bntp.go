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

	repository "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/goaoi"

	bntp "github.com/JonasMuehlmann/bntp.go/bntp"
	"github.com/barweiss/go-tuple"
	log "github.com/sirupsen/logrus"
)

type DocumentContentManager struct {
	Repository repository.DocumentContentRepository
	Hooks      *bntp.Hooks[string]
}

func NewDocumentContentRepository(hooks *bntp.Hooks[string], repository repository.DocumentContentRepository) (DocumentContentManager, error) {
	m := DocumentContentManager{}
	m.Repository = repository
	m.Hooks = hooks

	return m, nil
}

// TODO: Allow skipping certain hooks.
// TODO: Implement context handling.
func (m *DocumentContentManager) Add(ctx context.Context, pathContents []tuple.T2[string, string]) error {
	paths := bntp.TupleToSOA2(pathContents).V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Add(ctx, pathContents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) Update(ctx context.Context, pathContents []tuple.T2[string, string]) error {
	paths := bntp.TupleToSOA2(pathContents).V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Update(ctx, pathContents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) Move(ctx context.Context, pathChanges []tuple.T2[string, string]) error {
	paths := bntp.TupleToSOA2(pathChanges).V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Move(ctx, pathChanges)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterUpdateHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) Delete(ctx context.Context, paths []string) error {
	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	err = m.Repository.Delete(ctx, paths)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) Get(ctx context.Context, paths []string) (contents []string, err error) {
	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	contents, err = m.Repository.Get(ctx, paths)
	if err != nil {
		log.Error(err)

		return
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterSelectHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return
	}

	return
}

func (m *DocumentContentManager) AddTags(ctx context.Context, pathTags []tuple.T2[string, []string]) error {
	soa := bntp.TupleToSOA2(pathTags)
	paths := soa.V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	contents, err := m.Repository.Get(ctx, paths)
	if err != nil {
		log.Error(err)

		return err
	}

	newContents := make([]string, len(contents))

	contentTags := tuple.T2[[]string, [][]string]{V1: contents, V2: soa.V2}
	for i := range contentTags.V1 {
		newContent, err := AddTags(ctx, contentTags.V1[i], contentTags.V2[i])
		if err != nil {
			log.Error(err)

			return err
		}

		newContents = append(newContents, newContent)
	}

	newPathContents := bntp.TupleToAOS2(tuple.T2[[]string, []string]{V1: paths, V2: newContents})

	err = m.Repository.Update(ctx, newPathContents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) RemoveTags(ctx context.Context, pathTags []tuple.T2[string, []string]) error {
	soa := bntp.TupleToSOA2(pathTags)
	paths := soa.V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	contents, err := m.Repository.Get(ctx, paths)
	if err != nil {
		log.Error(err)

		return err
	}

	newContents := make([]string, len(contents))

	contentTags := tuple.T2[[]string, [][]string]{V1: contents, V2: soa.V2}
	for i := range contentTags.V1 {
		newContent, err := RemoveTags(ctx, contentTags.V1[i], contentTags.V2[i])
		if err != nil {
			log.Error(err)

			return err
		}

		newContents = append(newContents, newContent)
	}

	newPathContents := bntp.TupleToAOS2(tuple.T2[[]string, []string]{V1: paths, V2: newContents})

	err = m.Repository.Update(ctx, newPathContents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) AddLinks(ctx context.Context, pathLinks []tuple.T2[string, []string]) error {
	soa := bntp.TupleToSOA2(pathLinks)
	paths := soa.V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	contents, err := m.Repository.Get(ctx, paths)
	if err != nil {
		log.Error(err)

		return err
	}

	newContents := make([]string, len(contents))

	contentTags := tuple.T2[[]string, [][]string]{V1: contents, V2: soa.V2}
	for i := range contentTags.V1 {
		newContent, err := AddLinks(ctx, contentTags.V1[i], contentTags.V2[i])
		if err != nil {
			log.Error(err)

			return err
		}

		newContents = append(newContents, newContent)
	}

	newPathContents := bntp.TupleToAOS2(tuple.T2[[]string, []string]{V1: paths, V2: newContents})

	err = m.Repository.Update(ctx, newPathContents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) RemoveLinks(ctx context.Context, pathLinks []tuple.T2[string, []string]) error {
	soa := bntp.TupleToSOA2(pathLinks)
	paths := soa.V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	contents, err := m.Repository.Get(ctx, paths)
	if err != nil {
		log.Error(err)

		return err
	}

	newContents := make([]string, len(contents))

	contentTags := tuple.T2[[]string, [][]string]{V1: contents, V2: soa.V2}
	for i := range contentTags.V1 {
		newContent, err := RemoveLinks(ctx, contentTags.V1[i], contentTags.V2[i])
		if err != nil {
			log.Error(err)

			return err
		}

		newContents = append(newContents, newContent)
	}

	newPathContents := bntp.TupleToAOS2(tuple.T2[[]string, []string]{V1: paths, V2: newContents})

	err = m.Repository.Update(ctx, newPathContents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) AddBackLinks(ctx context.Context, pathBacklinks []tuple.T2[string, []string]) error {
	soa := bntp.TupleToSOA2(pathBacklinks)
	paths := soa.V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	contents, err := m.Repository.Get(ctx, paths)
	if err != nil {
		log.Error(err)

		return err
	}

	newContents := make([]string, len(contents))

	contentTags := tuple.T2[[]string, [][]string]{V1: contents, V2: soa.V2}
	for i := range contentTags.V1 {
		newContent, err := AddBacklinks(ctx, contentTags.V1[i], contentTags.V2[i])
		if err != nil {
			log.Error(err)

			return err
		}

		newContents = append(newContents, newContent)
	}

	newPathContents := bntp.TupleToAOS2(tuple.T2[[]string, []string]{V1: paths, V2: newContents})

	err = m.Repository.Update(ctx, newPathContents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterAddHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) RemoveBackLinks(ctx context.Context, pathBacklinks []tuple.T2[string, []string]) error {
	soa := bntp.TupleToSOA2(pathBacklinks)
	paths := soa.V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	contents, err := m.Repository.Get(ctx, paths)
	if err != nil {
		log.Error(err)

		return err
	}

	newContents := make([]string, len(contents))

	contentTags := tuple.T2[[]string, [][]string]{V1: contents, V2: soa.V2}
	for i := range contentTags.V1 {
		newContent, err := RemoveBacklinks(ctx, contentTags.V1[i], contentTags.V2[i])
		if err != nil {
			log.Error(err)

			return err
		}

		newContents = append(newContents, newContent)
	}

	newPathContents := bntp.TupleToAOS2(tuple.T2[[]string, []string]{V1: paths, V2: newContents})

	err = m.Repository.Update(ctx, newPathContents)
	if err != nil {
		log.Error(err)

		return err
	}

	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.AfterAnyHook|bntp.AfterDeleteHook))
	if err != nil {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}
