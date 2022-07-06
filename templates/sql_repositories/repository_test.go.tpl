package repository_test

import (
	"context"
	"testing"
    {{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
	"time"
    {{ end }}

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	repositoryCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository/sqlite3"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/stretchr/testify/assert"
)

func TestSQL{{.EntityName}}RepositoryAddTest(t *testing.T) {
	tests := []struct {
		err    error
		name   string
		models []*domain.{{.EntityName}}
	}{
		{
			name: "Empty input", models: []*domain.{{.EntityName}}{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.{{.EntityName}}{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.{{.EntityName}}{{"{{}}"}},
		},
		{
			name: "Two regular inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://example.com",
					Title:     optional.Make("My first bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					// This type does not exist
					{{.EntityName}}Type: optional.Make("Text"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
                    {{"}}"}},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           1,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{{"{{"}}
                            // This tag does not exist
                            Tag:        "Go",
                            ParentPath: []*domain.Tag{},
                            Subtags:    []*domain.Tag{},
                            ID:         2,
                        {{"}}"}},
						ID:         1,
					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://foo.example.com",
					Title:     optional.Make("My second bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					// This type does not exist
					{{.EntityName}}Type: optional.Make("Text"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
                    {{"}}"}},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{{"{{"}}
                            // This tag does not exist
                            Tag:        "Linux",
                            ParentPath: []*domain.Tag{},
                            Subtags:    []*domain.Tag{},
                            ID:         4,
                        {{"}}"}},
						ID:         3,

					{{end}}
				},
			},
		},
		{
			name: "Two minimal inputs", models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,

					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,

					{{end}}
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

			err = repo.Add(context.Background(), test.models)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQL{{.EntityName}}RepositoryReplaceTest(t *testing.T) {
	tests := []struct {
		err            error
		name           string
		previousModels []*domain.{{.EntityName}}
		models         []*domain.{{.EntityName}}
	}{
		{
			name: "Empty input", models: []*domain.{{.EntityName}}{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.{{.EntityName}}{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.{{.EntityName}}{{"{{}}"}}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Two existing minimal inputs, adding non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{},
			previousModels: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
					Tags: []*domain.Tag{},
					ID:           1,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,

					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
					Tags: []*domain.Tag{},
					ID:           2,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,

					{{end}}
				},
			},

			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://example.com",
					Title:     optional.Make("My first bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
                    {{"}}"}},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           1,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{{"{{"}}
                            // This tag does not exist
                            Tag:        "Go",
                            ParentPath: []*domain.Tag{},
                            Subtags:    []*domain.Tag{},
                            ID:         2,
                        {{"}}"}},
						ID:         1,

					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://foo.example.com",
					Title:     optional.Make("My second bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
                    {{"}}"}},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           2,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{{"{{"}}
                            // This tag does not exist
                            Tag:        "Linux",
                            ParentPath: []*domain.Tag{},
                            Subtags:    []*domain.Tag{},
                            ID:         4,
                        {{"}}"}},
						ID:         3,

					{{end}}
				},
			},
		},
		{
			name: "Two existing minimal inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},

			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{{"{{"}}
                            // This tag does not exist
                            Tag:        "Go",
                            ParentPath: []*domain.Tag{},
                            Subtags:    []*domain.Tag{},
                            ID:         2,
                        {{"}}"}},
						ID:         1,

					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate!
					URL:          "https://example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate!
					Path:       "path/to/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{{"{{"}}
                            // This tag does not exist
                            Tag:        "Linux",
                            ParentPath: []*domain.Tag{},
                            Subtags:    []*domain.Tag{},
                            ID:         4,
                        {{"}}"}},
						ID:         3,

					{{end}}
				},
			},
		},

		{
			name: "Two existing minimal inputs", models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: true,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
			previousModels: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: true,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryUpsertTest(t *testing.T) {
	tests := []struct {
		err            error
		name           string
		previousModels []*domain.{{.EntityName}}
		models         []*domain.{{.EntityName}}
	}{
		{
			name: "Empty input", models: []*domain.{{.EntityName}}{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.{{.EntityName}}{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.{{.EntityName}}{{"{{}}"}},
		},
		{
			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{},
			previousModels: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					Tags: []*domain.Tag{},
					ID: 1,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					Tags: []*domain.Tag{},
					ID:           2,
					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},

			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://example.com",
					Title:     optional.Make("My first bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
                    {{"}}"}},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID: 1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{{"{{"}}
                            // This tag does not exist
                            Tag:        "Go",
                            ParentPath: []*domain.Tag{},
                            Subtags:    []*domain.Tag{},
                            ID:         2,
                        {{"}}"}},
						ID:         1,

					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					URL:       "https://foo.example.com",
					Title:     optional.Make("My second bookmark"),
					// These tags do not exist!
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
					}},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					Tags: []*domain.Tag{{"{{"}}
						Tag:        "Test",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,
                    {{"}}"}},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{{"{{"}}
                            // This tag does not exist
                            Tag:        "Linux",
                            ParentPath: []*domain.Tag{},
                            Subtags:    []*domain.Tag{},
                            ID:         4,
                        {{"}}"}},
						ID:         3,

					{{end}}
				},
			},
		},
		{
			name: "Two existing inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},

			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate
					URL:          "https://example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate
					Path:       "path/to/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
			name: "Two existing minimal inputs",
			previousModels: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: true,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: true,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryUpdateTest(t *testing.T) {
	tests := []struct {
		err            error
		updater        *domain.{{.EntityName}}Updater
		name           string
		previousModels []*domain.{{.EntityName}}
		models         []*domain.{{.EntityName}}
	}{
		{
			name: "Empty input", models: []*domain.{{.EntityName}}{}, updater: &domain.{{.EntityName}}Updater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, updater: &domain.{{.EntityName}}Updater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil updater", models: []*domain.{{.EntityName}}{{"{{}}"}}, updater: nil, err: helper.NilInputError{},
		},
		{
			name: "Input containing nil value", models: []*domain.{{.EntityName}}{nil}, updater: &domain.{{.EntityName}}Updater{}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.{{.EntityName}}{{"{{}}"}}, updater: &domain.{{.EntityName}}Updater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Two existing minimal inputs, nop updater", updater: &domain.{{.EntityName}}Updater{}, err: helper.IneffectiveOperationError{},
			previousModels: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
			name: "Two existing inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					Tags:         []*domain.Tag{},
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					Tags:         []*domain.Tag{},
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
                        Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},

			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate
					URL:          "https://example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate
					Path:       "path/to/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},

		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal inputs, overwrite IsCollection",
			updater: &domain.{{.EntityName}}Updater{
				IsCollection: optional.Make(model.UpdateOperation[bool]{Operator: model.UpdateSet, Operand: true}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal inputs, prepend to Path",
			updater: &domain.{{.EntityName}}Updater{
                Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "new/"}),
			},
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, prepend to Tag",
			updater: &domain.{{.EntityName}}Updater{
                Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "New "}),
			},
            {{ end }}
			previousModels: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,

					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},

			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryUpdateWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.{{.EntityName}}Filter
		updater            *domain.{{.EntityName}}Updater
		name               string
		models             []*domain.{{.EntityName}}
		numAffectedRecords int64
		insertBeforeUpdate bool
	}{
		{
			name: "No entities", updater: &domain.{{.EntityName}}Updater{}, filter: &domain.{{.EntityName}}Filter{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil updater", updater: nil, filter: &domain.{{.EntityName}}Filter{}, err: helper.NilInputError{},
		},
		{
			name: "Nil filter", updater: &domain.{{.EntityName}}Updater{}, filter: nil, err: helper.NilInputError{},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal inputs, filter for title of first, update IsCollection", numAffectedRecords: 1, insertBeforeUpdate: true,
			updater: &domain.{{.EntityName}}Updater{
				IsCollection: optional.Make(model.UpdateOperation[bool]{Operator: model.UpdateSet, Operand: true}),
			},
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal inputs, filter for path of first, prepend to path", numAffectedRecords: 1, insertBeforeUpdate: true,
			updater: &domain.{{.EntityName}}Updater{
				Path : optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "new/"}),
			},
			filter: &domain.{{.EntityName}}Filter{
				Path : optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of first, prepend to Tag",numAffectedRecords: 1, insertBeforeUpdate: true,
			updater: &domain.{{.EntityName}}Updater{
                Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "New "}),
			},
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal inputs, overwrite IsCollection", numAffectedRecords: 2, insertBeforeUpdate: true, filter: &domain.{{.EntityName}}Filter{},
			updater: &domain.{{.EntityName}}Updater{
				IsCollection: optional.Make(model.UpdateOperation[bool]{Operator: model.UpdateSet, Operand: true}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal inputs, prepend to path", numAffectedRecords: 2, insertBeforeUpdate: true, filter: &domain.{{.EntityName}}Filter{},
			updater: &domain.{{.EntityName}}Updater{
				Path : optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "new/"}),
			},
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, prepend to Tag",numAffectedRecords: 2, insertBeforeUpdate: true, filter: &domain.{{.EntityName}}Filter{},
			updater: &domain.{{.EntityName}}Updater{
                Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "New "}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
			name: "Two existing minimal inputs, adding duplicated values", numAffectedRecords: 0, insertBeforeUpdate: true, filter: &domain.{{.EntityName}}Filter{}, err: helper.DuplicateInsertionError{},
            {{ if eq .EntityName "Bookmark" }}
			updater: &domain.{{.EntityName}}Updater{
				URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateSet, Operand: "https://example.com"}),
			},
            {{ else if eq .EntityName "Document" }}
			updater: &domain.{{.EntityName}}Updater{
				Path : optional.Make(model.UpdateOperation[string]{Operator: model.UpdateSet, Operand: "path/to/file"}),
			},
            {{ else if eq .EntityName "Tag" }}
			updater: &domain.{{.EntityName}}Updater{
                Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "New "}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryDeleteTest(t *testing.T) {
	tests := []struct {
		err                error
		name               string
		models             []*domain.{{.EntityName}}
		insertBeforeDelete bool
	}{
		{
			name: "Empty input", models: []*domain.{{.EntityName}}{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.{{.EntityName}}{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.{{.EntityName}}{{"{{}}"}},
		},
		{
			name: "Two minimal inputs", insertBeforeDelete: true, models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryDeleteWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.{{.EntityName}}Filter
		name               string
		models             []*domain.{{.EntityName}}
		numAffectedRecords int64
		insertBeforeDelete bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal inputs, filter for title of first", numAffectedRecords: 1, insertBeforeDelete: true,
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal inputs, filter for path of first", numAffectedRecords: 1, insertBeforeDelete: true,
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of first", numAffectedRecords: 1, insertBeforeDelete: true,
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two non-existing minimal inputs, filter for title",
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two non-existing minimal inputs, filter for path",
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two non-existing minimal inputs, filter for tag of first",
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryCountWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.{{.EntityName}}Filter
		name               string
		models             []*domain.{{.EntityName}}
		numAffectedRecords int64
		insertBeforeCount  bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal entities, filter for title of first", numAffectedRecords: 1, insertBeforeCount: true,
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal entities, filter for path of first", numAffectedRecords: 1, insertBeforeCount: true,
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of first", numAffectedRecords: 1, insertBeforeCount: true,
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two non-existing minimal entities, filter for title",
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal entities, filter for path",
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag",
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
            },
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryCountAllTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.{{.EntityName}}
		numRecords        int64
		insertBeforeCount bool
	}{
		{
			name: "Two existing minimal entities, filter for all", numRecords: 2, insertBeforeCount: true,
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
			name: "Two non-existing minimal entities, filter for all",
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryDoesExistTest(t *testing.T) {
	tests := []struct {
		err               error
		model             *domain.{{.EntityName}}
		name              string
		insertBeforeCheck bool
		doesExist         bool
	}{
		{
			name: "Nil input", err: helper.NilInputError{},
		},
		{
			name: "Existing minimal entity", doesExist: true, insertBeforeCheck: true,
			model: &domain.{{.EntityName}}{
				{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				DeletedAt:    optional.Make(time.Now()),
				URL:          "https://example.com",
				Title:        optional.Make("My first bookmark"),
				ID:           1,
				IsCollection: false,
				IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
			},
		},
		{
			name: "Non-existing minimal entities",
			model: &domain.{{.EntityName}}{
				{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				DeletedAt:    optional.Make(time.Now()),
				URL:          "https://example.com",
				Title:        optional.Make("My first bookmark"),
				ID:           1,
				IsCollection: false,
				IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

			if test.insertBeforeCheck {
				err = repo.Add(context.Background(), []*domain.{{.EntityName}}{test.model})
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

func TestSQL{{.EntityName}}RepositoryDoesExistWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.{{.EntityName}}Filter
		name              string
		models            []*domain.{{.EntityName}}
		insertBeforeCheck bool
		doesExist         bool
	}{
		{
			name: "Nil input", filter: nil, err: helper.NilInputError{},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal entities, filter for title of first", doesExist: true, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal inputs, filter for path of first", doesExist: true, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Path : optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of first",doesExist: true, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal entities, filter for IsCollection of both", doesExist: true, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				IsCollection: optional.Make(model.FilterOperation[bool]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[bool]{Operand: false},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal inputs, filter for path of both", doesExist: true, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterLike,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/%"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of both", doesExist: true, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterNEqual,
					Operand:  model.ScalarOperand[string]{Operand: ""},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two non-existing minimal entities, filter for title of first",
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal inputs, filter for path of first",
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of first",
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryGetWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.{{.EntityName}}Filter
		name              string
		models            []*domain.{{.EntityName}}
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil input", filter: nil, err: helper.NilInputError{},
		},
		{
			name: "Empty result", err: helper.IneffectiveOperationError{},
            {{ if eq .EntityName "Bookmark" }}
            filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal entities, filter for title of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal entities, filter for path of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of first",numRecords: 1, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal entities, filter for IsCollection of both", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				IsCollection: optional.Make(model.FilterOperation[bool]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[bool]{Operand: false},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal entities, filter for path of both", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterLike,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/%"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of ",numRecords: 2, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterNEqual,
					Operand:  model.ScalarOperand[string]{Operand: ""},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,

					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two non-existing minimal entities, filter for title of first", insertBeforeCheck: true, numRecords: 1,
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal entities, filter for path of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of first",numRecords: 1, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryGetFirstWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.{{.EntityName}}Filter
		name              string
		models            []*domain.{{.EntityName}}
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal entities, filter for title of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal inputs, filter for path of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of first",numRecords: 1, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two existing minimal entities, filter for IsCollection of both", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				IsCollection: optional.Make(model.FilterOperation[bool]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[bool]{Operand: false},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two existing minimal inputs, filter for path of both", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterLike,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/%"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of both",numRecords: 2, insertBeforeCheck: true,
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterNEqual,
					Operand:  model.ScalarOperand[string]{Operand: ""},
				}),
			},
            {{ end }}
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
				},
			},
		},
		{
            {{ if eq .EntityName "Bookmark" }}
			name: "Two non-existing minimal entities, filter for title of first", err: &helper.IneffectiveOperationError{},
			filter: &domain.{{.EntityName}}Filter{
				Title: optional.Make(model.FilterOperation[optional.Optional[string]]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[optional.Optional[string]]{Operand: optional.Make("My first bookmark")},
				}),
			},
            {{ else if eq .EntityName "Document" }}
			name: "Two non-existing minimal entities, filter for path of first", err: &helper.IneffectiveOperationError{},
			filter: &domain.{{.EntityName}}Filter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
            },
            {{ else if eq .EntityName "Tag" }}
			name: "Two existing minimal inputs, filter for tag of first",err: &helper.IneffectiveOperationError{},
			filter: &domain.{{.EntityName}}Filter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
            {{ end }}
        },
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryGetAllTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.{{.EntityName}}
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Two existing minimal entities", numRecords: 2, insertBeforeCheck: true,
			models: []*domain.{{.EntityName}}{
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://example.com",
					Title:        optional.Make("My first bookmark"),
					ID:           1,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/file",
                    ID:         1,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Programming",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         1,


					{{end}}
				},
				{
					{{ if eq .EntityName "Bookmark" }}
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    optional.Make(time.Now()),
					URL:          "https://foo.example.com",
					Title:        optional.Make("My second bookmark"),
					ID:           2,
					IsCollection: false,
					IsRead:       true,
					{{ end }}
					{{ if eq .EntityName "Document" }}
                    CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:       "path/to/other/file",
                    ID:         2,


					{{ end }}

					{{ if eq .EntityName "Tag" }}
						Tag:        "Operating Systems",
						ParentPath: []*domain.Tag{},
						Subtags:    []*domain.Tag{},
						ID:         3,


					{{end}}
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
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
func TestSQL{{.EntityName}}RepositoryAddTypeTest(t *testing.T) {
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
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryUpdateTypeTest(t *testing.T) {
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
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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

func TestSQL{{.EntityName}}RepositoryDeleteTypeTest(t *testing.T) {
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
			db, err := testCommon.GetDB()
			assert.NoErrorf(t, err, test.name)

			{{if or (eq .EntityName "Bookmark") (eq .EntityName "Document") }}
		    tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)
		    {{ end }}

			repo := new(repository.Sqlite3{{.EntityName}}Repository)

			{{if and (ne .EntityName "Bookmark") (ne .EntityName "Document") }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db })
		    {{ else }}
		    repoAbstract, err := repo.New(repository.Sqlite3{{.EntityName}}RepositoryConstructorArgs{DB: db, TagRepository: tagRepo})
		    {{ end }}
			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3{{.EntityName}}Repository)

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
{{ end }}
