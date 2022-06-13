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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/{{.EntityName}}_repository.go.tpl

package repository

import (
    "database/sql"
	repositoryCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/optional.go"
    "context"
	"fmt"
    "github.com/volatiletech/sqlboiler/v4/boil"
    "github.com/volatiletech/sqlboiler/v4/queries/qm"
    "github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/null/v8"
	"container/list"
    {{ if ne .DatabaseName "sqlite3" }}
    "time"
    {{end}}
)

{{template "structDefinition" .}}
{{template "repositoryHelperTypes" .}}

func GetDocumentDomainToSqlRepositoryModel(ctx context.Context, db *sql.DB) func(domainModel *domain.Document) (sqlRepositoryModel *Document, err error) {
    return func(domainModel *domain.Document) (sqlRepositoryModel *Document, err error) {
        return DocumentDomainToSqlRepositoryModel(ctx, db, domainModel)
    }
}

func GetDocumentSqlRepositoryToDomainModel(ctx context.Context, db *sql.DB) func(repositoryModel *Document) (domainModel *domain.Document, err error) {
    return func(sqlRepositoryModel *Document) (domainModel *domain.Document, err error) {
        return DocumentSqlRepositoryToDomainModel(ctx, db,sqlRepositoryModel)
    }
}

type {{$StructName}}ConstructorArgs struct {
    DB *sql.DB
}

func (repo *{{$StructName}}) New(args any) (repositoryCommon.DocumentRepository, error) {
    constructorArgs, ok := args.({{$StructName}}ConstructorArgs)
    if !ok {
        return repo, fmt.Errorf("expected type %T but got %T", {{$StructName}}ConstructorArgs{}, args)
    }

    repo.db = constructorArgs.DB

    return repo, nil
}

func (repo *{{$StructName}}) Add(ctx context.Context, domainModels []*domain.Document) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetDocumentDomainToSqlRepositoryModel(ctx, repo.db))
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

func (repo *{{$StructName}}) Replace(ctx context.Context, domainModels []*domain.Document) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetDocumentDomainToSqlRepositoryModel(ctx, repo.db))
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
func (repo *{{$StructName}}) Upsert(ctx context.Context, domainModels []*domain.Document) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetDocumentDomainToSqlRepositoryModel(ctx, repo.db))
	if err != nil {
		return err
	}

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, repositoryModel := range repositoryModels {
        {{ if eq .DatabaseName "mssql" }}
		err = repositoryModel.Upsert(ctx, tx, boil.Infer(), boil.Infer())
        {{else}}
		err = repositoryModel.Upsert(ctx, tx, false, []string{}, boil.Infer(), boil.Infer())
        {{end}}
		if err != nil {
			return err
		}
	}

	tx.Commit()

    return nil
}

func (repo *{{$StructName}}) Update(ctx context.Context, domainModels []*domain.Document, domainColumnUpdater *domain.DocumentUpdater) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetDocumentDomainToSqlRepositoryModel(ctx, repo.db))
	if err != nil {
		return err
	}

    columnUpdater, err := DocumentDomainToSqlRepositoryUpdater(ctx, repo.db, domainColumnUpdater)
    if err != nil {
        return err
    }

   	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

    for _, model := range   repositoryModels {
        columnUpdater.ApplyToModel(model)
        model.Update(ctx, tx, boil.Infer())
    }

    tx.Commit()

    return err
}

func (repo *{{$StructName}}) UpdateWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter, domainColumnUpdater *domain.DocumentUpdater) (numAffectedRecords int64, err error) {
	var modelsToUpdate DocumentSlice

    columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
    if err != nil {
        return
    }

    columnUpdater, err := DocumentDomainToSqlRepositoryUpdater(ctx, repo.db, domainColumnUpdater)
    if err != nil {
        return
    }



    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

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

func (repo *{{$StructName}}) Delete(ctx context.Context, domainModels []*domain.Document) error {
    repositoryModels, err := goaoi.TransformCopySlice(domainModels, GetDocumentDomainToSqlRepositoryModel(ctx, repo.db))
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

func (repo *{{$StructName}}) DeleteWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (numAffectedRecords int64, err error) {
    columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
    if err != nil {
        return
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

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

func (repo *{{$StructName}}) CountWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (int64, error) {
    columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
    if err != nil {
        return 0, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

	return Documents(queryFilters...).Count(ctx, repo.db)
}

func (repo *{{$StructName}}) CountAll(ctx context.Context) (int64, error) {
	return Documents().Count(ctx, repo.db)
}

func (repo *{{$StructName}}) DoesExist(ctx context.Context, domainModel *domain.Document) (bool, error) {
    repositoryModel, err := DocumentDomainToSqlRepositoryModel(ctx, repo.db, domainModel)
    if err != nil {
        return false, err
    }

	return DocumentExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *{{$StructName}}) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (bool, error) {
    columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
    if err != nil {
        return false, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

	return Documents(queryFilters...).Exists(ctx, repo.db)
}

func (repo *{{$StructName}}) GetWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) ([]*domain.Document, error) {
    columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
    if err != nil {
        return []*domain.Document{}, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

    repositoryModels, err := Documents(queryFilters...).All(ctx, repo.db)

    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetDocumentSqlRepositoryToDomainModel(ctx, repo.db))

    return domainModels, err
}

func (repo *{{$StructName}}) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.DocumentFilter) (*domain.Document, error) {
    columnFilter, err := DocumentDomainToSqlRepositoryFilter(ctx, repo.db, domainColumnFilter)
    if err != nil {
        return nil, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

    repositoryModel, err := Documents(queryFilters...).One(ctx, repo.db)

    var domainModel *domain.Document
    if err != nil {
        return domainModel, err
    }

    domainModel, err = DocumentSqlRepositoryToDomainModel(ctx, repo.db, repositoryModel)

    return domainModel, err
}

func (repo *{{$StructName}}) GetAll(ctx context.Context) ([]*domain.Document, error) {
    repositoryModels, err := Documents().All(ctx, repo.db)

    if err != nil {
        return []*domain.Document{}, err
    }

    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetDocumentSqlRepositoryToDomainModel(ctx, repo.db))

    return domainModels, err
}

func (repo *{{$StructName}}) AddType(ctx context.Context, types []string) error {
    for _, type_ := range types {
        repositoryModel := DocumentType{DocumentType: type_}

        err := repositoryModel.Insert(ctx, repo.db, boil.Infer())
        if err != nil {
            return err
        }
    }

    return nil
}

func (repo *{{$StructName}}) DeleteType(ctx context.Context, types []string) error {
    _, err := DocumentTypes(BookmarkTypeWhere.Type.IN(types)).DeleteAll(ctx, repo.db)

	return err
}

func (repo *{{$StructName}}) UpdateType(ctx context.Context, oldType string, newType string) error {
    repositoryModel, err := DocumentTypes(DocumentTypeWhere.DocumentType.EQ(oldType)).One(ctx, repo.db)
    if err != nil {
        return err
    }

    repositoryModel.DocumentType = newType

    _, err = repositoryModel.Update(ctx, repo.db, boil.Infer())

    return err
}
