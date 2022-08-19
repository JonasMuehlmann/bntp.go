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

	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/optional.go"

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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) Update(ctx context.Context, pathContents []tuple.T2[string, string]) error {
	paths := bntp.TupleToSOA2(pathContents).V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) Move(ctx context.Context, pathChanges []tuple.T2[string, string]) error {
	paths := bntp.TupleToSOA2(pathChanges).V1

	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeUpdateHook))
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) Delete(ctx context.Context, paths []string) error {
	err := goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeDeleteHook))
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) Get(ctx context.Context, paths []string) (contents []string, err error) {
	err = goaoi.ForeachSlice(paths, m.Hooks.PartiallySpecializeExecuteHooksForNoPointer(ctx, bntp.BeforeAnyHook|bntp.BeforeSelectHook))
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
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
	if err != nil && !errors.As(err, &goaoi.EmptyIterableError{}) {
		err = bntp.HookExecutionError{Inner: err}
		log.Error(err)

		return err
	}

	return err
}

func (m *DocumentContentManager) UpdateDocumentContentsFromNewModels(ctx context.Context, newDocuments []*domain.Document, documentManager *DocumentManager) error {
	addedPathLinks := make([]tuple.T2[string, []string], 0, len(newDocuments))
	removedPathLinks := make([]tuple.T2[string, []string], 0, len(newDocuments))
	addedPathBacklinks := make([]tuple.T2[string, []string], 0, len(newDocuments))
	removedPathBacklinks := make([]tuple.T2[string, []string], 0, len(newDocuments))
	addedPathTags := make([]tuple.T2[string, []string], 0, len(newDocuments))
	removedPathTags := make([]tuple.T2[string, []string], 0, len(newDocuments))

	for i, newDocument := range newDocuments {
		filter := &domain.DocumentFilter{ID: optional.Make(model.FilterOperation[int64]{Operand: model.ScalarOperand[int64]{newDocument.ID}, Operator: model.FilterEqual})}

		oldDocument, err := documentManager.GetFirstWhere(ctx, filter)
		if err != nil {
			return err
		}

		addedLinkIDs, err := GetAddedLinks(oldDocument, newDocument)
		if err != nil {
			return err
		}

		addedLinkDocuments, err := documentManager.GetWhere(ctx, &domain.DocumentFilter{
			LinkedDocumentIDs: optional.Make(model.FilterOperation[int64]{
				Operator: model.FilterIn,
				Operand:  model.ListOperand[int64]{addedLinkIDs}}),
		})
		if err != nil {
			return err
		}

		addedLinks, err := goaoi.TransformCopySliceUnsafe(addedLinkDocuments, (*domain.Document).GetPath)
		if err != nil {
			return err
		}

		addedPathLinks[i] = tuple.T2[string, []string]{oldDocument.Path, addedLinks}

		removedLinkIDs, err := GetRemovedLinks(oldDocument, newDocument)
		if err != nil {
			return err
		}

		removedLinkDocuments, err := documentManager.GetWhere(ctx, &domain.DocumentFilter{
			LinkedDocumentIDs: optional.Make(model.FilterOperation[int64]{
				Operator: model.FilterIn,
				Operand:  model.ListOperand[int64]{removedLinkIDs}}),
		})
		if err != nil {
			return err
		}

		removedLinks, err := goaoi.TransformCopySliceUnsafe(removedLinkDocuments, (*domain.Document).GetPath)
		if err != nil {
			return err
		}

		removedPathLinks[i] = tuple.T2[string, []string]{oldDocument.Path, removedLinks}

		addedBacklinkIDs, err := GetAddedBacklinks(oldDocument, newDocument)
		if err != nil {
			return err
		}

		addedBacklinkDocuments, err := documentManager.GetWhere(ctx, &domain.DocumentFilter{
			LinkedDocumentIDs: optional.Make(model.FilterOperation[int64]{
				Operator: model.FilterIn,
				Operand:  model.ListOperand[int64]{addedBacklinkIDs}}),
		})
		if err != nil {
			return err
		}

		addedBacklinks, err := goaoi.TransformCopySliceUnsafe(addedBacklinkDocuments, (*domain.Document).GetPath)
		if err != nil {
			return err
		}

		addedPathBacklinks[i] = tuple.T2[string, []string]{oldDocument.Path, addedBacklinks}

		removedBacklinkIDs, err := GetRemovedBacklinks(oldDocument, newDocument)
		if err != nil {
			return err
		}

		removedBacklinkDocuments, err := documentManager.GetWhere(ctx, &domain.DocumentFilter{
			LinkedDocumentIDs: optional.Make(model.FilterOperation[int64]{
				Operator: model.FilterIn,
				Operand:  model.ListOperand[int64]{removedBacklinkIDs}}),
		})
		if err != nil {
			return err
		}

		removedBacklinks, err := goaoi.TransformCopySliceUnsafe(removedBacklinkDocuments, (*domain.Document).GetPath)
		if err != nil {
			return err
		}

		removedPathBacklinks[i] = tuple.T2[string, []string]{oldDocument.Path, removedBacklinks}

		addedTagIDs, err := GetAddedTags(oldDocument, newDocument)
		if err != nil {
			return err
		}

		addedTagTags, err := documentManager.Repository.GetTagRepository().GetWhere(ctx, &domain.TagFilter{
			ID: optional.Make(model.FilterOperation[int64]{
				Operator: model.FilterIn,
				Operand:  model.ListOperand[int64]{addedTagIDs}}),
		})
		if err != nil {
			return err
		}

		addedTags, err := goaoi.TransformCopySliceUnsafe(addedTagTags, (*domain.Tag).GetTag)
		if err != nil {
			return err
		}

		addedPathTags[i] = tuple.T2[string, []string]{oldDocument.Path, addedTags}

		removedTagIDs, err := GetRemovedTags(oldDocument, newDocument)
		if err != nil {
			return err
		}

		removedTagTags, err := documentManager.Repository.GetTagRepository().GetWhere(ctx, &domain.TagFilter{
			ID: optional.Make(model.FilterOperation[int64]{
				Operator: model.FilterIn,
				Operand:  model.ListOperand[int64]{removedTagIDs}}),
		})
		if err != nil {
			return err
		}

		removedTags, err := goaoi.TransformCopySliceUnsafe(removedTagTags, (*domain.Tag).GetTag)
		if err != nil {
			return err
		}

		removedPathTags[i] = tuple.T2[string, []string]{oldDocument.Path, removedTags}
	}

	err := m.AddLinks(context.Background(), addedPathLinks)
	if err != nil {
		return err
	}
	err = m.RemoveLinks(context.Background(), removedPathLinks)
	if err != nil {
		return err
	}

	err = m.AddBackLinks(context.Background(), addedPathBacklinks)
	if err != nil {
		return err
	}

	err = m.RemoveBackLinks(context.Background(), removedPathBacklinks)
	if err != nil {
		return err
	}

	err = m.AddBackLinks(context.Background(), addedPathTags)
	if err != nil {
		return err
	}

	err = m.RemoveBackLinks(context.Background(), removedPathTags)
	if err != nil {
		return err
	}

	return err
}

