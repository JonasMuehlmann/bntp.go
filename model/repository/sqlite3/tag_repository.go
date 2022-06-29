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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/sql_repository.go.tpl

package repository

import (
	 repoCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
	"container/list"
	"fmt"
	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/volatiletech/null/v8"
    "context"
    "database/sql"
    "github.com/volatiletech/sqlboiler/v4/boil"
    "github.com/volatiletech/sqlboiler/v4/queries"
    "github.com/volatiletech/sqlboiler/v4/queries/qm"
    log "github.com/sirupsen/logrus"
    
    "strconv"
    "strings"
    
    
)


//******************************************************************//
//                        Types and constants                       //
//******************************************************************//
type Sqlite3TagRepository struct {
    db *sql.DB
    
}

type TagField string

var TagFields = struct {
    Tag  TagField
    Path  TagField
    Children  TagField
    ParentTag  TagField
    ID  TagField
    
}{
    Tag: "tag",
    Path: "path",
    Children: "children",
    ParentTag: "parent_tag",
    ID: "id",
    
}

var TagFieldsList = []TagField{
    TagField("Tag"),
    TagField("Path"),
    TagField("Children"),
    TagField("ParentTag"),
    TagField("ID"),
    
}

var TagRelationsList = []string{
    "ParentTagTag",
    "Bookmarks",
    "Documents",
    "ParentTagTags",
    
}

type TagFilter struct {
    Tag optional.Optional[model.FilterOperation[string]]
    Path optional.Optional[model.FilterOperation[string]]
    Children optional.Optional[model.FilterOperation[string]]
    ParentTag optional.Optional[model.FilterOperation[null.Int64]]
    ID optional.Optional[model.FilterOperation[int64]]
    
    ParentTagTag optional.Optional[model.FilterOperation[*Tag]]
    Bookmarks optional.Optional[model.FilterOperation[*Bookmark]]
    Documents optional.Optional[model.FilterOperation[*Document]]
    ParentTagTags optional.Optional[model.FilterOperation[*Tag]]
    
}

type TagFilterMapping[T any] struct {
    Field TagField
    FilterOperation model.FilterOperation[T]
}

func (filter *TagFilter) GetSetFilters() *list.List {
    setFilters := list.New()

    if filter.Tag.HasValue {
    setFilters.PushBack(TagFilterMapping[string]{Field: TagFields.Tag, FilterOperation: filter.Tag.Wrappee})
    }
    if filter.Path.HasValue {
    setFilters.PushBack(TagFilterMapping[string]{Field: TagFields.Path, FilterOperation: filter.Path.Wrappee})
    }
    if filter.Children.HasValue {
    setFilters.PushBack(TagFilterMapping[string]{Field: TagFields.Children, FilterOperation: filter.Children.Wrappee})
    }
    if filter.ParentTag.HasValue {
    setFilters.PushBack(TagFilterMapping[null.Int64]{Field: TagFields.ParentTag, FilterOperation: filter.ParentTag.Wrappee})
    }
    if filter.ID.HasValue {
    setFilters.PushBack(TagFilterMapping[int64]{Field: TagFields.ID, FilterOperation: filter.ID.Wrappee})
    }
    

    return setFilters
}

type TagUpdater struct {
    Tag optional.Optional[model.UpdateOperation[string]]
    Path optional.Optional[model.UpdateOperation[string]]
    Children optional.Optional[model.UpdateOperation[string]]
    ParentTag optional.Optional[model.UpdateOperation[null.Int64]]
    ID optional.Optional[model.UpdateOperation[int64]]
    
    ParentTagTag optional.Optional[model.UpdateOperation[*Tag]]
    Bookmarks optional.Optional[model.UpdateOperation[BookmarkSlice]]
    Documents optional.Optional[model.UpdateOperation[DocumentSlice]]
    ParentTagTags optional.Optional[model.UpdateOperation[TagSlice]]
    
}

type TagUpdaterMapping[T any] struct {
    Field TagField
    Updater model.UpdateOperation[T]
}

