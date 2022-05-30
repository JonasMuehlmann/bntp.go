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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/Document_repository.go.tpl

package repository

import (
    "database/sql"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/optional.go"
    "context"
	"fmt"
    "github.com/volatiletech/sqlboiler/v4/boil"
    "github.com/volatiletech/sqlboiler/v4/queries/qm"
    "github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/null/v8"
	"container/list"
)

type Sqlite3DocumentRepository struct {
    db *sql.DB
}
type DocumentField string

var DocumentFields = struct {
    CreatedAt  DocumentField
    UpdatedAt  DocumentField
    Path  DocumentField
    DeletedAt  DocumentField
    DocumentTypeID  DocumentField
    ID  DocumentField
    
}{
    CreatedAt: "created_at",
    UpdatedAt: "updated_at",
    Path: "path",
    DeletedAt: "deleted_at",
    DocumentTypeID: "document_type_id",
    ID: "id",
    
}

var DocumentFieldsList = []DocumentField{
    DocumentField("CreatedAt"),
    DocumentField("UpdatedAt"),
    DocumentField("Path"),
    DocumentField("DeletedAt"),
    DocumentField("DocumentTypeID"),
    DocumentField("ID"),
    
}

var DocumentRelationsList = []string{
    "DocumentType",
    "Tags",
    "SourceDocuments",
    "DestinationDocuments",
    
}

type DocumentFilter struct {
    CreatedAt optional.Optional[model.FilterOperation[string]]
    UpdatedAt optional.Optional[model.FilterOperation[string]]
    Path optional.Optional[model.FilterOperation[string]]
    DeletedAt optional.Optional[model.FilterOperation[null.String]]
    DocumentTypeID optional.Optional[model.FilterOperation[null.Int64]]
    ID optional.Optional[model.FilterOperation[int64]]
    
    DocumentType optional.Optional[model.FilterOperation[*DocumentType]]
    Tags optional.Optional[model.FilterOperation[*Tag]]
    SourceDocuments optional.Optional[model.FilterOperation[*Document]]
    DestinationDocuments optional.Optional[model.FilterOperation[*Document]]
    
}

type DocumentFilterMapping[T any] struct {
    Field DocumentField
    FilterOperation model.FilterOperation[T]
}

func (filter *DocumentFilter) GetSetFilters() *list.List {
    setFilters := list.New()

    if filter.CreatedAt.HasValue {
    setFilters.PushBack(DocumentFilterMapping[string]{Field: DocumentFields.CreatedAt, FilterOperation: filter.CreatedAt.Wrappee})
    }
    if filter.UpdatedAt.HasValue {
    setFilters.PushBack(DocumentFilterMapping[string]{Field: DocumentFields.UpdatedAt, FilterOperation: filter.UpdatedAt.Wrappee})
    }
    if filter.Path.HasValue {
    setFilters.PushBack(DocumentFilterMapping[string]{Field: DocumentFields.Path, FilterOperation: filter.Path.Wrappee})
    }
    if filter.DeletedAt.HasValue {
    setFilters.PushBack(DocumentFilterMapping[null.String]{Field: DocumentFields.DeletedAt, FilterOperation: filter.DeletedAt.Wrappee})
    }
    if filter.DocumentTypeID.HasValue {
    setFilters.PushBack(DocumentFilterMapping[null.Int64]{Field: DocumentFields.DocumentTypeID, FilterOperation: filter.DocumentTypeID.Wrappee})
    }
    if filter.ID.HasValue {
    setFilters.PushBack(DocumentFilterMapping[int64]{Field: DocumentFields.ID, FilterOperation: filter.ID.Wrappee})
    }
    

    return setFilters
}

type DocumentUpdater struct {
    CreatedAt optional.Optional[model.UpdateOperation[string]]
    UpdatedAt optional.Optional[model.UpdateOperation[string]]
    Path optional.Optional[model.UpdateOperation[string]]
    DeletedAt optional.Optional[model.UpdateOperation[null.String]]
    DocumentTypeID optional.Optional[model.UpdateOperation[null.Int64]]
    ID optional.Optional[model.UpdateOperation[int64]]
    
    DocumentType optional.Optional[model.UpdateOperation[*DocumentType]]
    Tags optional.Optional[model.UpdateOperation[TagSlice]]
    SourceDocuments optional.Optional[model.UpdateOperation[DocumentSlice]]
    DestinationDocuments optional.Optional[model.UpdateOperation[DocumentSlice]]
    
}

type DocumentUpdaterMapping[T any] struct {
    Field DocumentField
    Updater model.UpdateOperation[T]
}

func (updater *DocumentUpdater) GetSetUpdaters() *list.List {
    setUpdaters := list.New()

    if updater.CreatedAt.HasValue {
    setUpdaters.PushBack(DocumentUpdaterMapping[string]{Field: DocumentFields.CreatedAt, Updater: updater.CreatedAt.Wrappee})
    }
    if updater.UpdatedAt.HasValue {
    setUpdaters.PushBack(DocumentUpdaterMapping[string]{Field: DocumentFields.UpdatedAt, Updater: updater.UpdatedAt.Wrappee})
    }
    if updater.Path.HasValue {
    setUpdaters.PushBack(DocumentUpdaterMapping[string]{Field: DocumentFields.Path, Updater: updater.Path.Wrappee})
    }
    if updater.DeletedAt.HasValue {
    setUpdaters.PushBack(DocumentUpdaterMapping[null.String]{Field: DocumentFields.DeletedAt, Updater: updater.DeletedAt.Wrappee})
    }
    if updater.DocumentTypeID.HasValue {
    setUpdaters.PushBack(DocumentUpdaterMapping[null.Int64]{Field: DocumentFields.DocumentTypeID, Updater: updater.DocumentTypeID.Wrappee})
    }
    if updater.ID.HasValue {
    setUpdaters.PushBack(DocumentUpdaterMapping[int64]{Field: DocumentFields.ID, Updater: updater.ID.Wrappee})
    }
    

    return setUpdaters
}