// TODO: This mess needs a lot of cleaning up
func (m *DocumentContentManager) UpdateDocumentContentsFromFilterAndUpdater(ctx context.Context, filter *domain.DocumentFilter, updater *domain.DocumentUpdater, documentManager *DocumentManager) error {
	linksExtractorFromID := func(oldDocumentID int64) string {
		document, err := documentManager.GetFirstWhere(ctx, &domain.DocumentFilter{ID: optional.Make(model.FilterOperation[int64]{Operator: model.FilterEqual, Operand: model.ScalarOperand[int64]{Operand: oldDocumentID}})})
		if err != nil {

			// FIX: We should have proper error handling here
			panic(err)
		}

		return document.Path
	}
	tagsExtractorFromID := func(oldDocumentID int64) string {
		document, err := documentManager.Repository.GetTagRepository().GetFirstWhere(ctx, &domain.TagFilter{ID: optional.Make(model.FilterOperation[int64]{Operator: model.FilterEqual, Operand: model.ScalarOperand[int64]{Operand: oldDocumentID}})})
		if err != nil {

			// FIX: We should have proper error handling here
			panic(err)
		}

		return document.Tag
	}

	oldDocuments, err := documentManager.GetWhere(ctx, filter)

	if updater.LinkedDocumentIDs.HasValue {
		if updater.LinkedDocumentIDs.Wrappee.Operator == model.UpdateAppend || updater.LinkedDocumentIDs.Wrappee.Operator == model.UpdatePrepend {
			addedPathLinks := make([]tuple.T2[string, []string], 0, 10)

			addedLinks, err := goaoi.TransformCopySliceUnsafe(updater.LinkedDocumentIDs.Wrappee.Operand, linksExtractorFromID)
			if err != nil {
				return err
			}

			for _, oldDocument := range oldDocuments {
				addedPathLinks = append(addedPathLinks, tuple.T2[string, []string]{oldDocument.Path, addedLinks})
			}

			err = m.AddLinks(context.Background(), addedPathLinks)
			if err != nil {
				return err
			}
		} else if updater.LinkedDocumentIDs.Wrappee.Operator == model.UpdateClear {
			err := m.handleClearLinks(ctx, updater.LinkedDocumentIDs.Wrappee.Operand, linksExtractorFromID, oldDocuments)
			if err != nil {
				return err
			}
		} else if updater.LinkedDocumentIDs.Wrappee.Operator == model.UpdateSet {
			err := m.handleClearLinks(ctx, updater.LinkedDocumentIDs.Wrappee.Operand, linksExtractorFromID, oldDocuments)
			if err != nil {
				return err
			}
			err = m.handleSetLinks(ctx, updater.LinkedDocumentIDs.Wrappee.Operand, linksExtractorFromID, oldDocuments)
			if err != nil {
				return err
			}
		}
	}

	if updater.BacklinkedDocumentsIDs.HasValue {
		if updater.BacklinkedDocumentsIDs.Wrappee.Operator == model.UpdateAppend || updater.BacklinkedDocumentsIDs.Wrappee.Operator == model.UpdatePrepend {
			addedPathBacklinks := make([]tuple.T2[string, []string], 0, 10)

			addedBacklinks, err := goaoi.TransformCopySliceUnsafe(updater.BacklinkedDocumentsIDs.Wrappee.Operand, linksExtractorFromID)
			if err != nil {
				return err
			}

			for _, oldDocument := range oldDocuments {
				addedPathBacklinks = append(addedPathBacklinks, tuple.T2[string, []string]{oldDocument.Path, addedBacklinks})
			}

			err = m.AddBackLinks(context.Background(), addedPathBacklinks)
			if err != nil {
				return err
			}
		} else if updater.BacklinkedDocumentsIDs.Wrappee.Operator == model.UpdateClear {
			err := m.handleClearBacklinks(ctx, updater.BacklinkedDocumentsIDs.Wrappee.Operand, linksExtractorFromID, oldDocuments)
			if err != nil {
				return err
			}
		} else if updater.BacklinkedDocumentsIDs.Wrappee.Operator == model.UpdateSet {
			err := m.handleClearBacklinks(ctx, updater.BacklinkedDocumentsIDs.Wrappee.Operand, linksExtractorFromID, oldDocuments)
			if err != nil {
				return err
			}
			err = m.handlePushBacklinks(ctx, updater.BacklinkedDocumentsIDs.Wrappee.Operand, linksExtractorFromID, oldDocuments)
			if err != nil {
				return err
			}
		}
	}

	if updater.TagIDs.HasValue {
		if updater.TagIDs.Wrappee.Operator == model.UpdateAppend || updater.TagIDs.Wrappee.Operator == model.UpdatePrepend {
			addedPathTags := make([]tuple.T2[string, []string], 0, 10)

			addedTags, err := goaoi.TransformCopySliceUnsafe(updater.TagIDs.Wrappee.Operand, tagsExtractorFromID)
			if err != nil {
				return err
			}

			for _, oldDocument := range oldDocuments {
				addedPathTags = append(addedPathTags, tuple.T2[string, []string]{oldDocument.Path, addedTags})
			}

			err = m.AddTags(context.Background(), addedPathTags)
			if err != nil {
				return err
			}
		} else if updater.TagIDs.Wrappee.Operator == model.UpdateClear {
			err := m.handleClearTags(ctx, updater.TagIDs.Wrappee.Operand, tagsExtractorFromID, oldDocuments)
			if err != nil {
				return err
			}
		} else if updater.TagIDs.Wrappee.Operator == model.UpdateSet {
			err := m.handleClearTags(ctx, updater.TagIDs.Wrappee.Operand, tagsExtractorFromID, oldDocuments)
			if err != nil {
				return err
			}
			err = m.handleSetTags(ctx, updater.TagIDs.Wrappee.Operand, tagsExtractorFromID, oldDocuments)
			if err != nil {
				return err
			}
		}

	}

	return err

}