func (updater *TagUpdater) GetSetUpdaters() *list.List {
    setUpdaters := list.New()

    if updater.Tag.HasValue {
    setUpdaters.PushBack(TagUpdaterMapping[string]{Field: TagFields.Tag, Updater: updater.Tag.Wrappee})
    }
    if updater.Path.HasValue {
    setUpdaters.PushBack(TagUpdaterMapping[string]{Field: TagFields.Path, Updater: updater.Path.Wrappee})
    }
    if updater.Children.HasValue {
    setUpdaters.PushBack(TagUpdaterMapping[string]{Field: TagFields.Children, Updater: updater.Children.Wrappee})
    }
    if updater.ParentTag.HasValue {
    setUpdaters.PushBack(TagUpdaterMapping[null.Int64]{Field: TagFields.ParentTag, Updater: updater.ParentTag.Wrappee})
    }
    if updater.ID.HasValue {
    setUpdaters.PushBack(TagUpdaterMapping[int64]{Field: TagFields.ID, Updater: updater.ID.Wrappee})
    }
    

    return setUpdaters
}

func (updater *TagUpdater) ApplyToModel(tagModel *Tag) {
    if updater.Tag.HasValue {
        model.ApplyUpdater(&(*tagModel).Tag, updater.Tag.Wrappee)
    }
    if updater.Path.HasValue {
        model.ApplyUpdater(&(*tagModel).Path, updater.Path.Wrappee)
    }
    if updater.Children.HasValue {
        model.ApplyUpdater(&(*tagModel).Children, updater.Children.Wrappee)
    }
    if updater.ParentTag.HasValue {
        model.ApplyUpdater(&(*tagModel).ParentTag, updater.ParentTag.Wrappee)
    }
    if updater.ID.HasValue {
        model.ApplyUpdater(&(*tagModel).ID, updater.ID.Wrappee)
    }
    
}

type queryModSliceTag []qm.QueryMod

func (s queryModSliceTag) Apply(q *queries.Query) {
    qm.Apply(q, s...)
}

func buildQueryModFilterTag[T any](filterField TagField, filterOperation model.FilterOperation[T]) queryModSliceTag {
    var newQueryMod queryModSliceTag

    filterOperator := filterOperation.Operator

    switch filterOperator {
    case model.FilterEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" = ?", filterOperand.Operand))
    case model.FilterNEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterNEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" != ?", filterOperand.Operand))
    case model.FilterGreaterThan:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterGreaterThan operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" > ?", filterOperand.Operand))
    case model.FilterGreaterThanEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterGreaterThanEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" >= ?", filterOperand.Operand))
    case model.FilterLessThan:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterLessThan operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" < ?", filterOperand.Operand))
    case model.FilterLessThanEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterLessThanEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" <= ?", filterOperand.Operand))
    case model.FilterIn:
        filterOperand, ok := filterOperation.Operand.(model.ListOperand[any])
        if !ok {
            panic("expected a list operand for FilterIn operator")
        }

        newQueryMod = append(newQueryMod, qm.WhereIn(string(filterField)+" IN (?)", filterOperand.Operands))
    case model.FilterNotIn:
        filterOperand, ok := filterOperation.Operand.(model.ListOperand[any])
        if !ok {
            panic("expected a list operand for FilterNotIn operator")
        }

        newQueryMod = append(newQueryMod, qm.WhereNotIn(string(filterField)+" IN (?)", filterOperand.Operands))
    case model.FilterBetween:
        filterOperand, ok := filterOperation.Operand.(model.RangeOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterBetween operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
    case model.FilterNotBetween:
        filterOperand, ok := filterOperation.Operand.(model.RangeOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterNotBetween operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" NOT BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
    case model.FilterLike:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterLike operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" LIKE ?", filterOperand.Operand))
    case model.FilterNotLike:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterLike operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" NOT LIKE ?", filterOperand.Operand))
    case model.FilterOr:
        filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterOr operator")
        }
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Or2(qm.Expr(buildQueryModFilterTag(filterField, filterOperand.RHS))))
    case model.FilterAnd:
        filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
        if !ok {
            panic("expected a scalar operand for FilterAnd operator")
        }

        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.RHS)))
    default:
        panic("Unhandled FilterOperator")
    }

    return newQueryMod
}

