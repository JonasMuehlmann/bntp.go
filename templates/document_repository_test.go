package repository_test

import (
	"context"
	"testing"

	"time"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	repositoryCommon "github.com/JonasMuehlmann/bntp.go/model/repository"
	repository "github.com/JonasMuehlmann/bntp.go/model/repository/sqlite3"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/JonasMuehlmann/optional.go"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestSQLDocumentRepositoryAddTest(t *testing.T) {
	tests := []struct {
		err    error
		name   string
		models []*domain.Document
		tags   []*domain.Tag
	}{
		{
			name: "Empty input", models: []*domain.Document{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Document{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Document{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two regular inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					// These tags do not exist!
					TagIDs: []int64{1},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					// These tags do not exist!
					TagIDs: []int64{1},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           2,
				},
			},
		},
		{
			name: "Two minimal inputs", models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

			err = repo.Add(context.Background(), test.models)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLDocumentRepositoryReplaceTest(t *testing.T) {
	tests := []struct {
		err            error
		name           string
		previousModels []*domain.Document
		models         []*domain.Document
	}{
		{
			name: "Empty input", models: []*domain.Document{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Document{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Document{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, adding non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{},
			previousModels: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",

					ID: 1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",

					ID: 2,
				},
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					// These tags do not exist!
					TagIDs: []int64{1},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					// These tags do not exist!
					TagIDs: []int64{1},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           2,
				},
			},
		},
		{
			name: "Two existing minimal inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate!
					Path: "path/to/file",
					ID:   2,
				},
			},
		},

		{
			name: "Two existing minimal inputs", models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
			previousModels: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryUpsertTest(t *testing.T) {
	tests := []struct {
		err            error
		name           string
		previousModels []*domain.Document
		models         []*domain.Document
	}{
		{
			name: "Empty input", models: []*domain.Document{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Document{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Document{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{},
			previousModels: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",

					ID: 1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",

					ID: 2,
				},
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					// These tags do not exist!
					TagIDs: []int64{1},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					TagIDs:    []int64{1},
					// This type does not exist
					DocumentType: optional.Make("Note"),
					ID:           2,
				},
			},
		},

		{
			name: "Two existing inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate
					Path: "path/to/file",
					ID:   2,
				},
			},
		},

		{
			name: "Two existing minimal inputs",
			previousModels: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryUpdateTest(t *testing.T) {
	tests := []struct {
		err            error
		updater        *domain.DocumentUpdater
		name           string
		previousModels []*domain.Document
		models         []*domain.Document
	}{
		{
			name: "Empty input", models: []*domain.Document{}, updater: &domain.DocumentUpdater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, updater: &domain.DocumentUpdater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil updater", models: []*domain.Document{{}}, updater: nil, err: helper.NilInputError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Document{nil}, updater: &domain.DocumentUpdater{}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Document{{}}, updater: &domain.DocumentUpdater{}, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, nop updater", updater: &domain.DocumentUpdater{}, err: helper.IneffectiveOperationError{},
			previousModels: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{
			name: "Two existing inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					// This is a duplicate
					Path: "path/to/file",
					ID:   2,
				},
			},
		},

		{

			name: "Two existing minimal inputs, prepend to Path",
			updater: &domain.DocumentUpdater{
				Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "new/"}),
			},

			previousModels: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryUpdateWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.DocumentFilter
		updater            *domain.DocumentUpdater
		name               string
		models             []*domain.Document
		numAffectedRecords int64
		insertBeforeUpdate bool
	}{
		{
			name: "No entities", updater: &domain.DocumentUpdater{}, filter: &domain.DocumentFilter{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil updater", updater: nil, filter: &domain.DocumentFilter{}, err: helper.NilInputError{},
		},
		{
			name: "Nil filter", updater: &domain.DocumentUpdater{}, filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal inputs, filter for path of first, prepend to path", numAffectedRecords: 1, insertBeforeUpdate: true,
			updater: &domain.DocumentUpdater{
				Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "new/"}),
			},
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{

			name: "Two existing minimal inputs, prepend to path", numAffectedRecords: 2, insertBeforeUpdate: true, filter: &domain.DocumentFilter{},
			updater: &domain.DocumentUpdater{
				Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "new/"}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},

		{
			name: "Two existing minimal inputs, adding duplicated values", numAffectedRecords: 0, insertBeforeUpdate: true, filter: &domain.DocumentFilter{}, err: helper.DuplicateInsertionError{},

			updater: &domain.DocumentUpdater{
				Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateSet, Operand: "path/to/file"}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryDeleteTest(t *testing.T) {
	tests := []struct {
		err                error
		name               string
		models             []*domain.Document
		insertBeforeDelete bool
	}{
		{
			name: "Empty input", models: []*domain.Document{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Document{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Document{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two minimal inputs", insertBeforeDelete: true, models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryDeleteWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.DocumentFilter
		name               string
		models             []*domain.Document
		numAffectedRecords int64
		insertBeforeDelete bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal inputs, filter for path of first", numAffectedRecords: 1, insertBeforeDelete: true,
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{

			name: "Two non-existing minimal inputs, filter for path",
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryCountWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.DocumentFilter
		name               string
		models             []*domain.Document
		numAffectedRecords int64
		insertBeforeCount  bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal entities, filter for path of first", numAffectedRecords: 1, insertBeforeCount: true,
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{

			name: "Two existing minimal entities, filter for path",
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryCountAllTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.Document
		numRecords        int64
		insertBeforeCount bool
	}{
		{
			name: "Two existing minimal entities, filter for all", numRecords: 2, insertBeforeCount: true,
			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{
			name: "Two non-existing minimal entities, filter for all",
			models: []*domain.Document{
				{},
				{},
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryDoesExistTest(t *testing.T) {
	tests := []struct {
		err               error
		model             *domain.Document
		name              string
		insertBeforeCheck bool
		doesExist         bool
	}{
		{
			name: "Nil input", err: helper.NilInputError{},
		},
		{
			name: "Existing minimal entity", doesExist: true, insertBeforeCheck: true,
			model: &domain.Document{

				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: optional.Make(time.Now()),
				Path:      "path/to/file",
				ID:        1,
			},
		},
		{
			name: "Non-existing minimal entities",
			model: &domain.Document{

				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: optional.Make(time.Now()),
				Path:      "path/to/file",
				ID:        1,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

			if test.insertBeforeCheck {
				err = repo.Add(context.Background(), []*domain.Document{test.model})
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

func TestSQLDocumentRepositoryDoesExistWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.DocumentFilter
		name              string
		models            []*domain.Document
		insertBeforeCheck bool
		doesExist         bool
	}{
		{
			name: "Nil input", filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal inputs, filter for path of first", doesExist: true, insertBeforeCheck: true,
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{

			name: "Two existing minimal inputs, filter for path of both", doesExist: true, insertBeforeCheck: true,
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterLike,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/%"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{

			name: "Two existing minimal inputs, filter for path of first",
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryGetWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.DocumentFilter
		name              string
		models            []*domain.Document
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil input", filter: nil, err: helper.NilInputError{},
		},
		{
			name: "Empty result", err: helper.IneffectiveOperationError{},

			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},
		},
		{

			name: "Two existing minimal entities, filter for path of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{

			name: "Two existing minimal entities, filter for path of both", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterLike,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/%"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{

			name: "Two existing minimal entities, filter for path of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryGetFirstWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.DocumentFilter
		name              string
		models            []*domain.Document
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal inputs, filter for path of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{

			name: "Two existing minimal inputs, filter for path of both", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterLike,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/%"},
				}),
			},

			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
				},
			},
		},
		{

			name: "Two non-existing minimal entities, filter for path of first", err: &helper.IneffectiveOperationError{},
			filter: &domain.DocumentFilter{
				Path: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "path/to/file"},
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryGetAllTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.Document
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Two existing minimal entities", numRecords: 2, insertBeforeCheck: true,
			models: []*domain.Document{
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/file",
					ID:        1,
				},
				{

					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: optional.Make(time.Now()),
					Path:      "path/to/other/file",
					ID:        2,
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryGetFromIDsTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.Document
		ids               []int64
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil input", ids: nil, err: helper.NilInputError{},
		},
		{

			name: "Two existing minimal inputs, get all", numRecords: 2, insertBeforeCheck: true,
			models: []*domain.Document{
				{
					Path: "https://foo.example.com",
					ID:   1,
				},
				{
					Path: "https://bar.example.com",
					ID:   3,
				},
			},
			ids: []int64{1, 3},
		},
		{
			name: "Two existing minimal inputs, get one ", numRecords: 1, insertBeforeCheck: true,
			models: []*domain.Document{
				{
					Path: "https://foo.example.com",
					ID:   1,
				},
				{
					Path: "https://bar.example.com",
					ID:   3,
				},
			},
			ids: []int64{1},
		},
		{
			name: "Two existing minimal inputs, IDs don't exist", numRecords: 0, insertBeforeCheck: true,
			models: []*domain.Document{
				{
					Path: "https://foo.example.com",
					ID:   1,
				},
				{
					Path: "https://bar.example.com",
					ID:   3,
				},
			},
			ids: []int64{4, 5},
			err: helper.NonExistentPrimaryDataError{},
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

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

			if test.insertBeforeCheck {
				err = repo.Add(context.Background(), test.models)
				assert.NoErrorf(t, err, test.name)
			}

			records, err := repo.GetFromIDs(context.Background(), test.ids)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
			assert.Equalf(t, test.numRecords, len(records), test.name)
		})
	}
}

func TestSQLDocumentRepositoryAddTypeTest(t *testing.T) {
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryUpdateTypeTest(t *testing.T) {
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
			name: "Non-existent old type", oldTypes: []string{"Text"}, newType: "Text", err: helper.NonExistentPrimaryDataError{},
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryDeleteTypeTest(t *testing.T) {
	tests := []struct {
		err           error
		name          string
		type_         []string
		preAddedTypes []string
	}{
		{
			name: "Deleting non-existent type", type_: []string{"Text"}, err: helper.IneffectiveOperationError{},
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

			tagRepo := new(repository.Sqlite3TagRepository)

			tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
			assert.NoErrorf(t, err, test.name)

			tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

			repo := new(repository.Sqlite3DocumentRepository)

			repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

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

func TestSQLDocumentRepositoryTagModelConverter(t *testing.T) {
	t.Parallel()
	defer testCommon.HandlePanic(t, t.Name())

	db, err := testCommon.GetDB()
	require.NoError(t, err, ", db open")
	defer db.Close()

	tagRepo := new(repository.Sqlite3TagRepository)

	tagRepoAbstract, err := tagRepo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
	assert.NoError(t, err)

	tagRepo = tagRepoAbstract.(*repository.Sqlite3TagRepository)

	repo := new(repository.Sqlite3DocumentRepository)

	repoAbstract, err := repo.New(repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger(), TagRepository: tagRepo})

	assert.NoError(t, err)

	repo = repoAbstract.(*repository.Sqlite3DocumentRepository)

	refOut := &domain.Document{
		Path:                   "path/to//other/file",
		DocumentType:           optional.Make("Note"),
		BacklinkedDocumentsIDs: []int64{1},
		ID:                     2,
	}
	refIn := &domain.Document{
		Path:              "path/to//some/file",
		DocumentType:      optional.Make("Note"),
		LinkedDocumentIDs: []int64{1},
		ID:                3,
	}

	original := &domain.Document{
		Path:                   "path/to/file",
		DocumentType:           optional.Make("Note"),
		LinkedDocumentIDs:      []int64{2},
		BacklinkedDocumentsIDs: []int64{3},
		TagIDs:                 []int64{1},
		ID:                     1,
	}

	tag := &domain.Tag{
		Tag: "Programming languages",
		ID:  1,
	}

	boil.DebugMode = true

	err = repo.GetTagRepository().Add(context.Background(), []*domain.Tag{tag})
	assert.NoError(t, err)

	err = repo.AddType(context.Background(), []string{original.DocumentType.Wrappee})
	assert.NoError(t, err)

	documents := []*domain.Document{refOut, refIn, original}

	err = repo.Add(context.Background(), documents)
	assert.NoError(t, err)

	for _, document := range documents {
		repositoryModel, err := repo.DocumentDomainToRepositoryModel(context.Background(), document)
		assert.NoError(t, err)

		convertedBack, err := repo.DocumentRepositoryToDomainModel(context.Background(), repositoryModel.(*repository.Document))
		assert.NoError(t, err)
		assert.EqualValues(t, document, convertedBack)

		err = repositoryModel.(*repository.Document).Reload(context.Background(), db)
		assert.NoError(t, err)
		assert.EqualValues(t, document, convertedBack)

		convertedBack, err = repo.DocumentRepositoryToDomainModel(context.Background(), repositoryModel.(*repository.Document))
		assert.NoError(t, err)
		assert.EqualValues(t, document, convertedBack)
	}
}
