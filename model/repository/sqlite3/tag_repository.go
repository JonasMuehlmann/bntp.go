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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/Tag_repository.go.tpl

package repository

import (
    "database/sql"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/optional.go"
    "context"
	"fmt"
    "github.com/volatiletech/sqlboiler/v4/boil"
    "github.com/volatiletech/sqlboiler/v4/queries/qm"
    "github.com/volatiletech/sqlboiler/v4/queries"
	"container/list"
)

type Sqlite3TagRepository struct {
    db *sql.DB
}
type TagField string

var TagFields = struct {
    Tag  TagField
    ID  TagField
    
}{
    Tag: "tag",
    ID: "id",
    
}

var TagFieldsList = []TagField{
    TagField("Tag"),
    TagField("ID"),
    
}

var TagRelationsList = []string{
    "Bookmarks",
    "ParentTagTags",
    "ChildTagTags",
    "Documents",
    "ParentTagTagParentPaths",
    "TagParentPaths",
    
}

type TagFilter struct {
    Tag optional.Optional[model.FilterOperation[string]]
    ID optional.Optional[model.FilterOperation[int64]]
    
    Bookmarks optional.Optional[model.FilterOperation[*Bookmark]]
    ParentTagTags optional.Optional[model.FilterOperation[*Tag]]
    ChildTagTags optional.Optional[model.FilterOperation[*Tag]]
    Documents optional.Optional[model.FilterOperation[*Document]]
    ParentTagTagParentPaths optional.Optional[model.FilterOperation[TagParentPathSlice]]
    TagParentPaths optional.Optional[model.FilterOperation[TagParentPathSlice]]
    
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
    if filter.ID.HasValue {
    setFilters.PushBack(TagFilterMapping[int64]{Field: TagFields.ID, FilterOperation: filter.ID.Wrappee})
    }
    

    return setFilters
}

type TagUpdater struct {
    Tag optional.Optional[model.UpdateOperation[string]]
    ID optional.Optional[model.UpdateOperation[int64]]
    
    Bookmarks optional.Optional[model.UpdateOperation[BookmarkSlice]]
    ParentTagTags optional.Optional[model.UpdateOperation[TagSlice]]
    ChildTagTags optional.Optional[model.UpdateOperation[TagSlice]]
    Documents optional.Optional[model.UpdateOperation[DocumentSlice]]
    ParentTagTagParentPaths optional.Optional[model.UpdateOperation[TagParentPathSlice]]
    TagParentPaths optional.Optional[model.UpdateOperation[TagParentPathSlice]]
    
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
    if updater.ID.HasValue {
    setUpdaters.PushBack(TagUpdaterMapping[int64]{Field: TagFields.ID, Updater: updater.ID.Wrappee})
    }
    

    return setUpdaters
}

func (updater *TagUpdater) ApplyToModel(tagModel *Tag) {
    if updater.Tag.HasValue {
        model.ApplyUpdater(&(*tagModel).Tag, updater.Tag.Wrappee)
    }
    if updater.ID.HasValue {
        model.ApplyUpdater(&(*tagModel).ID, updater.ID.Wrappee)
    }
    
}

type Sqlite3TagRepositoryHook func(context.Context, Sqlite3TagRepository) error

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
            panic("Expected a scalar operand for FilterEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" = ?", filterOperand.Operand))
    case model.FilterNEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterNEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" != ?", filterOperand.Operand))
    case model.FilterGreaterThan:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterGreaterThan operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" > ?", filterOperand.Operand))
    case model.FilterGreaterThanEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterGreaterThanEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" >= ?", filterOperand.Operand))
    case model.FilterLessThan:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterLessThan operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" < ?", filterOperand.Operand))
    case model.FilterLessThanEqual:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterLessThanEqual operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" <= ?", filterOperand.Operand))
    case model.FilterIn:
        filterOperand, ok := filterOperation.Operand.(model.ListOperand[any])
        if !ok {
            panic("Expected a list operand for FilterIn operator")
        }

        newQueryMod = append(newQueryMod, qm.WhereIn(string(filterField)+" IN (?)", filterOperand.Operands))
    case model.FilterNotIn:
        filterOperand, ok := filterOperation.Operand.(model.ListOperand[any])
        if !ok {
            panic("Expected a list operand for FilterNotIn operator")
        }

        newQueryMod = append(newQueryMod, qm.WhereNotIn(string(filterField)+" IN (?)", filterOperand.Operands))
    case model.FilterBetween:
        filterOperand, ok := filterOperation.Operand.(model.RangeOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterBetween operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
    case model.FilterNotBetween:
        filterOperand, ok := filterOperation.Operand.(model.RangeOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterNotBetween operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" NOT BETWEEN ? AND ?", filterOperand.Start, filterOperand.End))
    case model.FilterLike:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterLike operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" LIKE ?", filterOperand.Operand))
    case model.FilterNotLike:
        filterOperand, ok := filterOperation.Operand.(model.ScalarOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterLike operator")
        }

        newQueryMod = append(newQueryMod, qm.Where(string(filterField)+" NOT LIKE ?", filterOperand.Operand))
    case model.FilterOr:
        filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterOr operator")
        }
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Or2(qm.Expr(buildQueryModFilterTag(filterField, filterOperand.RHS))))
    case model.FilterAnd:
        filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterAnd operator")
        }

        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterTag(filterField, filterOperand.RHS)))
    default:
        panic("Unhandled FilterOperator")
    }

    return newQueryMod
}