func buildQueryModListFromFilterTag(setFilters list.List) queryModSliceTag {
	queryModList := make(queryModSliceTag, 0, 5)

	for filter := setFilters.Front(); filter != nil; filter = filter.Next() {
		filterMapping, ok := filter.Value.(TagFilterMapping[any])
		if !ok {
			panic(fmt.Sprintf("expected type %T but got %T", TagFilterMapping[any]{}, filter))
		}

        newQueryMod := buildQueryModFilterTag(filterMapping.Field, filterMapping.FilterOperation)

        queryModList = append(queryModList, newQueryMod...)
	}

	return queryModList
}


type Sqlite3TagRepositoryConstructorArgs struct {
    DB *sql.DB
    
}

func (repo *Sqlite3TagRepository) New(args any) (newRepo repoCommon.TagRepository, err error) {
    constructorArgs, ok := args.(Sqlite3TagRepositoryConstructorArgs)
    if !ok {
        err = fmt.Errorf("expected type %T but got %T", Sqlite3TagRepositoryConstructorArgs{}, args)

        return
    }

    repo.db = constructorArgs.DB
    

    newRepo = repo

    return
}


//******************************************************************//
//                              Methods                             //
//******************************************************************//
func (repo *Sqlite3TagRepository) Add(ctx context.Context, domainModels []*domain.Tag)  (err error){
    if len(domainModels) == 0 {
        log.Debug(helper.LogMessageEmptyInput)

        return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
    }

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil{
		err = helper.NilInputError{}
		log.Error(err)

		return err
	}

    var repositoryModels []any
    repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

    var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
        repoModel, ok := repositoryModel.(*Tag)
        if !ok {
            return fmt.Errorf("expected type *Tag but got %T", repoModel)
        }

		err = repoModel.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	tx.Commit()

    return nil
}

func (repo *Sqlite3TagRepository) Replace(ctx context.Context, domainModels []*domain.Tag)  (err error){
    
    if len(domainModels) == 0 {
        log.Debug(helper.LogMessageEmptyInput)

        return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
    }

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil{
		err = helper.NilInputError{}
		log.Error(err)

		return err
	}

    var repositoryModels []any
    repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

    var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
        repoModel, ok := repositoryModel.(*Tag)
        if !ok {
            return fmt.Errorf("expected type *Tag but got %T", repoModel)
        }

        var numAffectedRecords int64
		numAffectedRecords, err = repoModel.Update(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}

        if numAffectedRecords == 0 {
            return helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError}
        }
	}

	tx.Commit()

    return nil
}
func (repo *Sqlite3TagRepository) Upsert(ctx context.Context, domainModels []*domain.Tag)  (err error){
    if len(domainModels) == 0 {
        log.Debug(helper.LogMessageEmptyInput)

        return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
    }

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil{
		err = helper.NilInputError{}
		log.Error(err)

		return err
	}

    var repositoryModels []any
    repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

    var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
        repoModel, ok := repositoryModel.(*Tag)
        if !ok {
            return fmt.Errorf("expected type *Tag but got %T", repoModel)
        }

        
		err = repoModel.Upsert(ctx, tx, false, []string{}, boil.Infer(), boil.Infer())
        
		if err != nil {
			return err
		}
	}

	tx.Commit()

    return nil
}

func (repo *Sqlite3TagRepository) Update(ctx context.Context, domainModels []*domain.Tag, domainColumnUpdater *domain.TagUpdater)  (err error){
    if len(domainModels) == 0 {
        log.Debug(helper.LogMessageEmptyInput)

        return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
    }

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil{
		err = helper.NilInputError{}
		log.Error(err)

		return err
	}

	if  domainColumnUpdater == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return err
    }

    var repositoryModels []any
    repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

    var repositoryUpdater any
    repositoryUpdater, err = repo.TagDomainToRepositoryUpdater(ctx, domainColumnUpdater)
    if err != nil {
        return err
    }

    var tx *sql.Tx

   	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

    var numAffectedRecords int64
    for _, repositoryModel := range   repositoryModels {
        repoModel, ok := repositoryModel.(*Tag)
        if !ok {
            return fmt.Errorf("expected type *Tag but got %T", repoModel)
        }

        repoUpdater, ok := repositoryUpdater.(*TagUpdater)
        if !ok {
            return fmt.Errorf("expected type *Tag but got %T", repoModel)
        }

        repoUpdater.ApplyToModel(repoModel)
        numAffectedRecords, err = repoModel.Update(ctx, tx, boil.Infer())
        if err != nil {
            return err
        }

        if numAffectedRecords == 0 {
            return helper.IneffectiveOperationError{Inner: helper.NonExistentPrimaryDataError}
        }
    }

    err = tx.Commit()

    return err
}