func (m *DocumentContentManager) handleClearLinks(ctx context.Context, documents []int64, linksExtractor func(oldDocument int64) string, oldDocuments []*domain.Document) error {
	removedPathLinks := make([]tuple.T2[string, []string], 0, 10)

	removedLinks, err := goaoi.TransformCopySliceUnsafe(documents, linksExtractor)
	if err != nil {
		return err
	}

	for _, oldDocument := range oldDocuments {
		removedPathLinks = append(removedPathLinks, tuple.T2[string, []string]{oldDocument.Path, removedLinks})
	}
	err = m.RemoveLinks(ctx, removedPathLinks)
	if err != nil {
		return err
	}
	return nil
}

func (m *DocumentContentManager) handleSetLinks(ctx context.Context, documents []int64, linksExtractor func(oldDocument int64) string, oldDocuments []*domain.Document) error {
	addedPathLinks := make([]tuple.T2[string, []string], 0, 10)

	addedLinks, err := goaoi.TransformCopySliceUnsafe(documents, linksExtractor)
	if err != nil {
		return err
	}

	for _, oldDocument := range oldDocuments {
		addedPathLinks = append(addedPathLinks, tuple.T2[string, []string]{oldDocument.Path, addedLinks})
	}
	err = m.AddLinks(ctx, addedPathLinks)
	if err != nil {
		return err
	}
	return nil
}

