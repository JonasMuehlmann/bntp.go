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

package libdocuments_test

import (
	"context"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/bntp"
	"github.com/JonasMuehlmann/bntp.go/bntp/libdocuments"
	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	fsRepo "github.com/JonasMuehlmann/bntp.go/model/repository/fs"
	sqlite3Repo "github.com/JonasMuehlmann/bntp.go/model/repository/sqlite3"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/barweiss/go-tuple"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

// FIX: This tests the repository
// func TestDocumentContentManagerAdd(t *testing.T) {
// 	tests := []struct {
// 		err          error
// 		errorMatcher testCommon.OutputValidator
// 		name         string
// 		pathContents []tuple.T2[string, string]
// 	}{
// 		{
// 			name:         "empty args",
// 			pathContents: nil,
// 			err:          helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name:         "zero value args",
// 			pathContents: []tuple.T2[string, string]{},
// 			err:          helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name:         "good minimal args",
// 			pathContents: []tuple.T2[string, string]{{V1: "path1", V2: "content1"}, {V1: "path2", V2: "content2"}},
// 		},
// 		{
// 			name:         "add duplicate",
// 			pathContents: []tuple.T2[string, string]{{V1: "path1", V2: "content1"}, {V1: "path1", V2: "content1"}},
// 			err:          helper.DuplicateInsertionError{},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			defer testCommon.HandlePanic(t, test.name)

// 			fs := afero.NewMemMapFs()

// 			repoConcrete := &fsRepo.FSDocumentContentRepository{}
// 			repoAbstract, err := repoConcrete.New(fsRepo.FSDocumentContentRepositoryConstructorArgs{Fs: fs, Logger: logrus.StandardLogger()})
// 			assert.NoError(t, err, test.name+", assert repository creation")

// 			repoConcrete = repoAbstract.(*fsRepo.FSDocumentContentRepository)

// 			manager, err := libdocuments.NewDocumentContentManager(repoConcrete.Logger, &bntp.Hooks[string]{}, repoConcrete)
// 			assert.NoError(t, err, test.name+", assert manager creation")

// 			err = manager.Add(context.Background(), test.pathContents)

// 			if test.err != nil {
// 				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
// 			} else if test.errorMatcher != nil {
// 				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
// 			} else {
// 				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
// 			}
// 		})
// 	}
// }

func TestDocumentContentManagerUpdateDocumentContentsFromNewModels(t *testing.T) {
	tests := []struct {
		err                     error
		errorMatcher            testCommon.OutputValidator
		name                    string
		oldDocuments            []*domain.Document
		oldDocumentPathContents []tuple.T2[string, string]
		addedDocuments          []*domain.Document
		newDocumentPathContents []tuple.T2[string, string]
		addNew                  bool
		tags                    []*domain.Tag
	}{
		{
			name: "empty args",
			err:  helper.IneffectiveOperationError{},
		},
		{
			name: "zero value args",
			err:  helper.IneffectiveOperationError{},
		},
		{
			name:           "no change",
			oldDocuments:   []*domain.Document{{ID: 1, Path: "Foo"}, {ID: 2, Path: "Bar"}},
			addedDocuments: []*domain.Document{{ID: 1, Path: "Foo"}, {ID: 2, Path: "Bar"}},
			addNew:         true,
		},
		{
			name:           "no old documents",
			addedDocuments: []*domain.Document{{ID: 1, Path: "Foo"}, {ID: 2, Path: "Bar"}},
			addNew:         true,
		},
		{
			name:         "add to all",
			oldDocuments: []*domain.Document{{ID: 1, Path: "Foo"}, {ID: 2, Path: "Bar"}},
			oldDocumentPathContents: []tuple.T2[string, string]{
				{V1: "Foo", V2: "# Tags\n# Links\n# Backlinks"},
				{V1: "Bar", V2: "# Tags\n# Links\n# Backlinks"},
			},
			tags: []*domain.Tag{{Tag: "tag1", ID: 1}},
			addedDocuments: []*domain.Document{
				{
					ID:                     1,
					Path:                   "Foo",
					TagIDs:                 []int64{1},
					LinkedDocumentIDs:      []int64{2},
					BacklinkedDocumentsIDs: []int64{2},
				}, {
					ID:                     2,
					Path:                   "Bar",
					TagIDs:                 []int64{1},
					LinkedDocumentIDs:      []int64{1},
					BacklinkedDocumentsIDs: []int64{1},
				},
			},
			newDocumentPathContents: []tuple.T2[string, string]{
				{V1: "Foo", V2: "# Tags\ntag1\n# Links\n- (Bar)[Bar]\n# Backlinks\n- (Bar)[Bar]\n"},
				{V1: "Bar", V2: "# Tags\ntag1\n# Links\n- (Foo)[Foo]\n# Backlinks\n- (Foo)[Foo]\n"},
			},
			addNew: true,
		},
		{
			name:         "remove from all",
			oldDocuments: []*domain.Document{{ID: 1, Path: "Foo"}, {ID: 2, Path: "Bar"}},
			newDocumentPathContents: []tuple.T2[string, string]{
				{V1: "Foo", V2: "# Tags\ntag1\n# Links\n- (Bar)[Bar]\n# Backlinks\n- (Bar)[Bar]\n"},
				{V1: "Bar", V2: "# Tags\ntag1\n# Links\n- (Foo)[Foo]\n# Backlinks\n- (Foo)[Foo]\n"},
			},
			tags: []*domain.Tag{{Tag: "tag1", ID: 1}},
			addedDocuments: []*domain.Document{
				{
					ID:                     1,
					Path:                   "Foo",
					TagIDs:                 []int64{1},
					LinkedDocumentIDs:      []int64{2},
					BacklinkedDocumentsIDs: []int64{2},
				}, {
					ID:                     2,
					Path:                   "Bar",
					TagIDs:                 []int64{1},
					LinkedDocumentIDs:      []int64{1},
					BacklinkedDocumentsIDs: []int64{1},
				},
			},
			oldDocumentPathContents: []tuple.T2[string, string]{
				{V1: "Foo", V2: "# Tags\n# Links\n# Backlinks"},
				{V1: "Bar", V2: "# Tags\n# Links\n# Backlinks"},
			},
			addNew: true,
		},

		{
			name:           "added documents don't exist",
			addedDocuments: []*domain.Document{{ID: 1, Path: "Foo"}, {ID: 2, Path: "Bar"}},
			err:            helper.NonExistentPrimaryDataError{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			//*****************    Setup FS and create files    ****************//
			fs := afero.NewMemMapFs()

			if test.oldDocumentPathContents != nil {
				pathContentWriter := func(pathContent tuple.T2[string, string]) error {
					return afero.WriteFile(fs, pathContent.V1, []byte(pathContent.V2), 0o644)
				}

				err := goaoi.ForeachSlice(test.oldDocumentPathContents, pathContentWriter)
				assert.NoError(t, err, test.name+", assert file creation")
			}

			//**************    Setup document content manager    **************//
			repoConcrete := &fsRepo.FSDocumentContentRepository{}
			repoAbstract, err := repoConcrete.New(fsRepo.FSDocumentContentRepositoryConstructorArgs{Fs: fs, Logger: logrus.StandardLogger()})
			assert.NoError(t, err, test.name+", assert document content repository creation")

			repoConcrete = repoAbstract.(*fsRepo.FSDocumentContentRepository)

			manager, err := libdocuments.NewDocumentContentManager(repoConcrete.Logger, &bntp.Hooks[string]{}, repoConcrete)
			assert.NoError(t, err, test.name+", assert document content manager creation")

			//*******************    Setup tag repository    *******************//
			tagRepoConcrete := &sqlite3Repo.Sqlite3TagRepository{}
			tagRepoAbstract, err := tagRepoConcrete.New(sqlite3Repo.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: logrus.StandardLogger()})
			assert.NoError(t, err, test.name+", assert tag repository creation")

			tagRepoConcrete = tagRepoAbstract.(*sqlite3Repo.Sqlite3TagRepository)

			documentRepoConcrete := &sqlite3Repo.Sqlite3DocumentRepository{}

			//******************    Setup document manager    ******************//
			documentRepoAbstract, err := documentRepoConcrete.New(sqlite3Repo.Sqlite3DocumentRepositoryConstructorArgs{DB: db, TagRepository: tagRepoConcrete, Logger: repoConcrete.Logger})
			assert.NoError(t, err, test.name+", assert document repository creation")

			documentRepoConcrete = documentRepoAbstract.(*sqlite3Repo.Sqlite3DocumentRepository)

			documentManager, err := libdocuments.NewDocumentManager(repoConcrete.Logger, &bntp.Hooks[domain.Document]{}, documentRepoConcrete)
			assert.NoError(t, err, test.name+", assert document manager creation")

			//******************    Add tags and documents    ******************//
			if test.tags != nil {
				err = documentManager.Repository.GetTagRepository().Add(context.Background(), test.tags)
				assert.NoError(t, err, test.name+", assert adding tags")
			}

			if test.addNew {
				err = documentManager.Add(context.Background(), test.addedDocuments)
				assert.NoError(t, err, test.name+", assert creating added documents")
			}

			//*********************    Run main function    ********************//
			err = manager.UpdateDocumentContentsFromNewModels(context.Background(), test.addedDocuments, &documentManager)

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}

			if test.err == nil && !assert.ObjectsAreEqualValues(test.oldDocuments, test.addedDocuments) && test.oldDocuments != nil {
				newPathsContentSOA := bntp.TupleToSOA2(test.oldDocumentPathContents)

				newDocumentContents, err := manager.Get(context.Background(), newPathsContentSOA.V1)
				assert.NoError(t, err, test.name+", assert getting new document contents")
				assert.Equal(t, newPathsContentSOA.V2, newDocumentContents, test.name+", assert new document contents match")
			}
		})
	}
}