func (repo *Sqlite3TagRepository) UpdateWhere(ctx context.Context, domainColumnFilter *domain.TagFilter, domainColumnUpdater *domain.TagUpdater) (numAffectedRecords int64, err error) {
	var modelsToUpdate TagSlice

	if  domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return 0, err
    }

	if  domainColumnUpdater == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return 0, err
    }

    var repositoryFilter any
    repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
    if err != nil {
        return
    }

    var repositoryUpdater any
    repositoryUpdater, err = repo.TagDomainToRepositoryUpdater(ctx, domainColumnUpdater)
    if err != nil {
        return
    }

    repoUpdater, ok := repositoryUpdater.(*TagUpdater)
    if !ok {
        err = fmt.Errorf("expected type *TagUpdater but got %T", repoUpdater)

        return
    }


    repoFilter, ok := repositoryFilter.(*TagFilter)
    if !ok {
        err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

        return
    }

    setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	modelsToUpdate, err = Tags(queryFilters...).All(ctx, repo.db)
	if err != nil {
		return
	}

    numAffectedRecords = int64(len(modelsToUpdate))

    var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

    for _, repoModel := range modelsToUpdate {
        repoUpdater.ApplyToModel(repoModel)
        repoModel.Update(ctx, tx, boil.Infer())
    }

    tx.Commit()

    return
}

func (repo *Sqlite3TagRepository) Delete(ctx context.Context, domainModels []*domain.Tag)  (err error){
    if len(domainModels) == 0 {
        log.Debug(helper.LogMessageEmptyInput)

        return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
    }

	err = goaoi.AnyOfSlice(domainModels, goaoi.AreEqualPartial[*domain.Tag](nil))
	if err == nil{
		err = helper.NilInputError{}
		log.Error(err)

		return err
	}

    var repositoryModels []any
    repositoryModels, err = goaoi.TransformCopySlice(domainModels, repo.GetTagDomainToRepositoryModel(ctx))
	if err != nil {
		return err
	}

    var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
        repoModel, ok := repositoryModel.(*Tag)
        if !ok {
            return fmt.Errorf("expected type *Tag but got %T", repoModel)
        }

		_, err = repoModel.Delete(ctx, tx)
		if err != nil {
			return err
		}
	}

	tx.Commit()

    return nil
}

func (repo *Sqlite3TagRepository) DeleteWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (numAffectedRecords int64, err error) {
	if  domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return 0, err
    }

    var repositoryFilter any
    repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
    if err != nil {
        return
    }

    repoFilter, ok := repositoryFilter.(*TagFilter)
    if !ok {
        err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

        return
    }

    setFilters := * repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

    var tx *sql.Tx

	tx, err = repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Tags(queryFilters...).DeleteAll(ctx, tx)

    tx.Commit()

    return
}

func (repo *Sqlite3TagRepository) CountWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (numRecords int64, err error) {
	if  domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return 0, err
    }

    var repositoryFilter any
    repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
    if err != nil {
        return 0, err
    }

    repoFilter, ok := repositoryFilter.(*TagFilter)
    if !ok {
        return 0, fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

    }

    setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	return Tags(queryFilters...).Count(ctx, repo.db)
}

func (repo *Sqlite3TagRepository) CountAll(ctx context.Context) (numRecords int64, err error) {
	return Tags().Count(ctx, repo.db)
}

func (repo *Sqlite3TagRepository) DoesExist(ctx context.Context, domainModel *domain.Tag) (doesExist bool, err error) {
	if domainModel == nil {
        err = helper.NilInputError{}
		log.Error(err)

		return
	}

    var repositoryModel any
    repositoryModel, err = repo.TagDomainToRepositoryModel(ctx, domainModel)
    if err != nil {
        return
    }

    repoModel, ok := repositoryModel.(*Tag)
    if !ok {
        err = fmt.Errorf("expected type *Tag but got %T", repoModel)
        return
    }


	return TagExists(ctx, repo.db, repoModel.ID)
}

