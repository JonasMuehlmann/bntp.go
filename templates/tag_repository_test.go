package repository_test

import (
	"context"
	"testing"

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
)

// FIX: Some tests show false positives because we do not check for non-existent tags.
func TestSQLTagRepositoryAddTest(t *testing.T) {
	tests := []struct {
		err    error
		name   string
		models []*domain.Tag
		tags   []*domain.Tag
	}{
		{
			name: "Empty input", models: []*domain.Tag{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Tag{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Tag{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two regular inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{}, models: []*domain.Tag{
				{
					Tag: "Programming",

					// This tag does not exist
					SubtagIDs: []int64{2},
					ID:        1,
				},
				{
					Tag: "Operating Systems",

					// This tag does not exist
					SubtagIDs: []int64{4},
					ID:        3,
				},
			},
		},
		{
			name: "Two minimal inputs", models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "One minimal input, generating ID", models: []*domain.Tag{
				{
					Tag: "Programming",
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

			err = repo.Add(context.Background(), test.models)
			if test.err == nil {
				assert.NoErrorf(t, err, test.name)
			} else {
				assert.ErrorAsf(t, err, &test.err, test.name)
			}
		})
	}
}

func TestSQLTagRepositoryReplaceTest(t *testing.T) {
	tests := []struct {
		err            error
		name           string
		previousModels []*domain.Tag
		models         []*domain.Tag
		tags           []*domain.Tag
	}{
		{
			name: "Empty input", models: []*domain.Tag{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Tag{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Tag{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, adding non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{},
			previousModels: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},

			models: []*domain.Tag{
				{
					Tag:       "Programming",
					SubtagIDs: []int64{2},
					ID:        1,
				},
				{
					Tag:       "Operating Systems",
					SubtagIDs: []int64{4},
					ID:        3,
				},
			},
		},
		{
			name: "Two existing minimal inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					// This tag does not exist
					SubtagIDs: []int64{2},
					ID:        1,
				},
				{
					Tag: "Operating Systems",

					// This tag does not exist
					SubtagIDs: []int64{4},
					ID:        3,
				},
			},
		},
		{
			name: "Two existing minimal inputs", models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
			previousModels: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryUpsertTest(t *testing.T) {
	tests := []struct {
		err            error
		name           string
		previousModels []*domain.Tag
		models         []*domain.Tag
		tags           []*domain.Tag
	}{
		{
			name: "Empty input", models: []*domain.Tag{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Tag{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Tag{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two existing inputs, non-existent dependencies", err: repositoryCommon.ReferenceToNonExistentDependencyError{},
			previousModels: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					// This tag does not exist
					SubtagIDs: []int64{2},
					ID:        1,
				},
				{
					Tag: "Operating Systems",

					// This tag does not exist
					SubtagIDs: []int64{4},
					ID:        3,
				},
			},
		},

		{
			name: "Two existing minimal inputs",
			previousModels: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryUpdateTest(t *testing.T) {
	tests := []struct {
		err            error
		updater        *domain.TagUpdater
		name           string
		previousModels []*domain.Tag
		models         []*domain.Tag
		tags           []*domain.Tag
	}{
		{
			name: "Empty input", models: []*domain.Tag{}, updater: &domain.TagUpdater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, updater: &domain.TagUpdater{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil updater", models: []*domain.Tag{{}}, updater: nil, err: helper.NilInputError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Tag{nil}, updater: &domain.TagUpdater{}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Tag{{}}, updater: &domain.TagUpdater{}, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, nop updater", updater: &domain.TagUpdater{}, err: helper.IneffectiveOperationError{},
			previousModels: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two existing inputs, adding duplicated values", err: helper.DuplicateInsertionError{},
			previousModels: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Programming",

					ID: 3,
				},
			},
		},

		{
			name: "Two existing minimal inputs, prepend to Tag",
			updater: &domain.TagUpdater{
				Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "New "}),
			},

			previousModels: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryUpdateWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.TagFilter
		updater            *domain.TagUpdater
		name               string
		models             []*domain.Tag
		tags               []*domain.Tag
		numAffectedRecords int64
		insertBeforeUpdate bool
	}{
		{
			name: "No entities", updater: &domain.TagUpdater{}, filter: &domain.TagFilter{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil updater", updater: nil, filter: &domain.TagFilter{}, err: helper.NilInputError{},
		},
		{
			name: "Nil filter", updater: &domain.TagUpdater{}, filter: nil, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, filter for tag of first, prepend to Tag", numAffectedRecords: 1, insertBeforeUpdate: true,
			updater: &domain.TagUpdater{
				Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "New "}),
			},
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two existing minimal inputs, prepend to Tag", numAffectedRecords: 2, insertBeforeUpdate: true, filter: &domain.TagFilter{},
			updater: &domain.TagUpdater{
				Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdatePrepend, Operand: "New "}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryDeleteTest(t *testing.T) {
	tests := []struct {
		err                error
		name               string
		models             []*domain.Tag
		tags               []*domain.Tag
		insertBeforeDelete bool
	}{
		{
			name: "Empty input", models: []*domain.Tag{}, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Nil input", models: nil, err: helper.IneffectiveOperationError{},
		},
		{
			name: "Input containing nil value", models: []*domain.Tag{nil}, err: helper.NilInputError{},
		},
		{
			name: "One default-constructed input", models: []*domain.Tag{{}}, err: helper.NilInputError{},
		},
		{
			name: "Two minimal inputs", insertBeforeDelete: true, models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryDeleteWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.TagFilter
		name               string
		models             []*domain.Tag
		tags               []*domain.Tag
		numAffectedRecords int64
		insertBeforeDelete bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, filter for tag of first", numAffectedRecords: 1, insertBeforeDelete: true,
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two non-existing minimal inputs, filter for tag of first",
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryCountWhereTest(t *testing.T) {
	tests := []struct {
		err                error
		filter             *domain.TagFilter
		name               string
		models             []*domain.Tag
		tags               []*domain.Tag
		numAffectedRecords int64
		insertBeforeCount  bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, filter for tag of first", numAffectedRecords: 1, insertBeforeCount: true,
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two existing minimal inputs, filter for tag",
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryCountAllTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.Tag
		tags              []*domain.Tag
		numRecords        int64
		insertBeforeCount bool
	}{
		{
			name: "Two existing minimal entities, filter for all", numRecords: 2, insertBeforeCount: true,
			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two non-existing minimal entities, filter for all",
			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryDoesExistTest(t *testing.T) {
	tests := []struct {
		err               error
		model             *domain.Tag
		name              string
		insertBeforeCheck bool
		doesExist         bool
	}{
		{
			name: "Nil input", err: helper.NilInputError{},
		},
		{
			name: "Existing minimal entity", doesExist: true, insertBeforeCheck: true,
			model: &domain.Tag{
				Tag: "Programming",

				ID: 1,
			},
		},
		{
			name: "Non-existing minimal entities",
			model: &domain.Tag{
				Tag: "Programming",

				ID: 1,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

			if test.insertBeforeCheck {
				err = repo.Add(context.Background(), []*domain.Tag{test.model})
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

func TestSQLTagRepositoryDoesExistWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.TagFilter
		name              string
		models            []*domain.Tag
		tags              []*domain.Tag
		insertBeforeCheck bool
		doesExist         bool
	}{
		{
			name: "Nil input", filter: nil, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, filter for tag of first", doesExist: true, insertBeforeCheck: true,
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two existing minimal inputs, filter for tag of both", doesExist: true, insertBeforeCheck: true,
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterNEqual,
					Operand:  model.ScalarOperand[string]{Operand: ""},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two existing minimal inputs, filter for tag of first",
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryGetWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.TagFilter
		name              string
		models            []*domain.Tag
		tags              []*domain.Tag
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil input", filter: nil, err: helper.NilInputError{},
		},
		{
			name: "Empty result", err: helper.IneffectiveOperationError{},

			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},
		},
		{
			name: "Two existing minimal inputs, filter for tag of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two existing minimal inputs, filter for tag of ", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterNEqual,
					Operand:  model.ScalarOperand[string]{Operand: ""},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two existing minimal inputs, filter for tag of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryGetFirstWhereTest(t *testing.T) {
	tests := []struct {
		err               error
		filter            *domain.TagFilter
		name              string
		models            []*domain.Tag
		tags              []*domain.Tag
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil filter", filter: nil, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, filter for tag of first", numRecords: 1, insertBeforeCheck: true,
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two existing minimal inputs, filter for tag of both", numRecords: 2, insertBeforeCheck: true,
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterNEqual,
					Operand:  model.ScalarOperand[string]{Operand: ""},
				}),
			},

			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
				},
			},
		},
		{
			name: "Two existing minimal inputs, filter for tag of first", err: &helper.IneffectiveOperationError{},
			filter: &domain.TagFilter{
				Tag: optional.Make(model.FilterOperation[string]{
					Operator: model.FilterEqual,
					Operand:  model.ScalarOperand[string]{Operand: "Programming"},
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryGetAllTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.Tag
		tags              []*domain.Tag
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Two existing minimal entities", numRecords: 2, insertBeforeCheck: true,
			models: []*domain.Tag{
				{
					Tag: "Programming",

					ID: 1,
				},
				{
					Tag: "Operating Systems",

					ID: 3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryGetFromIDsTest(t *testing.T) {
	tests := []struct {
		err               error
		name              string
		models            []*domain.Tag
		ids               []int64
		numRecords        int
		insertBeforeCheck bool
	}{
		{
			name: "Nil input", ids: nil, err: helper.NilInputError{},
		},
		{
			name: "Two existing minimal inputs, get all", numRecords: 2, insertBeforeCheck: true,
			models: []*domain.Tag{
				{
					Tag: "Programming",
					ID:  1,
				},
				{
					Tag: "Operating Systems",
					ID:  3,
				},
			},
			ids: []int64{1, 3},
		},
		{
			name: "Two existing minimal inputs, get one ", numRecords: 1, insertBeforeCheck: true,
			models: []*domain.Tag{
				{
					Tag: "Programming",
					ID:  1,
				},
				{
					Tag: "Operating Systems",
					ID:  3,
				},
			},
			ids: []int64{1},
		},
		{
			name: "Two existing minimal inputs, IDs don't exist", numRecords: 0, insertBeforeCheck: true,
			models: []*domain.Tag{
				{
					Tag: "Programming",
					ID:  1,
				},
				{
					Tag: "Operating Systems",
					ID:  3,
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

			repo := new(repository.Sqlite3TagRepository)

			repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})

			assert.NoErrorf(t, err, test.name)

			repo = repoAbstract.(*repository.Sqlite3TagRepository)

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

func TestSQLTagRepositoryTagModelConverter(t *testing.T) {
	t.Parallel()
	defer testCommon.HandlePanic(t, t.Name())

	db, err := testCommon.GetDB()
	require.NoError(t, err, ", db open")
	defer db.Close()

	repo := new(repository.Sqlite3TagRepository)

	repoAbstract, err := repo.New(repository.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: log.StandardLogger()})
	assert.NoError(t, err)

	repo = repoAbstract.(*repository.Sqlite3TagRepository)

	parent1 := &domain.Tag{
		Tag: "Software development",

		ID:        1,
		SubtagIDs: []int64{2},
	}

	parent2 := &domain.Tag{
		Tag: "Computer science",

		ID:            2,
		SubtagIDs:     []int64{5},
		ParentPathIDs: []int64{1},
	}

	child1 := &domain.Tag{
		Tag: "Golang",

		ID:            3,
		ParentPathIDs: []int64{5, 2, 1},
	}

	child2 := &domain.Tag{
		Tag: "CPP",

		ID:            4,
		ParentPathIDs: []int64{5, 2, 1},
	}

	original := &domain.Tag{
		Tag: "Programming languages",

		ID:            5,
		SubtagIDs:     []int64{3, 4},
		ParentPathIDs: []int64{2, 1},
	}

	tags := []*domain.Tag{parent1, parent2, original, child1, child2}

	err = repo.Add(context.Background(), tags)
	assert.NoError(t, err)

	for _, tag := range tags {
		repositoryModel, err := repo.TagDomainToRepositoryModel(context.Background(), tag)
		assert.NoError(t, err)

		convertedBack, err := repo.TagRepositoryToDomainModel(context.Background(), repositoryModel.(*repository.Tag))
		assert.NoError(t, err)
		assert.EqualValues(t, tag, convertedBack)
	}
}
