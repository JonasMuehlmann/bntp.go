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

{{template "structDefinition" .}}
{{template "repositoryHelperTypes" .}}

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

type {{$StructName}}ConstructorArgs struct {
    DB *sql.DB
}

func (repo *{{$StructName}}) New(args any) (*{{$StructName}}, error) {
    constructorArgs, ok := args.({{$StructName}}ConstructorArgs)
    if !ok {
        return fmt.Errorf("expected type %T but got %T", {{$StructName}}ConstructorArgs{}, args)
    }

    repo.db = constructorArgs.DB

    return repo, nil
}

func (repo *{{$StructName}}) Add(ctx context.Context, domainModels []*domain.Tag) error {
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

func (repo *{{$StructName}}) Replace(ctx context.Context, domainModels []*domain.Tag) error {
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

func (repo *{{$StructName}}) UpdateWhere(ctx context.Context, domainColumnFilter *domain.TagFilter, domainColumnUpdater *domain.TagUpdater) (numAffectedRecords int64, err error) {
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

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

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

func (repo *{{$StructName}}) Delete(ctx context.Context, domainModels []*domain.Tag) error {
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

func (repo *{{$StructName}}) DeleteWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (numAffectedRecords int64, err error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Tags(queryFilters...).DeleteAll(ctx, tx)

    tx.Commit()

    return
}

func (repo *{{$StructName}}) CountWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (int64, error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return 0, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

	return Tags(queryFilters...).Count(ctx, repo.db)
}

func (repo *{{$StructName}}) CountAll(ctx context.Context) (int64, error) {
	return Tags().Count(ctx, repo.db)
}

func (repo *{{$StructName}}) DoesExist(ctx context.Context, domainModel *domain.Tag) (bool, error) {
    repositoryModel, err := TagDomainToSqlRepositoryModel(repo.db, domainModel)
    if err != nil {
        return false, err
    }

	return TagExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *{{$StructName}}) DoesExistWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (bool, error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return false, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

	return Tags(queryFilters...).Exists(ctx, repo.db)
}

func (repo *{{$StructName}}) GetWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) ([]*domain.Tag, error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return []*domain.Tag{}, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

    repositoryModels, err := Tags(queryFilters...).All(ctx, repo.db)
    if err != nil {
        return []*domain.Tag{}, err
    }


    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetTagSqlRepositoryToDomainModel(repo.db))

    return domainModels, err
}

func (repo *{{$StructName}}) GetFirstWhere(ctx context.Context, domainColumnFilter *domain.TagFilter) (*domain.Tag, error) {
    columnFilter, err := TagDomainToSqlRepositoryFilter(repo.db, domainColumnFilter)
    if err != nil {
        return nil, err
    }

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := buildQueryModListFromFilter{{$EntityName}}(setFilters)

    repositoryModel, err := Tags(queryFilters...).One(ctx, repo.db)
    if err != nil {
        return nil, err
    }

    return TagSqlRepositoryToDomainModel(repo.db, repositoryModel)
}

func (repo *{{$StructName}}) GetAll(ctx context.Context) ([]*domain.Tag, error) {
    repositoryModels, err := Tags().All(ctx, repo.db)
    if err != nil {
        return []*domain.Tag{}, err
    }

    domainModels, err := goaoi.TransformCopySlice(repositoryModels, GetTagSqlRepositoryToDomainModel(repo.db))

    return domainModels, err
}