func (repo *Sqlite3TagRepository) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (doesExist bool, err error) {
	if  domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
    }

    var repositoryFilter any
    repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
    if err != nil {
        return
    }

    repoFilter, ok := repositoryFilter.(*TagFilter)
    if !ok {
        err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

        return
    }

    setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	return Tags(queryFilters...).Exists(ctx, repo.db)
}

func (repo *Sqlite3TagRepository) GetWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (records []*domain.Tag, err error) {
	if  domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
    }

    var repositoryFilter any
    repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
    if err != nil {
        return
    }

    repoFilter, ok := repositoryFilter.(*TagFilter)
    if !ok {
        err = fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

        return
    }


    setFilters := *repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

    var repositoryModels TagSlice
    repositoryModels, err = Tags(queryFilters...).All(ctx, repo.db)

    records = make([]*domain.Tag, 0, len(repositoryModels))

    var domainModel *domain.Tag
    for _, repoModel := range repositoryModels {
        domainModel, err = repo.TagRepositoryToDomainModel(ctx, repoModel)
        if err != nil {
            return
        }

        records = append(records, domainModel)
    }

    return
}

func (repo *Sqlite3TagRepository) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (record *domain.Tag, err error) {
	if  domainColumnFilter == nil {
		err = helper.NilInputError{}
		log.Error(err)

		return
    }

    var repositoryFilter any
    repositoryFilter, err = repo.TagDomainToRepositoryFilter(ctx, domainColumnFilter)
    if err != nil {
        return
    }

    repoFilter, ok := repositoryFilter.(*TagFilter)
    if !ok {
        err =  fmt.Errorf("expected type *TagFilter but got %T", repoFilter)

        return
    }

    setFilters := * repoFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

    var repositoryModel *Tag
    repositoryModel, err = Tags(queryFilters...).One(ctx, repo.db)
    if err != nil {
        return
    }

    record , err =repo.TagRepositoryToDomainModel(ctx, repositoryModel)

    return
}

func (repo *Sqlite3TagRepository) GetAll(ctx context.Context) (records []*domain.Tag, err error) {
    var repositoryModels TagSlice
    repositoryModels, err = Tags().All(ctx, repo.db)
    if err != nil {
        return
    }

    records = make([]*domain.Tag, 0, len(repositoryModels))

    var domainModel *domain.Tag
    for _, repoModel := range repositoryModels {
        domainModel, err = repo.TagRepositoryToDomainModel(ctx, repoModel)
        if err != nil {
            return
        }

        records = append(records, domainModel)
    }

    return
}




//******************************************************************//
//                            Converters                            //
//******************************************************************//
func (repo *Sqlite3TagRepository) GetTagDomainToRepositoryModel(ctx context.Context) func(domainModel *domain.Tag) (repositoryModel any, err error) {
    return func(domainModel *domain.Tag) (repositoryModel any, err error) {
        return repo.TagDomainToRepositoryModel(ctx, domainModel)
    }
}

func (repo *Sqlite3TagRepository) GetTagRepositoryToDomainModel(ctx context.Context) func(repositoryModel any) (domainModel *domain.Tag, err error) {
    return func(repositoryModel any) (domainModel *domain.Tag, err error) {

        return repo.TagRepositoryToDomainModel(ctx,repositoryModel)
    }
}


//******************************************************************//
//                          Model Converter                         //
//******************************************************************//



