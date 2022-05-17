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

// THIS CODE IS GENERATED BY GO GENERATE, IT'S TEMPLATE IS /templates/sql_repositories/{{LowercaseBeginning .EntityName}}_repository.go.tpl

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

{{template "structDefinition" .}}
{{template "repositoryHelperTypes" .}}

func (repo * {{$StructName}}) New(args ...any) ({{$StructName}}, error) {
        panic("not implemented") // TODO: Implement
}

func (repo *{{$StructName}}) Add(ctx context.Context, repositoryModels []Bookmark) error {
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

func (repo *{{$StructName}}) Replace(ctx context.Context, repositoryModels []Bookmark) error {
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

func (repo *{{$StructName}}) UpdateWhere(ctx context.Context, columnFilter BookmarkFilter, columnUpdater BookmarkUpdater) (numAffectedRecords int64, err error) {
    // NOTE: This kind of update is inefficient, since we do a read just to do a write later, but at the moment there is no better way
    // Either SQLboiler adds support for this usecase or (preferably), we use the caching and hook system to avoid database interaction, when it is not needed

    // TODO: Implement translator from domainColumnFilter to repositoryColumnFilter and updater
	var modelsToUpdate BookmarkSlice

    setFilters := *columnFilter.GetSetFilters()

	queryFilters := BuildQueryModListFromFilter(setFilters)

	modelsToUpdate, err = Bookmarks(queryFilters...).All(ctx, repo.db)

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

func (repo *{{$StructName}}) Delete(ctx context.Context, repositoryModels []Bookmark) error {
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

func (repo *{{$StructName}}) DeleteWhere(ctx context.Context, columnFilter BookmarkFilter) (numAffectedRecords int64, err error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := BuildQueryModListFromFilter(setFilters)

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	numAffectedRecords, err = Bookmarks(queryFilters...).DeleteAll(ctx, tx)

    tx.Commit()

    return
}

func (repo *{{$StructName}}) CountWhere(ctx context.Context, columnFilter BookmarkFilter) (int64, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := BuildQueryModListFromFilter(setFilters)

	return Bookmarks(queryFilters...).Count(ctx, repo.db)
}

func (repo *{{$StructName}}) CountAll(ctx context.Context) (int64, error) {
	return Bookmarks().Count(ctx, repo.db)
}

func (repo *{{$StructName}}) DoesExist(ctx context.Context, repositoryModel Bookmark) (bool, error) {
	return BookmarkExists(ctx, repo.db, repositoryModel.ID)
}

func (repo *{{$StructName}}) DoesExistWhere(ctx context.Context, columnFilter BookmarkFilter) (bool, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := BuildQueryModListFromFilter(setFilters)

	return Bookmarks(queryFilters...).Exists(ctx, repo.db)
}

func (repo *{{$StructName}}) GetWhere(ctx context.Context, columnFilter BookmarkFilter) ([]*Bookmark, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := BuildQueryModListFromFilter(setFilters)

	return Bookmarks(queryFilters...).All(ctx, repo.db)
}

func (repo *{{$StructName}}) GetFirstWhere(ctx context.Context, columnFilter BookmarkFilter) (*Bookmark, error) {
    setFilters := *columnFilter.GetSetFilters()

	queryFilters := BuildQueryModListFromFilter(setFilters)

	return Bookmarks(queryFilters...).One(ctx, repo.db)
}

func (repo *{{$StructName}}) GetAll(ctx context.Context) ([]*Bookmark, error) {
	return Bookmarks().All(ctx, repo.db)
}