func buildQueryModListFromFilterTag(setFilters list.List) queryModSliceTag {
	queryModList := make(queryModSliceTag, 0, 2)

	for filter := setFilters.Front(); filter != nil; filter = filter.Next() {
		filterMapping, ok := filter.Value.(TagFilterMapping[any])
		if !ok {
			panic(fmt.Sprintf("Expected type %T but got %T", TagFilterMapping[any]{}, filter))
		}

        newQueryMod := buildQueryModFilterTag(filterMapping.Field, filterMapping.FilterOperation)

        queryModList = append(queryModList, newQueryMod...)
	}

	return queryModList
}

func GetTagDomainToSqlRepositoryModel(db *sql.DB) func(domainModel *domain.Tag) (sqlRepositoryModel *Tag, err error) {
    return func(domainModel *domain.Tag) (sqlRepositoryModel *Tag, err error) {
        return TagDomainToSqlRepositoryModel(db, domainModel)
    }
}

func GetTagSqlRepositoryToDomainModel(db *sql.DB) func(repositoryModel *Tag) (domainModel *domain.Tag, err error) {
    return func(sqlRepositoryModel *Tag) (domainModel *domain.Tag, err error) {
        return TagSqlRepositoryToDomainModel(db,sqlRepositoryModel)
    }
}

func (repo * Sqlite3TagRepository) New(args ...any) (Sqlite3TagRepository, error) {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3TagRepository) Add(ctx context.Context, domainModels []*domain.Tag) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetTagDomainToSqlRepositoryModel(repo.db))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		err = repositoryModel.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	tx.Commit()

    return nil
}

func (repo *Sqlite3TagRepository) Replace(ctx context.Context, domainModels []*domain.Tag) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetTagDomainToSqlRepositoryModel(repo.db))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		_, err = repositoryModel.Update(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}
	}

	tx.Commit()

    return nil
}

func (repo *Sqlite3TagRepository) UpdateWhere(ctx context.Context, domainColumnFilter *domain.TagFilter, domainColumnUpdater *domain.TagUpdater) (numAffectedRecords int64, err error) {
    // NOTE: This kind of update is inefficient, since we do a read just to do a write later, but at the moment there is no better way
    // Either SQLboiler adds support for this usecase or (preferably), we use the caching and hook system to avoid database interaction, when it is not needed

	var modelsToUpdate TagSlice

    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return
    }

    columnUpdater, err := TagDomainToSqlRepositoryUpdater(repo.db, domainColumnUpdater)
    if err != nil {
        return
    }



    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	modelsToUpdate, err = Tags(queryFilters...).All(ctx, repo.db)
	if err != nil {
		return
	}

    numAffectedRecords = int64(len(modelsToUpdate))

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

    for _, model := range modelsToUpdate {
        columnUpdater.ApplyToModel(model)
        model.Update(ctx, tx, boil.Infer())
    }

    tx.Commit()

    return
}

func (repo *Sqlite3TagRepository) Delete(ctx context.Context, domainModels []*domain.Tag) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetTagDomainToSqlRepositoryModel(repo.db))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
		_, err = repositoryModel.Delete(ctx, tx)
		if err != nil {
			return err
		}
	}

	tx.Commit()

    return nil
}

func (repo *Sqlite3TagRepository) DeleteWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (numAffectedRecords int64, err error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Tags(queryFilters...).DeleteAll(ctx, tx)

    tx.Commit()

    return
}

func (repo *Sqlite3TagRepository) CountWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (int64, error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return 0, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	return Tags(queryFilters...).Count(ctx, repo.db)
}

func (repo *Sqlite3TagRepository) CountAll(ctx context.Context) (int64, error) {
	return Tags().Count(ctx, repo.db)
}

func (repo *Sqlite3TagRepository) DoesExist(ctx context.Context, domainModel *domain.Tag) (bool, error) {
    repositoryModel, err := TagDomainToSqlRepositoryModel(repo.db, domainModel)
    if err != nil {
        return false, err
    }

	return TagExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *Sqlite3TagRepository) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (bool, error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return false, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

	return Tags(queryFilters...).Exists(ctx, repo.db)
}

func (repo *Sqlite3TagRepository) GetWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) ([]*domain.Tag, error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return []*domain.Tag{}, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

    repositoryModels, err := Tags(queryFilters...).All(ctx, repo.db)
    if err != nil {
        return []*domain.Tag{}, err
    }


    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetTagSqlRepositoryToDomainModel(repo.db))

    return domainModels, err
}

func (repo *Sqlite3TagRepository) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (*domain.Tag, error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return nil, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterTag(setFilters)

    repositoryModel, err := Tags(queryFilters...).One(ctx, repo.db)
    if err != nil {
        return nil, err
    }

    return TagSqlRepositoryToDomainModel(repo.db, repositoryModel)
}

func (repo *Sqlite3TagRepository) GetAll(ctx context.Context) ([]*domain.Tag, error) {
    repositoryModels, err := Tags().All(ctx, repo.db)
    if err != nil {
        return []*domain.Tag{}, err
    }

    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetTagSqlRepositoryToDomainModel(repo.db))

    return domainModels, err
}