func (repo *Sqlite3TagRepository) TagDomainToRepositoryModel(ctx context.Context, domainModel *domain.Tag) (repositoryModel any, err error)  {

// TODO: make sure to insert all tags in ParentPath and Subtags into db
    repositoryModelConcrete := new(Tag)
    repositoryModelConcrete.R = repositoryModelConcrete.R.NewStruct()

    repositoryModelConcrete.ID = domainModel.ID
    repositoryModelConcrete.Tag = domainModel.Tag


//***********************    Set ParentTag    **********************//
    if len(domainModel.ParentPath) > 0 {
        repositoryModelConcrete.ParentTag = null.NewInt64(domainModel.ParentPath[len(domainModel.ParentPath) - 1].ID, true)
    }

//*************************    Set Path    *************************//
for _, tag := range domainModel.ParentPath {
    repositoryModelConcrete.Path += strconv.FormatInt(tag.ID, 10) + ";"
}

repositoryModelConcrete.Path += strconv.FormatInt(domainModel.ID, 10)

//************************    Set Children  ************************//
for _, tag := range domainModel.Subtags {
    repositoryModelConcrete.Children += strconv.FormatInt(tag.ID, 10) + ";"
}

    repositoryModel = repositoryModelConcrete

    return
}

// TODO: These functions should be context aware
func (repo *Sqlite3TagRepository) TagRepositoryToDomainModel(ctx context.Context, repositoryModel any) (domainModel *domain.Tag, err error) {
// TODO: make sure to insert all tags in ParentPath and Subtags into db
    domainModel = new(domain.Tag)

    repositoryModelConcrete := repositoryModel.(Tag)

    domainModel.ID = repositoryModelConcrete.ID
    domainModel.Tag = repositoryModelConcrete.Tag

//***********************    Set ParentPath    **********************//
var parentTagID int64
var parentTag *Tag
var domainParentTag *domain.Tag

for _, parentTagIDRaw := range strings.Split(repositoryModelConcrete.Path, ";")[:len(repositoryModelConcrete.Path)-2]{
    parentTagID, err = strconv.ParseInt(parentTagIDRaw, 10, 64)
    if err != nil {
        return
    }

    parentTag, err = Tags(TagWhere.ID.EQ(parentTagID)).One(ctx, repo.db)
    if err != nil {
        return
    }

    domainParentTag, err = repo.TagRepositoryToDomainModel(ctx, parentTag)
    if err != nil {
        return
    }

    domainModel.ParentPath = append(domainModel.ParentPath, domainParentTag)
}

//************************    Set Subtags ************************//
var childTagID int64
var childTag *Tag
var domainChildTag *domain.Tag

for _, childTagIDRaw := range strings.Split(repositoryModelConcrete.Children, ";")[:len(repositoryModelConcrete.Children)-2]{
    childTagID, err = strconv.ParseInt(childTagIDRaw, 10, 64)
    if err != nil {
        return
    }

    childTag, err = Tags(TagWhere.ID.EQ(childTagID)).One(ctx, repo.db)
    if err != nil {
        return
    }

    domainChildTag, err = repo.TagRepositoryToDomainModel(ctx, childTag)
    if err != nil {
        return
    }

    domainModel.Subtags = append(domainModel.Subtags, domainChildTag)
}

    repositoryModel = repositoryModelConcrete

    return
}

//******************************************************************//
//                         Filter Converter                         //
//******************************************************************//



func (repo *Sqlite3TagRepository) TagDomainToRepositoryFilter(ctx context.Context, domainFilter *domain.TagFilter) (repositoryFilter any, err error)  {
    repositoryFilterConcrete := new(TagFilter)

	repositoryFilterConcrete.ID = domainFilter.ID
	repositoryFilterConcrete.Tag = domainFilter.Tag

	if domainFilter.ParentPath.HasValue {

		//*********************    Set ParentPath    *********************//
		var convertedParentTagTagFilter model.FilterOperation[*Tag]

		convertedParentTagTagFilter, err = model.ConvertFilter[*Tag, *domain.Tag](domainFilter.ParentPath.Wrappee, repoCommon.MakeDomainToRepositoryEntityConverterGeneric[domain.Tag, Tag](ctx, repo.TagDomainToRepositoryModel))
		if err != nil {
			return
		}

		repositoryFilterConcrete.ParentTagTag.Push(convertedParentTagTagFilter)
		//*************************    Set Path    *************************//
		var convertedPathFilter model.FilterOperation[string]

		convertedPathFilter, err = model.ConvertFilter[string, *domain.Tag](domainFilter.ParentPath.Wrappee, func(tag *domain.Tag) (string, error) { return strconv.FormatInt(tag.ID, 10), nil })
		if err != nil {
			return
		}

		repositoryFilterConcrete.Path.Push(convertedPathFilter)
		//**********************    Set ParentTag    ***********************//
		var convertedParentTagFilter model.FilterOperation[null.Int64]

		convertedParentTagFilter, err = model.ConvertFilter[null.Int64, *domain.Tag](domainFilter.ParentPath.Wrappee, func(tag *domain.Tag) (null.Int64, error) { return null.NewInt64(tag.ID, true), nil })
		if err != nil {
			return
		}

		repositoryFilterConcrete.ParentTag.Push(convertedParentTagFilter)
	}

	//**********************    Set child tags *********************//
	if domainFilter.Subtags.HasValue {
		var convertedFilter model.FilterOperation[string]

		convertedFilter, err = model.ConvertFilter[string, *domain.Tag](domainFilter.Subtags.Wrappee, func(tag *domain.Tag) (string, error) { return strconv.FormatInt(tag.ID, 10), nil })
		if err != nil {
			return
		}

		repositoryFilterConcrete.Children.Push(convertedFilter)
	}

    repositoryFilter = repositoryFilterConcrete

	return
}

