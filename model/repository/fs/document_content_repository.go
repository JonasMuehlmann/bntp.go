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

package repository

import (
	"context"
	"fmt"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	commonRepo "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/barweiss/go-tuple"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

type FSDocumentContentRepositoryConstructorArgs struct {
	Fs     afero.Fs
	Logger *log.Logger
}

type FSDocumentContentRepository struct {
	Logger *log.Logger
	fs     afero.Fs
}

func (repo *FSDocumentContentRepository) New(args any) (commonRepo.DocumentContentRepository, error) {
	constructorArgs, ok := args.(FSDocumentContentRepositoryConstructorArgs)
	if !ok {
		return repo, fmt.Errorf("expected type %T but got %T", repo.fs, args)
	}

	repo.fs = constructorArgs.Fs
	repo.Logger = constructorArgs.Logger

	return repo, nil
}

func (repo *FSDocumentContentRepository) Add(ctx context.Context, pathContents []tuple.T2[string, string]) error {
	if len(pathContents) == 0 {
		repo.Logger.Debug(helper.LogMessageEmptyInput)

		return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}

	err := goaoi.AnyOfSlice(pathContents, func(e tuple.T2[string, string]) bool { return e == tuple.T2[string, string]{} })
	if err == nil {
		err = helper.NilInputError{}
		repo.Logger.Error(err)

		return err
	}

	transformer := func(pathContent tuple.T2[string, string]) error {
		file, err := repo.fs.Create(pathContent.V1)
		if err != nil {
			return err
		}

		_, err = file.WriteString(pathContent.V2)
		if err != nil {
			return err
		}

		return nil
	}

	return goaoi.ForeachSlice(pathContents, transformer)
}

func (repo *FSDocumentContentRepository) Update(ctx context.Context, pathContents []tuple.T2[string, string]) error {
	if len(pathContents) == 0 {
		repo.Logger.Debug(helper.LogMessageEmptyInput)

		return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}

	err := goaoi.AnyOfSlice(pathContents, func(e tuple.T2[string, string]) bool { return e == tuple.T2[string, string]{} })
	if err == nil {
		err = helper.NilInputError{}
		repo.Logger.Error(err)

		return err
	}

	transformer := func(pathChange tuple.T2[string, string]) error {
		file, err := repo.fs.Open(pathChange.V1)
		if err != nil {
			return err
		}

		_, err = file.WriteString(pathChange.V2)
		if err != nil {
			return err
		}

		return err
	}

	return goaoi.ForeachSlice(pathContents, transformer)
}

func (repo *FSDocumentContentRepository) Move(ctx context.Context, pathChanges []tuple.T2[string, string]) error {
	if len(pathChanges) == 0 {
		repo.Logger.Debug(helper.LogMessageEmptyInput)

		return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}

	err := goaoi.AnyOfSlice(pathChanges, func(e tuple.T2[string, string]) bool { return e == tuple.T2[string, string]{} })
	if err == nil {
		err = helper.NilInputError{}
		repo.Logger.Error(err)

		return err
	}

	transformer := func(pathChange tuple.T2[string, string]) error {
		return repo.fs.Rename(pathChange.V1, pathChange.V2)
	}

	return goaoi.ForeachSlice(pathChanges, transformer)
}

func (repo *FSDocumentContentRepository) Delete(ctx context.Context, paths []string) error {
	if len(paths) == 0 {
		repo.Logger.Debug(helper.LogMessageEmptyInput)

		return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}

	return goaoi.ForeachSlice(paths, repo.fs.Remove)
}

func (repo *FSDocumentContentRepository) Get(ctx context.Context, paths []string) (contents []string, err error) {
	if len(paths) == 0 {
		repo.Logger.Debug(helper.EmptyInputError{})

		return nil, helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
	}

	transformer := func(path string) (content string, err error) {
		contentRaw, err := afero.ReadFile(repo.fs, path)
		if err == nil {
			content = string(contentRaw)
		}

		return
	}

	return goaoi.TransformCopySlice(paths, transformer)
}
