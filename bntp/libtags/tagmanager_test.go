package libtags_test

import (
	"context"
	"testing"

	bntp "github.com/JonasMuehlmann/bntp.go/bntp"
	"github.com/JonasMuehlmann/bntp.go/bntp/libtags"
	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	sqlite3Repo "github.com/JonasMuehlmann/bntp.go/model/repository/sqlite3"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLibtagsMarshalPath(t *testing.T) {
	tests := []struct {
		err          error
		errorMatcher testCommon.OutputValidator
		tag          *domain.Tag
		name         string
		path         string
		tags         []*domain.Tag
	}{
		{
			name: "no args",
			err:  helper.IneffectiveOperationError{},
		},
		{
			name: "no parent",
			tags: []*domain.Tag{{ID: 1, Tag: "foo"}},
			tag:  &domain.Tag{ID: 1, Tag: "foo"},
			path: "foo",
		},
		{
			name: "two parents",
			tags: []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}, {ID: 3, Tag: "baz"}},
			tag:  &domain.Tag{ID: 1, Tag: "foo", ParentPathIDs: []int64{2, 3}},
			path: "bar::baz::foo",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			tagRepoConcrete := &sqlite3Repo.Sqlite3TagRepository{}
			tagRepoAbstract, err := tagRepoConcrete.New(sqlite3Repo.Sqlite3TagRepositoryConstructorArgs{DB: db, Logger: logrus.StandardLogger()})
			tagRepoConcrete = tagRepoAbstract.(*sqlite3Repo.Sqlite3TagRepository)
			assert.NoError(t, err, test.name+", assert tag repository creation")

			tagManager, err := libtags.NewTagmanager(tagRepoConcrete.Logger, &bntp.Hooks[domain.Tag]{}, tagRepoConcrete)
			assert.NoError(t, err, test.name+", assert tag manager creation")

			if test.tags != nil {
				err = tagManager.Add(context.Background(), test.tags)
				assert.NoError(t, err, test.name+", assert tag creation")
			}

			path, err := tagManager.MarshalPath(context.Background(), test.tag, false)

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
				assert.Equal(t, test.path, path, test.name+", assert returned path matches")
			}

		})
	}
}