//******************************************************************//
//                         Updater Converter                        //
//******************************************************************//



func (repo *Sqlite3TagRepository) TagDomainToRepositoryUpdater(ctx context.Context, domainUpdater *domain.TagUpdater) (repositoryUpdater any, err error)  {
    repositoryUpdaterConcrete := new(TagUpdater)

	//**************************    Set tag    *************************//
	if domainUpdater.Tag.HasValue {
		repositoryUpdaterConcrete.Tag.Push(model.UpdateOperation[string]{Operator: domainUpdater.Tag.Wrappee.Operator, Operand: repositoryUpdaterConcrete.Tag.Wrappee.Operand})
	}

	//***********    Set ParentPath    ***********//
	if domainUpdater.ParentPath.HasValue {
		var convertedTagRaw any
		tag := domainUpdater.ParentPath.Wrappee.Operand[len(domainUpdater.ParentPath.Wrappee.Operand)-1]
		convertedTagRaw, err =  repo.TagDomainToRepositoryModel(ctx, tag)
		if err != nil {
			return
		}

		repositoryUpdaterConcrete.ParentTagTag.Push(model.UpdateOperation[*Tag]{Operator: domainUpdater.ParentPath.Wrappee.Operator, Operand: convertedTagRaw.(*Tag)})
		repositoryUpdaterConcrete.ParentTag.Push(model.UpdateOperation[null.Int64]{Operator: domainUpdater.ParentPath.Wrappee.Operator, Operand: null.NewInt64(convertedTagRaw.(*Tag).ID, true)})

		pathIDs := make([]string, 0, len(domainUpdater.ParentPath.Wrappee.Operand)+1)
		for _, tag := range domainUpdater.ParentPath.Wrappee.Operand {
			pathIDs = append(pathIDs, strconv.FormatInt(tag.ID, 10))
		}

		pathIDs = append(pathIDs, strconv.FormatInt(tag.ID, 10))

		repositoryUpdaterConcrete.Path.Push(model.UpdateOperation[string]{Operator: domainUpdater.ParentPath.Wrappee.Operator, Operand: strings.Join(pathIDs, ";")})
	}

	//***********************    Set Children    ***********************//
	if domainUpdater.Subtags.HasValue {
		pathIDs := make([]string, 0, len(domainUpdater.Subtags.Wrappee.Operand)+1)
		for _, tag := range domainUpdater.Subtags.Wrappee.Operand {
			pathIDs = append(pathIDs, strconv.FormatInt(tag.ID, 10))
		}

		repositoryUpdaterConcrete.Children.Push(model.UpdateOperation[string]{Operator: domainUpdater.Subtags.Wrappee.Operator, Operand: strings.Join(pathIDs, ";")})
	}

	//**************************    Set ID    **************************//
	if domainUpdater.ID.HasValue {
		repositoryUpdaterConcrete.ID.Push(model.UpdateOperation[int64]{Operator: domainUpdater.ID.Wrappee.Operator, Operand: repositoryUpdaterConcrete.ID.Wrappee.Operand})
	}

    repositoryUpdater = repositoryUpdaterConcrete

	return
}

