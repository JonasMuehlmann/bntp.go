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

func TestSQLBookmarkRepositoryTest(t *testing.T) {
	tests := []struct {
		name   string
		models []*domain.Bookmark
		err    error
	}{
		{
			name: "Empty input", models: []*domain.Bookmark{},
		},
		{
			name: "Nil input", models: []*domain.Bookmark{},
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
