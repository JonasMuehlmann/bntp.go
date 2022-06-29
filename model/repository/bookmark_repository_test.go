package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	repositoryCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository/sqlite3"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/stretchr/testify/assert"
)

func TestSQLBookmarkRepositoryAddTest(t *testing.T) {
	tests := []struct {
		name   string
		models []*domain.Bookmark
		err    error
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Bookmark{{}},
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
		t.Run(test.name, func(t *testing.T) {
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
		name               string
		models             []*domain.Bookmark
		err                error
		insertBeforeUpdate bool
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Bookmark{{}}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, insertBeforeUpdate: true, models: []*domain.Bookmark{
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
			name: "Two existing minimal inputs", insertBeforeUpdate: true, models: []*domain.Bookmark{
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
		t.Run(test.name, func(t *testing.T) {
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

			if test.insertBeforeUpdate {
				err = repo.Add(context.Background(), test.models)
				if test.err == nil {
					assert.NoErrorf(t, err, test.name)
				} else {
					assert.ErrorAsf(t, err, &test.err, test.name)
				}
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
		name               string
		models             []*domain.Bookmark
		err                error
		insertBeforeUpdate bool
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: []*domain.Bookmark{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Bookmark{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Bookmark{{}},
		},
		{
			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, insertBeforeUpdate: true, models: []*domain.Bookmark{
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
			name: "Two existing minimal inputs", insertBeforeUpdate: true, models: []*domain.Bookmark{
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
		t.Run(test.name, func(t *testing.T) {
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

			if test.insertBeforeUpdate {
				err = repo.Add(context.Background(), test.models)
				if test.err == nil {
					assert.NoErrorf(t, err, test.name)
				} else {
					assert.ErrorAsf(t, err, &test.err, test.name)
				}
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
		name               string
		models             []*domain.Bookmark
		updater            *domain.BookmarkUpdater
		err                error
		insertBeforeUpdate bool
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{}, updater: &domain.BookmarkUpdater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: []*domain.Bookmark{}, updater: &domain.BookmarkUpdater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil updater", models: []*domain.Bookmark{{}}, updater: nil, err: helper.NilInputError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Bookmark{nil}, updater: &domain.BookmarkUpdater{}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Bookmark{{}}, updater: &domain.BookmarkUpdater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, updater: &domain.BookmarkUpdater{}, insertBeforeUpdate: true, models: []*domain.Bookmark{
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
			name: "Two existing minimal inputs", insertBeforeUpdate: true, updater: &domain.BookmarkUpdater{}, models: []*domain.Bookmark{
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
		t.Run(test.name, func(t *testing.T) {
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

			if test.insertBeforeUpdate {
				err = repo.Add(context.Background(), test.models)
				if test.err == nil {
					assert.NoErrorf(t, err, test.name)
				} else {
					assert.ErrorAsf(t, err, &test.err, test.name)
				}
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
