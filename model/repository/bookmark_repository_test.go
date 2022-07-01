package repository_test

import (
	"context"
	"testing"

	// repositoryCommon "github.com/JonasMuehlmann/bntp.go/model/repository".
	repository "github.com/JonasMuehlmann/bntp.go/model/repository/sqlite3"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
)

// func TestSQLBookmarkRepositoryAddTest(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		models []*domain.Bookmark
// 		err    error
// 	}{
// 		{
// 			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "One default-constructed input", models: []*domain.Bookmark{{}},
// 		},
// 		{
// 			name: "Two regular inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, models: []*domain.Bookmark{
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://example.com",
// 					Title:     optional.Make("My first bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://foo.example.com",
// 					Title:     optional.Make("My second bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two minimal inputs", models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			err = repo.Add(context.Background(), test.models)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryReplaceTest(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		models             []*domain.Bookmark
// 		err                error
// 		insertBeforeUpdate bool
// 	}{
// 		{
// 			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "One default-constructed input", models: []*domain.Bookmark{{}}, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, insertBeforeUpdate: true, models: []*domain.Bookmark{
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://example.com",
// 					Title:     optional.Make("My first bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://foo.example.com",
// 					Title:     optional.Make("My second bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two existing minimal inputs", insertBeforeUpdate: true, models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeUpdate {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			err = repo.Replace(context.Background(), test.models)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryUpsertTest(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		models             []*domain.Bookmark
// 		err                error
// 		insertBeforeUpdate bool
// 	}{
// 		{
// 			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "One default-constructed input", models: []*domain.Bookmark{{}},
// 		},
// 		{
// 			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, insertBeforeUpdate: true, models: []*domain.Bookmark{
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://example.com",
// 					Title:     optional.Make("My first bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://foo.example.com",
// 					Title:     optional.Make("My second bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two existing minimal inputs", insertBeforeUpdate: true, models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeUpdate {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			err = repo.Upsert(context.Background(), test.models)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryUpdateTest(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		models             []*domain.Bookmark
// 		updater            *domain.BookmarkUpdater
// 		err                error
// 		insertBeforeUpdate bool
// 	}{
// 		{
// 			name: "Empty input", models: []*domain.Bookmark{}, updater: &domain.BookmarkUpdater{}, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Nil input", models: nil, updater: &domain.BookmarkUpdater{}, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Nil updater", models: []*domain.Bookmark{{}}, updater: nil, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "Input containing nil value", models: []*domain.Bookmark{nil}, updater: &domain.BookmarkUpdater{}, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "One default-constructed input", models: []*domain.Bookmark{{}}, updater: &domain.BookmarkUpdater{}, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, updater: &domain.BookmarkUpdater{}, insertBeforeUpdate: true, models: []*domain.Bookmark{
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://example.com",
// 					Title:     optional.Make("My first bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://foo.example.com",
// 					Title:     optional.Make("My second bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two existing minimal inputs", insertBeforeUpdate: true, updater: &domain.BookmarkUpdater{}, models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two existing minimal inputs, overwrite IsCollection", insertBeforeUpdate: true,
// 			updater: &domain.BookmarkUpdater{
// 				IsCollection: optional.Make(model.UpdateOperation[bool]{Operator: model.UpdateSet, Operand: true}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeUpdate {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			err = repo.Update(context.Background(), test.models, test.updater)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryUpdateWhereTest(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		models             []*domain.Bookmark
// 		filter             *domain.BookmarkFilter
// 		updater            *domain.BookmarkUpdater
// 		err                error
// 		insertBeforeUpdate bool
// 		numAffectedRecords int64
// 	}{
// 		{
// 			name: "No entities", updater: &domain.BookmarkUpdater{}, filter: &domain.BookmarkFilter{}, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Nil updater", updater: nil, filter: &domain.BookmarkFilter{}, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "Nil filter", updater: &domain.BookmarkUpdater{}, filter: nil, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, updater: &domain.BookmarkUpdater{}, filter: &domain.BookmarkFilter{}, insertBeforeUpdate: true, models: []*domain.Bookmark{
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://example.com",
// 					Title:     optional.Make("My first bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 					DeletedAt: optional.Make(time.Now()),
// 					URL:       "https://foo.example.com",
// 					Title:     optional.Make("My second bookmark"),
// 					// These tags do not exist!
// 					Tags: []*domain.Tag{{
// 						Tag:        "Test",
// 						ParentPath: []*domain.Tag{},
// 						Subtags:    []*domain.Tag{},
// 						ID:         1,
// 					}},
// 					// This type does not exist
// 					BookmarkType: optional.Make("Text"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two existing minimal inputs, filter for title of first", numAffectedRecords: 1, insertBeforeUpdate: true, updater: &domain.BookmarkUpdater{},
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two existing minimal inputs, overwrite title", numAffectedRecords: 2, insertBeforeUpdate: true, filter: &domain.BookmarkFilter{},
// 			updater: &domain.BookmarkUpdater{
// 				Title: optional.Make(model.UpdateOperation[optional.Optional[string]]{Operator: model.UpdateSet, Operand: optional.Make("foo")}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeUpdate {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			numAffectedRecords, err := repo.UpdateWhere(context.Background(), test.filter, test.updater)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 			assert.Equalf(t, test.numAffectedRecords, numAffectedRecords, test.name)
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryDeleteTest(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		models             []*domain.Bookmark
// 		insertBeforeDelete bool
// 		err                error
// 	}{
// 		{
// 			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
// 		},
// 		{
// 			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "One default-constructed input", models: []*domain.Bookmark{{}},
// 		},
// 		{
// 			name: "Two minimal inputs", insertBeforeDelete: true, models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeDelete {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			err = repo.Delete(context.Background(), test.models)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryDeleteWhereTest(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		models             []*domain.Bookmark
// 		filter             *domain.BookmarkFilter
// 		err                error
// 		insertBeforeDelete bool
// 		numAffectedRecords int64
// 	}{
// 		{
// 			name: "Nil filter", filter: nil, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "Two existing minimal inputs, filter for title of first", numAffectedRecords: 1, insertBeforeDelete: true,
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two non-existing minimal inputs, filter for title of first",
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeDelete {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			numAffectedRecords, err := repo.DeleteWhere(context.Background(), test.filter)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 			assert.Equalf(t, test.numAffectedRecords, numAffectedRecords, test.name)
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryCountWhereTest(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		models             []*domain.Bookmark
// 		filter             *domain.BookmarkFilter
// 		err                error
// 		insertBeforeCount  bool
// 		numAffectedRecords int64
// 	}{
// 		{
// 			name: "Nil filter", filter: nil, err: helper.NilInputError{},
// 		},
// 		{
// 			name: "Two existing minimal entities, filter for title of first", numAffectedRecords: 1, insertBeforeCount: true,
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two non-existing minimal entities, filter for title of first",
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeCount {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			numAffectedRecords, err := repo.CountWhere(context.Background(), test.filter)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 			assert.Equalf(t, test.numAffectedRecords, numAffectedRecords, test.name)
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryCountAllTest(t *testing.T) {
// 	tests := []struct {
// 		name              string
// 		models            []*domain.Bookmark
// 		err               error
// 		insertBeforeCount bool
// 		numRecords        int64
// 	}{
// 		{
// 			name: "Two existing minimal entities, filter for title of first", numRecords: 2, insertBeforeCount: true,
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two non-existing minimal entities, filter for title of first",
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeCount {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			numRecords, err := repo.CountAll(context.Background())
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 			assert.Equalf(t, test.numRecords, numRecords, test.name)
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryDoesExistTest(t *testing.T) {
// 	tests := []struct {
// 		name              string
// 		model             *domain.Bookmark
// 		err               error
// 		insertBeforeCheck bool
// 		doesExist         bool
// 	}{
// 		{
// 			name: "Existing minimal entity", doesExist: true, insertBeforeCheck: true,
// 			model: &domain.Bookmark{
// 				CreatedAt:    time.Now(),
// 				UpdatedAt:    time.Now(),
// 				DeletedAt:    optional.Make(time.Now()),
// 				URL:          "https://example.com",
// 				Title:        optional.Make("My first bookmark"),
// 				ID:           1,
// 				IsCollection: false,
// 				IsRead:       true,
// 			},
// 		},
// 		{
// 			name: "Non-existing minimal entities",
// 			model: &domain.Bookmark{
// 				CreatedAt:    time.Now(),
// 				UpdatedAt:    time.Now(),
// 				DeletedAt:    optional.Make(time.Now()),
// 				URL:          "https://example.com",
// 				Title:        optional.Make("My first bookmark"),
// 				ID:           1,
// 				IsCollection: false,
// 				IsRead:       true,
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeCheck {
// 				err = repo.Add(context.Background(), []*domain.Bookmark{test.model})
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			doesExist, err := repo.DoesExist(context.Background(), test.model)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 			assert.Equalf(t, test.doesExist, doesExist, test.name)
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryDoesExistWhereTest(t *testing.T) {
// 	tests := []struct {
// 		name              string
// 		filter            *domain.BookmarkFilter
// 		models            []*domain.Bookmark
// 		err               error
// 		insertBeforeCheck bool
// 		doesExist         bool
// 	}{
// 		{
// 			name: "Two existing minimal entities, filter for title of first", doesExist: true, insertBeforeCheck: true,
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two existing minimal entities, filter for IsCollection of both", doesExist: true, insertBeforeCheck: true,
// 			filter: &domain.BookmarkFilter{
// 				IsCollection: optional.Make(model.FilterOperation[bool]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[bool]{Operand: false},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two non-existing minimal entities, filter for title of first",
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeCheck {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			doesExist, err := repo.DoesExistWhere(context.Background(), test.filter)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 			assert.Equalf(t, test.doesExist, doesExist, test.name)
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryGetWhereTest(t *testing.T) {
// 	tests := []struct {
// 		name              string
// 		filter            *domain.BookmarkFilter
// 		models            []*domain.Bookmark
// 		err               error
// 		insertBeforeCheck bool
// 		numRecords        int
// 	}{
// 		{
// 			name: "Two existing minimal entities, filter for title of first", numRecords: 1, insertBeforeCheck: true,
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two existing minimal entities, filter for IsCollection of both", numRecords: 2, insertBeforeCheck: true,
// 			filter: &domain.BookmarkFilter{
// 				IsCollection: optional.Make(model.FilterOperation[bool]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[bool]{Operand: false},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two non-existing minimal entities, filter for title of first",
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeCheck {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			records, err := repo.GetWhere(context.Background(), test.filter)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 			assert.Equalf(t, test.numRecords, len(records), test.name)
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryGetFirstWhereTest(t *testing.T) {
// 	tests := []struct {
// 		name              string
// 		filter            *domain.BookmarkFilter
// 		models            []*domain.Bookmark
// 		err               error
// 		insertBeforeCheck bool
// 		numRecords        int
// 	}{
// 		{
// 			name: "Two existing minimal entities, filter for title of first", numRecords: 1, insertBeforeCheck: true,
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two existing minimal entities, filter for IsCollection of both", numRecords: 2, insertBeforeCheck: true,
// 			filter: &domain.BookmarkFilter{
// 				IsCollection: optional.Make(model.FilterOperation[bool]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[bool]{Operand: false},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "Two non-existing minimal entities, filter for title of first", err: &helper.IneffectiveOperationError{},
// 			filter: &domain.BookmarkFilter{
// 				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
// 					Operator: model.FilterEqual,
// 					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
// 				}),
// 			},
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeCheck {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			_, err = repo.GetFirstWhere(context.Background(), test.filter)
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 		})
// 	}
// }

// func TestSQLBookmarkRepositoryGetAllTest(t *testing.T) {
// 	tests := []struct {
// 		name              string
// 		models            []*domain.Bookmark
// 		err               error
// 		insertBeforeCheck bool
// 		numRecords        int
// 	}{
// 		{
// 			name: "Two existing minimal entities", numRecords: 2, insertBeforeCheck: true,
// 			models: []*domain.Bookmark{
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://example.com",
// 					Title:        optional.Make("My first bookmark"),
// 					ID:           1,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 				{
// 					CreatedAt:    time.Now(),
// 					UpdatedAt:    time.Now(),
// 					DeletedAt:    optional.Make(time.Now()),
// 					URL:          "https://foo.example.com",
// 					Title:        optional.Make("My second bookmark"),
// 					ID:           2,
// 					IsCollection: false,
// 					IsRead:       true,
// 				},
// 			},
// 		},
// 		{
// 			name: "No entities", numRecords: 0,
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			t.Parallel()
// 			db, err := testCommon.GetDB()
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo := new(repository.Sqlite3TagRepository)

// 			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
// 			assert.NoErrorf(t, err, test.name)

// 			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

// 			repo := new(repository.Sqlite3BookmarkRepository)

// 			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
// 			assert.NoErrorf(t, err, test.name)

// 			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

// 			if test.insertBeforeCheck {
// 				err = repo.Add(context.Background(), test.models)
// 				if test.err == nil {
// 					assert.NoErrorf(t, err, test.name)
// 				} else {
// 					assert.ErrorAsf(t, err, &test.err, test.name)
// 				}
// 			}

// 			records, err := repo.GetAll(context.Background())
// 			if test.err == nil {
// 				assert.NoErrorf(t, err, test.name)
// 			} else {
// 				assert.ErrorAsf(t, err, &test.err, test.name)
// 			}
// 			assert.Equalf(t, test.numRecords, len(records), test.name)
// 		})
// 	}
// }

func TestSQLBookmarkRepositoryAddTypeTest(t *testing.T) {
	tests := []struct {
		name         string
		type_        []string
		preAddedType []string
		err          error
	}{
		{
			name: "One new type", type_: []string{"Text"},
		},
		{
			name: "Two new types", type_: []string{"Text", "Image"},
		},
		{
			name: "One duplicate type", preAddedType: []string{"Text"}, type_: []string{"Text"},
		},
		{
			name: "Two types, one duplicate", preAddedType: []string{"Text"}, type_: []string{"Text", "Image"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3BookmarkRepository)

			repoAbstract, err := repo.New(repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3BookmarkRepository)

			if test.preAddedType != nil {
				err = repo.AddType(context.Background(), test.preAddedType)
				if test.err == nil {
					assert.NoErrorf(t, err, test.name)
				} else {
					assert.ErrorAsf(t, err, &test.err, test.name)
				}
			}

			err = repo.AddType(context.Background(), test.type_)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}