func (m *DocumentContentManager) handleClearBacklinks(ctx context.Context, documents []int64, linksExtractor func(oldDocument int64) string, oldDocuments []*domain.Document) error {
	removedPathBacklinks := make([]tuple.T2[string, []string], 0, 10)

	removedBacklinks, err := goaoi.TransformCopySliceUnsafe(documents, linksExtractor)
	if err != nil {
		return err
	}

	for _, oldDocument := range oldDocuments {
		removedPathBacklinks = append(removedPathBacklinks, tuple.T2[string, []string]{oldDocument.Path, removedBacklinks})
	}
	err = m.RemoveBackLinks(ctx, removedPathBacklinks)
	if err != nil {
		return err
	}
	return nil
}

func (m *DocumentContentManager) handlePushBacklinks(ctx context.Context, documents []int64, linksExtractor func(oldDocument int64) string, oldDocuments []*domain.Document) error {
	addedPathBacklinks := make([]tuple.T2[string, []string], 0, 10)

	addedBacklinks, err := goaoi.TransformCopySliceUnsafe(documents, linksExtractor)
	if err != nil {
		return err
	}

	for _, oldDocument := range oldDocuments {
		addedPathBacklinks = append(addedPathBacklinks, tuple.T2[string, []string]{oldDocument.Path, addedBacklinks})
	}
	err = m.AddBackLinks(ctx, addedPathBacklinks)
	if err != nil {
		return err
	}
	return nil
}

func (m *DocumentContentManager) handleClearTags(ctx context.Context, tags []int64, tagsExtractor func(oldDocument int64) string, oldDocuments []*domain.Document) error {
	removedPathTags := make([]tuple.T2[string, []string], 0, 10)

	removedBacklinks, err := goaoi.TransformCopySliceUnsafe(tags, tagsExtractor)
	if err != nil {
		return err
	}

	for _, oldDocument := range oldDocuments {
		removedPathTags = append(removedPathTags, tuple.T2[string, []string]{oldDocument.Path, removedBacklinks})
	}
	err = m.RemoveTags(ctx, removedPathTags)
	if err != nil {
		return err
	}
	return nil
}

func (m *DocumentContentManager) handleSetTags(ctx context.Context, tags []int64, tagsExtractor func(oldDocument int64) string, oldDocuments []*domain.Document) error {
	addedPathTags := make([]tuple.T2[string, []string], 0, 10)

	addedBacklinks, err := goaoi.TransformCopySliceUnsafe(tags, tagsExtractor)
	if err != nil {
		return err
	}

	for _, oldDocument := range oldDocuments {
		addedPathTags = append(addedPathTags, tuple.T2[string, []string]{oldDocument.Path, addedBacklinks})
	}
	err = m.AddTags(ctx, addedPathTags)
	if err != nil {
		return err
	}
	return nil
}