func (updater *DocumentUpdater) ApplyToModel(documentModel *Document) {
    if updater.CreatedAt.HasValue {
        model.ApplyUpdater(&(*documentModel).CreatedAt, updater.CreatedAt.Wrappee)
    }
    if updater.UpdatedAt.HasValue {
        model.ApplyUpdater(&(*documentModel).UpdatedAt, updater.UpdatedAt.Wrappee)
    }
    if updater.Path.HasValue {
        model.ApplyUpdater(&(*documentModel).Path, updater.Path.Wrappee)
    }
    if updater.DeletedAt.HasValue {
        model.ApplyUpdater(&(*documentModel).DeletedAt, updater.DeletedAt.Wrappee)
    }
    if updater.DocumentTypeID.HasValue {
        model.ApplyUpdater(&(*documentModel).DocumentTypeID, updater.DocumentTypeID.Wrappee)
    }
    if updater.ID.HasValue {
        model.ApplyUpdater(&(*documentModel).ID, updater.ID.Wrappee)
    }
    
}

type Sqlite3DocumentRepositoryHook func(context.Context, Sqlite3DocumentRepository) error

type queryModSliceDocument []qm.QueryMod

func (s queryModSliceDocument) Apply(q *queries.Query) {
    qm.Apply(q, s...)
}

func buildQueryModFilterDocument[T any](filterField DocumentField, filterOperation model.FilterOperation[T]) queryModSliceDocument {
    var newQueryMod queryModSliceDocument

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
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterDocument(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Or2(qm.Expr(buildQueryModFilterDocument(filterField, filterOperand.RHS))))
    case model.FilterAnd:
        filterOperand, ok := filterOperation.Operand.(model.CompoundOperand[any])
        if !ok {
            panic("Expected a scalar operand for FilterAnd operator")
        }

        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterDocument(filterField, filterOperand.LHS)))
        newQueryMod = append(newQueryMod, qm.Expr(buildQueryModFilterDocument(filterField, filterOperand.RHS)))
    default:
        panic("Unhandled FilterOperator")
    }

    return newQueryMod
}

func buildQueryModListFromFilterDocument(setFilters list.List) queryModSliceDocument {
	queryModList := make(queryModSliceDocument, 0, 6)

	for filter := setFilters.Front(); filter != nil; filter = filter.Next() {
		filterMapping, ok := filter.Value.(DocumentFilterMapping[any])
		if !ok {
			panic(fmt.Sprintf("Expected type %T but got %T", DocumentFilterMapping[any]{}, filter))
		}

        newQueryMod := buildQueryModFilterDocument(filterMapping.Field, filterMapping.FilterOperation)

        queryModList = append(queryModList, newQueryMod...)
	}

	return queryModList
}

func (repo * Sqlite3DocumentRepository) New(args ...any) (Sqlite3DocumentRepository, error) {
        panic("not implemented") // TODO: Implement
}

func (repo *Sqlite3DocumentRepository) Add(ctx context.Context, repositoryModels []Document) error {
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

func (repo *Sqlite3DocumentRepository) Replace(ctx context.Context, repositoryModels []Document) error {
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

func (repo *Sqlite3DocumentRepository) UpdateWhere(ctx context.Context, columnFilter DocumentFilter, columnUpdater DocumentUpdater) (numAffectedRecords int64, err error) {
    // NOTE: This kind of update is inefficient, since we do a read just to do a write later, but at the moment there is no better way
    // Either SQLboiler adds support for this usecase or (preferably), we use the caching and hook system to avoid database interaction, when it is not needed

	var modelsToUpdate DocumentSlice

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	modelsToUpdate, err = Documents(queryFilters...).All(ctx, repo.db)
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

func (repo *Sqlite3DocumentRepository) Delete(ctx context.Context, repositoryModels []Document) error {
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

func (repo *Sqlite3DocumentRepository) DeleteWhere(ctx context.Context, columnFilter DocumentFilter) (numAffectedRecords int64, err error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Documents(queryFilters...).DeleteAll(ctx, tx)
	if err != nil {
		return
	}

    tx.Commit()

    return
}

func (repo *Sqlite3DocumentRepository) CountWhere(ctx context.Context, columnFilter DocumentFilter) (int64, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	return Documents(queryFilters...).Count(ctx, repo.db)
}

func (repo *Sqlite3DocumentRepository) CountAll(ctx context.Context) (int64, error) {
	return Documents().Count(ctx, repo.db)
}

func (repo *Sqlite3DocumentRepository) DoesExist(ctx context.Context, repositoryModel Document) (bool, error) {
	return DocumentExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *Sqlite3DocumentRepository) DoesExistWhere(ctx context.Context, columnFilter DocumentFilter) (bool, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	return Documents(queryFilters...).Exists(ctx, repo.db)
}

func (repo *Sqlite3DocumentRepository) GetWhere(ctx context.Context, columnFilter DocumentFilter) ([]*Document, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	return Documents(queryFilters...).All(ctx, repo.db)
}

func (repo *Sqlite3DocumentRepository) GetFirstWhere(ctx context.Context, columnFilter DocumentFilter) (*Document, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilterDocument(setFilters)

	return Documents(queryFilters...).One(ctx, repo.db)
}

func (repo *Sqlite3DocumentRepository) GetAll(ctx context.Context) ([]*Document, error) {
	return Documents().All(ctx, repo.db)
}
