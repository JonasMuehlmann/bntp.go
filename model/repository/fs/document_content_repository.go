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

	commonRepo "github.com/JonasMuehlmann/bntp.go/model/repository"
	"github.com/spf13/afero"
)

type FSDocumentContentRepository struct {
	fs afero.Fs
}

func (repo *FSDocumentContentRepository) New(args any) (commonRepo.DocumentContentRepository, error) {
	fs, ok := args.(afero.Fs)
	if !ok {
		return repo, fmt.Errorf("expected type %T but got %T", repo.fs, args)
	}

	repo.fs = fs

	return repo, nil
}

func (repo *FSDocumentContentRepository) Add(ctx context.Context, path string, content string) error {
	file, err := repo.fs.Create(path)
	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func (repo *FSDocumentContentRepository) Update(ctx context.Context, path string, newContent string) error {
	file, err := repo.fs.Open(path)
	if err != nil {
		return err
	}

	_, err = file.WriteString(newContent)
	if err != nil {
		return err
	}

	return nil
}

func (repo *FSDocumentContentRepository) Move(ctx context.Context, oldPath string, newPath string) error {
	return repo.fs.Rename(oldPath, newPath)
}

func (repo *FSDocumentContentRepository) Delete(ctx context.Context, path string) error {
	return repo.fs.Remove(path)
}

func (repo *FSDocumentContentRepository) Get(ctx context.Context, path string) (content string, err error) {
	file, err := repo.fs.Open(path)
	if err != nil {
		return "", err
	}

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 0, stat.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
