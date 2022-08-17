package repository_test

import (
	"context"
	"testing"

	"time"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	repositoryCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository/mssql"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSQLBookmarkRepositoryAddTest(t *testing.T) {
	tests := []struct {
		err    error
		name   string
		models []*domain.Bookmark
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Bookmark{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two regular inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, models: []*domain.Bookmark{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://example.com",
					Title:     optional.Make("My first bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					// This type does not exist
					BookmarkType: optional.Make("Text"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://foo.example.com",
					Title:     optional.Make("My second bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					// This type does not exist
					BookmarkType: optional.Make("Text"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{
			name: "Two minimal inputs", models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			err = repo.Add(context.Background(), test.models)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLBookmarkRepositoryReplaceTest(t *testing.T) {
	tests := []struct {
		err            error
		name           string
		previousModels []*domain.Bookmark
		models         []*domain.Bookmark
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Bookmark{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, adding non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{},
			previousModels: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},

			models: []*domain.Bookmark{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://example.com",
					Title:     optional.Make("My first bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://foo.example.com",
					Title:     optional.Make("My second bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{
			name: "Two existing minimal inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate!
					URL:          "https://example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},

		{
			name: "Two existing minimal inputs", models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: true,
					IsRead:       true,
				},
			},
			previousModels: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: true,
					IsRead:       true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.previousModels != nil {
				err = repo.Add(context.Background(), test.previousModels)
				assert.NoErrorf(t, err, test.name)
			}

			err = repo.Replace(context.Background(), test.models)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLBookmarkRepositoryUpsertTest(t *testing.T) {
	tests := []struct {
		err            error
		name           string
		previousModels []*domain.Bookmark
		models         []*domain.Bookmark
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Bookmark{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{},
			previousModels: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},

			models: []*domain.Bookmark{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://example.com",
					Title:     optional.Make("My first bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://foo.example.com",
					Title:     optional.Make("My second bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{
			name: "Two existing inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate
					URL:          "https://example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{
			name: "Two existing minimal inputs",
			previousModels: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: true,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: true,
					IsRead:       true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.previousModels != nil {
				err = repo.Add(context.Background(), test.previousModels)
				assert.NoErrorf(t, err, test.name)
			}

			err = repo.Upsert(context.Background(), test.models)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLBookmarkRepositoryUpdateTest(t *testing.T) {
	tests := []struct {
		err            error
		updater        *domain.BookmarkUpdater
		name           string
		previousModels []*domain.Bookmark
		models         []*domain.Bookmark
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{}, updater: &domain.BookmarkUpdater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, updater: &domain.BookmarkUpdater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil updater", models: []*domain.Bookmark{{}}, updater: nil, err: helper.NilInputError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Bookmark{nil}, updater: &domain.BookmarkUpdater{}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Bookmark{{}}, updater: &domain.BookmarkUpdater{}, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, nop updater", updater: &domain.BookmarkUpdater{}, err: helper.IneffectiveOperationError{},
			previousModels: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{
			name: "Two existing inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate
					URL:          "https://example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},

		{

			name: "Two existing minimal inputs, overwrite IsCollection",
			updater: &domain.BookmarkUpdater{
				IsCollection: optional.Make(model.UpdateOperation[bool]{Operator: model.UpdateSet, Operand: true}),
			},

			previousModels: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.previousModels != nil {
				err = repo.Add(context.Background(), test.previousModels)
				assert.NoErrorf(t, err, test.name)
			}

			err = repo.Update(context.Background(), test.models, test.updater)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLBookmarkRepositoryUpdateWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.BookmarkFilter
		updater            *domain.BookmarkUpdater
		name               string
		models             []*domain.Bookmark
		numAffectedRecords int64
		insertBeforeUpdate bool
	}{
		{
			name: "No entities", updater: &domain.BookmarkUpdater{}, filter: &domain.BookmarkFilter{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil updater", updater: nil, filter: &domain.BookmarkFilter{}, err: helper.NilInputError{},
		},
		{
			name: "Nil filter", updater: &domain.BookmarkUpdater{}, filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal inputs, filter for title of first, update IsCollection", numAffectedRecords: 1, insertBeforeUpdate: true,
			updater: &domain.BookmarkUpdater{
				IsCollection: optional.Make(model.UpdateOperation[bool]{Operator: model.UpdateSet, Operand: true}),
			},
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{

			name: "Two existing minimal inputs, overwrite IsCollection", numAffectedRecords: 2, insertBeforeUpdate: true, filter: &domain.BookmarkFilter{},
			updater: &domain.BookmarkUpdater{
				IsCollection: optional.Make(model.UpdateOperation[bool]{Operator: model.UpdateSet, Operand: true}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{
			name: "Two existing minimal inputs, adding duplicated values", numAffectedRecords: 0, insertBeforeUpdate: true, filter: &domain.BookmarkFilter{}, err: helper.DuplicateInsertionError{},

			updater: &domain.BookmarkUpdater{
				URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateSet, Operand: "https://example.com"}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeUpdate {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			numAffectedRecords, err := repo.UpdateWhere(context.Background(), test.filter, test.updater)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
			assert.Equalf(t, test.numAffectedRecords, numAffectedRecords, test.name)
		})
	}
}

func TestSQLBookmarkRepositoryDeleteTest(t *testing.T) {
	tests := []struct {
		err                error
		name               string
		models             []*domain.Bookmark
		insertBeforeDelete bool
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Bookmark{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two minimal inputs", insertBeforeDelete: true, models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeDelete {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			err = repo.Delete(context.Background(), test.models)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLBookmarkRepositoryDeleteWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.BookmarkFilter
		name               string
		models             []*domain.Bookmark
		numAffectedRecords int64
		insertBeforeDelete bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal inputs, filter for title of first", numAffectedRecords: 1, insertBeforeDelete: true,
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{

			name: "Two non-existing minimal inputs, filter for title",
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeDelete {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			numAffectedRecords, err := repo.DeleteWhere(context.Background(), test.filter)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
			assert.Equalf(t, test.numAffectedRecords, numAffectedRecords, test.name)
		})
	}
}

func TestSQLBookmarkRepositoryCountWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.BookmarkFilter
		name               string
		models             []*domain.Bookmark
		numAffectedRecords int64
		insertBeforeCount  bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal entities, filter for title of first", numAffectedRecords: 1, insertBeforeCount: true,
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{

			name: "Two non-existing minimal entities, filter for title",
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeCount {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			numAffectedRecords, err := repo.CountWhere(context.Background(), test.filter)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
			assert.Equalf(t, test.numAffectedRecords, numAffectedRecords, test.name)
		})
	}
}

func TestSQLBookmarkRepositoryCountAllTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.Bookmark
		numRecords        int64
		insertBeforeCount bool
	}{
		{
			name: "Two existing minimal entities, filter for all", numRecords: 2, insertBeforeCount: true,
			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{
			name: "Two non-existing minimal entities, filter for all",
			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeCount {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			numRecords, err := repo.CountAll(context.Background())
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
			assert.Equalf(t, test.numRecords, numRecords, test.name)
		})
	}
}

func TestSQLBookmarkRepositoryDoesExistTest(t *testing.T) {
	tests := []struct {
		err               error
		model             *domain.Bookmark
		name              string
		insertBeforeCheck bool
		doesExist         bool
	}{
		{
			name: "Nil input", err: helper.NilInputError{},
		},
		{
			name: "Existing minimal entity", doesExist: true, insertBeforeCheck: true,
			model: &domain.Bookmark{

				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				DeletedAt:    optional.Make(time.Now()),
				URL:          "https://example.com",
				Title:        optional.Make("My first bookmark"),
				ID:           1,
				IsCollection: false,
				IsRead:       true,
			},
		},
		{
			name: "Non-existing minimal entities",
			model: &domain.Bookmark{

				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				DeletedAt:    optional.Make(time.Now()),
				URL:          "https://example.com",
				Title:        optional.Make("My first bookmark"),
				ID:           1,
				IsCollection: false,
				IsRead:       true,
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeCheck {
				err = repo.Add(context.Background(), []*domain.Bookmark{test.model})
				assert.NoErrorf(t, err, test.name)
			}

			doesExist, err := repo.DoesExist(context.Background(), test.model)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
			assert.Equalf(t, test.doesExist, doesExist, test.name)
		})
	}
}

func TestSQLBookmarkRepositoryDoesExistWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.BookmarkFilter
		name              string
		models            []*domain.Bookmark
		insertBeforeCheck bool
		doesExist         bool
	}{
		{
			name: "Nil input", filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal entities, filter for title of first", doesExist: true, insertBeforeCheck: true,
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{

			name: "Two existing minimal entities, filter for IsCollection of both", doesExist: true, insertBeforeCheck: true,
			filter: &domain.BookmarkFilter{
				IsCollection: optional.Make(model.FilterOperation[bool]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[bool]{Operand: false},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{

			name: "Two non-existing minimal entities, filter for title of first",
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeCheck {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			doesExist, err := repo.DoesExistWhere(context.Background(), test.filter)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
			assert.Equalf(t, test.doesExist, doesExist, test.name)
		})
	}
}

func TestSQLBookmarkRepositoryGetWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.BookmarkFilter
		name              string
		models            []*domain.Bookmark
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil input", filter: nil, err: helper.NilInputError{},
		},
		{
			name: "Empty result", err: helper.IneffectiveOperationError{},

			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
		},
		{

			name: "Two existing minimal entities, filter for title of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{

			name: "Two existing minimal entities, filter for IsCollection of both", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.BookmarkFilter{
				IsCollection: optional.Make(model.FilterOperation[bool]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[bool]{Operand: false},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{

			name: "Two non-existing minimal entities, filter for title of first", insertBeforeCheck: true, numRecords: 1,
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeCheck {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			records, err := repo.GetWhere(context.Background(), test.filter)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
			assert.Equalf(t, test.numRecords, len(records), test.name)
		})
	}
}

func TestSQLBookmarkRepositoryGetFirstWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.BookmarkFilter
		name              string
		models            []*domain.Bookmark
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal entities, filter for title of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{

			name: "Two existing minimal entities, filter for IsCollection of both", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.BookmarkFilter{
				IsCollection: optional.Make(model.FilterOperation[bool]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[bool]{Operand: false},
				}),
			},

			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{

			name: "Two non-existing minimal entities, filter for title of first", err: &helper.IneffectiveOperationError{},
			filter: &domain.BookmarkFilter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeCheck {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			_, err = repo.GetFirstWhere(context.Background(), test.filter)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLBookmarkRepositoryGetAllTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.Bookmark
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Two existing minimal entities", numRecords: 2, insertBeforeCheck: true,
			models: []*domain.Bookmark{
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
				},
				{

					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
				},
			},
		},
		{
			name: "No entities", numRecords: 0, err: helper.IneffectiveOperationError{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeCheck {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			records, err := repo.GetAll(context.Background())
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
			assert.Equalf(t, test.numRecords, len(records), test.name)
		})
	}
}

func TestSQLBookmarkRepositoryAddTypeTest(t *testing.T) {
	tests := []struct {
		err          error
		name         string
		type_        []string
		preAddedType []string
	}{
		{
			name: "One new type", type_: []string{"Text"},
		},
		{
			name: "Two new types", type_: []string{"Text", "Image"},
		},
		{
			name: "One duplicate type", preAddedType: []string{"Text"}, type_: []string{"Text"}, err: helper.DuplicateInsertionError{},
		},
		{
			name: "Two types, one duplicate", preAddedType: []string{"Text"}, type_: []string{"Text", "Image"}, err: helper.DuplicateInsertionError{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.preAddedType != nil {
				err = repo.AddType(context.Background(), test.preAddedType)
				assert.NoErrorf(t, err, test.name)
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

func TestSQLBookmarkRepositoryUpdateTypeTest(t *testing.T) {
	tests := []struct {
		err                error
		name               string
		newType            string
		oldTypes           []string
		insertBeforeUpdate bool
	}{
		{
			name: "New type", oldTypes: []string{"Text"}, newType: "Image", insertBeforeUpdate: true,
		},
		{
			name: "Rename to duplicate type", oldTypes: []string{"Text", "Image"}, newType: "Image", insertBeforeUpdate: true, err: helper.DuplicateInsertionError{},
		},
		{
			name: "Non-existent old type", oldTypes: []string{"Text"}, newType: "Text", err: helper.NonExistentPrimaryDataError,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.insertBeforeUpdate {
				err = repo.AddType(context.Background(), test.oldTypes)
				assert.NoErrorf(t, err, test.name)
			}

			err = repo.UpdateType(context.Background(), test.oldTypes[0], test.newType)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLBookmarkRepositoryDeleteTypeTest(t *testing.T) {
	tests := []struct {
		err           error
		name          string
		type_         []string
		preAddedTypes []string
	}{
		{
			name: "Deleting non-existent type", type_: []string{"Text"},
		},
		{
			name: "Deleting existing type", preAddedTypes: []string{"Text"}, type_: []string{"Text"},
		},
		{
			name: "Two types, one non-existent", preAddedTypes: []string{"Text"}, type_: []string{"Text", "Image"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)

			db, err := testCommon.GetDB()
			require.NoErrorf(t, err, test.name+", db open")
			defer db.Close()

			tagRepo := new(repository.MssqlTagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

			repo := new(repository.MssqlBookmarkRepository)

			repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.MssqlBookmarkRepository)

			if test.preAddedTypes != nil {
				err = repo.AddType(context.Background(), test.preAddedTypes)
				assert.NoErrorf(t, err, test.name)
			}

			err = repo.DeleteType(context.Background(), test.type_)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLBookmarkRepositoryTagModelConverter(t *testing.T) {
	t.Parallel()
	defer testCommon.HandlePanic(t, t.Name())

	db, err := testCommon.GetDB()
	require.NoError(t, err, ", db open")
	defer db.Close()

	tagRepo := new(repository.MssqlTagRepository)

	tagRepoAbstract, err := tagRepo.New(repository.MssqlTagRepositoryConstructorArgs{DB: db})
	assert.NoError(t, err)

	tagRepo = tagRepoAbstract.(*repository.MssqlTagRepository)

	repo := new(repository.MssqlBookmarkRepository)

	repoAbstract, err := repo.New(repository.MssqlBookmarkRepositoryConstructorArgs{DB: db, TagRepository: tagRepo})

	assert.NoError(t, err)

	repo = repoAbstract.(*repository.MssqlBookmarkRepository)

	parent1 := &domain.Tag{
		Tag:        "Software development",
		ParentPath: []*domain.Tag{},
		Subtags:    []*domain.Tag{},
		ID:         1,
	}

	parent2 := &domain.Tag{
		Tag:        "Computer science",
		ParentPath: []*domain.Tag{},
		Subtags:    []*domain.Tag{},
		ID:         2,
	}

	child1 := &domain.Tag{
		Tag:        "Golang",
		ParentPath: []*domain.Tag{},
		Subtags:    []*domain.Tag{},
		ID:         3,
	}

	child2 := &domain.Tag{
		Tag:        "CPP",
		ParentPath: []*domain.Tag{},
		Subtags:    []*domain.Tag{},
		ID:         4,
	}

	original := &domain.Tag{
		Tag:        "Programming languages",
		ParentPath: []*domain.Tag{},
		Subtags:    []*domain.Tag{},
		ID:         5,
	}

	original.AddChildren([]*domain.Tag{child1, child2})
	original.AddDirectParent(parent2)
	parent2.AddDirectParent(parent1)

	err = repo.GetTagRepository().Add(context.Background(), []*domain.Tag{parent1, parent2, child1, child2, original})

	assert.NoError(t, err)

	repositoryModel, err := repo.GetTagRepository().TagDomainToRepositoryModel(context.Background(), original)

	assert.NoError(t, err)

	convertedBack, err := repo.GetTagRepository().TagRepositoryToDomainModel(context.Background(), repositoryModel.(*repository.Tag))

	assert.NoError(t, err)
	assert.EqualValues(t, original, convertedBack)
}
