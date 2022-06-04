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

	"github.com/barweiss/go-tuple"
)

type DocumentContent struct {
	Path    string
	Content string
}

type DocumentContentRepository interface {
	New(args any) (DocumentContentRepository, error)

	Add(ctx context.Context, pathContents []tuple.T2[string, string]) error
	Update(ctx context.Context, pathContents []tuple.T2[string, string]) error
	Move(ctx context.Context, pathChanges []tuple.T2[string, string]) error
	Delete(ctx context.Context, paths []string) error
	Get(ctx context.Context, paths []string) (contents []string, err error)
	// GetAll(context.Context) (records []DocumentContent, err error)
	// DoesExist(ctx context.Context, path string) (doesExist bool, err error)
	// CountAll(ctx context.Context) (numRecords int64, err error)
}
