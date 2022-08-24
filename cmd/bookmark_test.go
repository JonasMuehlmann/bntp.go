package cmd_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/cmd"
	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/bntp.go/model"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/JonasMuehlmann/drop-return-values.go"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCmdBookmarkAdd(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "no args",
			args: []string{
				"bookmark",
				"add",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad args",
			args: []string{
				"bookmark",
				"add",
				"foo",
				"bar",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Good minimal arg",
			args: []string{
				"bookmark",
				"add",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Add duplicate",
			args: []string{
				"bookmark",
				"add",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "bar"}))),
			},
			err:             helper.DuplicateInsertionError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("duplicate"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			outputBuffer := testCommon.NewBufferString("")
			errorBuffer := testCommon.NewBufferString("")
			fs := afero.NewMemMapFs()
			cli := cmd.NewCli(cmd.WithStdErrOverride(errorBuffer), cmd.WithDbOverride(db), cmd.WithFsOverride(fs), cmd.WithAll())
			cli.RootCmd.SetOut(outputBuffer)

			cli.RootCmd.SetArgs(test.args)

			err = cli.Execute()

			stdout := outputBuffer.String()
			stderr := errorBuffer.String()

			if test.outputValidator != nil {
				test.outputValidator(t, stdout, test.name+", assert stdout matches")
			}
			if test.errorValidator != nil {
				test.errorValidator(t, stderr, test.name+", assert stderr matches")
			}

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorValidator(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdBookmarkReplace(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldBookmarks    []*domain.Bookmark
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name:            "No args",
			args:            []string{"bookmark", "replace"},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad args",
			args: []string{
				"bookmark",
				"replace",
				"foo",
				"bar",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name:         "Good minimal arg",
			oldBookmarks: []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			args: []string{
				"bookmark",
				"replace",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name:         "Some tags don't exist yet",
			oldBookmarks: []*domain.Bookmark{{ID: 1, URL: "foo"}},
			args: []string{
				"bookmark",
				"replace",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar2"}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "All tags don't exist yet",
			args: []string{
				"bookmark",
				"replace",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 3, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 4, URL: "bar2"}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			outputBuffer := testCommon.NewBufferString("")
			errorBuffer := testCommon.NewBufferString("")
			fs := afero.NewMemMapFs()
			cli := cmd.NewCli(cmd.WithStdErrOverride(errorBuffer), cmd.WithDbOverride(db), cmd.WithFsOverride(fs), cmd.WithAll())
			cli.RootCmd.SetOut(outputBuffer)

			cli.RootCmd.SetArgs(test.args)

			if test.oldBookmarks != nil {
				cli.BookmarkReplaceCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.BookmarkManager.Add(context.Background(), test.oldBookmarks)
					assert.NoError(t, err, test.name+", assert adding old tags")
				}
			}

			err = cli.Execute()

			stdout := outputBuffer.String()
			stderr := errorBuffer.String()

			if test.outputValidator != nil {
				test.outputValidator(t, stdout, test.name+", assert stdout matches")
			}
			if test.errorValidator != nil {
				test.errorValidator(t, stderr, test.name+", assert stderr matches")
			}

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorValidator(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdBookmarkUpsert(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldBookmarks    []*domain.Bookmark
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name:            "No args",
			args:            []string{"bookmark", "upsert"},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad args",
			args: []string{
				"bookmark",
				"upsert",
				"foo",
				"bar",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name:         "Good minimal args",
			oldBookmarks: []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			args: []string{
				"bookmark",
				"upsert",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name:         "Updating and adding tags",
			oldBookmarks: []*domain.Bookmark{{ID: 1, URL: "foo"}},
			args: []string{
				"bookmark",
				"upsert",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Updating tags",
			args: []string{
				"bookmark",
				"upsert",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 3, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 4, URL: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			outputBuffer := testCommon.NewBufferString("")
			errorBuffer := testCommon.NewBufferString("")
			fs := afero.NewMemMapFs()
			cli := cmd.NewCli(cmd.WithStdErrOverride(errorBuffer), cmd.WithDbOverride(db), cmd.WithFsOverride(fs), cmd.WithAll())
			cli.RootCmd.SetOut(outputBuffer)

			cli.RootCmd.SetArgs(test.args)

			if test.oldBookmarks != nil {
				cli.BookmarkUpsertCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.BookmarkManager.Add(context.Background(), test.oldBookmarks)
					assert.NoError(t, err, test.name+", assert adding old tags")
				}
			}

			err = cli.Execute()

			stdout := outputBuffer.String()
			stderr := errorBuffer.String()

			if test.outputValidator != nil {
				test.outputValidator(t, stdout, test.name+", assert stdout matches")
			}
			if test.errorValidator != nil {
				test.errorValidator(t, stderr, test.name+", assert stderr matches")
			}

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorValidator(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdBookmarkEdit(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldBookmarks    []*domain.Bookmark
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "Only updater",
			args: []string{
				"bookmark",
				"edit",
				"--updater",
				string(drop.From2To1(json.Marshal(domain.BookmarkUpdater{URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad models",
			args: []string{
				"bookmark",
				"edit",
				"foo",
				"bar",
				"--updater",
				string(drop.From2To1(json.Marshal(domain.BookmarkUpdater{URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Minimal args, missing updater",
			args: []string{
				"bookmark",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar2"}))),
			},
			errorMatcher:    testCommon.ValidatorContains("required", "updater"),
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("required"),
		},
		{
			name: "Minimal args, bad updater",
			args: []string{
				"bookmark",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar2"}))),
				"--updater",
				"foo",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name:         "Good minimal args",
			oldBookmarks: []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			args: []string{
				"bookmark",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar"}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.BookmarkUpdater{URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			outputValidator: testCommon.ValidatorEqual("2\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"bookmark",
				"edit",
				"--filter",
				"foo",
				"--updater",
				string(drop.From2To1(json.Marshal(domain.BookmarkUpdater{URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Filter and model",
			args: []string{
				"bookmark",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.BookmarkUpdater{URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             cmd.ConflictingPositionalArgsAndFlagError{"filter"},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("together with"),
		},
		{
			name:         "Good minimal args, filter instead of models",
			oldBookmarks: []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			args: []string{
				"bookmark",
				"edit",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.BookmarkUpdater{URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			outputValidator: testCommon.ValidatorEqual("1\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags don't exist yet",
			args: []string{
				"bookmark",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar2"}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.BookmarkUpdater{URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "All tags don't exist yet",
			args: []string{
				"bookmark",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 3, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 4, URL: "bar2"}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.BookmarkUpdater{URL: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			outputBuffer := testCommon.NewBufferString("")
			errorBuffer := testCommon.NewBufferString("")
			fs := afero.NewMemMapFs()
			cli := cmd.NewCli(cmd.WithStdErrOverride(errorBuffer), cmd.WithDbOverride(db), cmd.WithFsOverride(fs), cmd.WithAll())
			cli.RootCmd.SetOut(outputBuffer)

			cli.RootCmd.SetArgs(test.args)

			if test.oldBookmarks != nil {
				cli.BookmarkEditCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.BookmarkManager.Add(context.Background(), test.oldBookmarks)
					assert.NoError(t, err, test.name+", assert adding old tags")
				}
			}

			err = cli.Execute()

			stdout := outputBuffer.String()
			stderr := errorBuffer.String()

			if test.outputValidator != nil {
				test.outputValidator(t, stdout, test.name+", assert stdout matches")
			}
			if test.errorValidator != nil {
				test.errorValidator(t, stderr, test.name+", assert stderr matches")
			}

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorValidator(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdBookmarkList(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tags            []*domain.Bookmark
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No tags",
			args: []string{
				"bookmark",
				"list",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Some tags",
			args: []string{
				"bookmark",
				"list",
			},
			tags:            []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			outputValidator: testCommon.ValidatorContains("foo", "bar"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"bookmark",
				"list",
				"--filter",
				"foo",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "no tags, using filter",
			args: []string{
				"bookmark",
				"list",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags, using filter",
			args: []string{
				"bookmark",
				"list",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tags:            []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			outputValidator: testCommon.ValidatorContains("foo"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			outputBuffer := testCommon.NewBufferString("")
			errorBuffer := testCommon.NewBufferString("")
			fs := afero.NewMemMapFs()
			cli := cmd.NewCli(cmd.WithStdErrOverride(errorBuffer), cmd.WithDbOverride(db), cmd.WithFsOverride(fs), cmd.WithAll())
			cli.RootCmd.SetOut(outputBuffer)

			cli.RootCmd.SetArgs(test.args)

			if test.tags != nil {
				cli.BookmarkListCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.BookmarkManager.Add(context.Background(), test.tags)
					assert.NoError(t, err, test.name+", assert adding old tags")
				}
			}

			err = cli.Execute()

			stdout := outputBuffer.String()
			stderr := errorBuffer.String()

			if test.outputValidator != nil {
				test.outputValidator(t, stdout, test.name+", assert stdout matches")
			}
			if test.errorValidator != nil {
				test.errorValidator(t, stderr, test.name+", assert stderr matches")
			}

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorValidator(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdBookmarkRemove(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldBookmarks    []*domain.Bookmark
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "Bad models",
			args: []string{
				"bookmark",
				"remove",
				"foo",
				"bar",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name:         "Good minimal args",
			oldBookmarks: []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			args: []string{
				"bookmark",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar"}))),
			},
			outputValidator: testCommon.ValidatorEqual("2\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"bookmark",
				"remove",
				"--filter",
				"foo",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Filter and model",
			args: []string{
				"bookmark",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             cmd.ConflictingPositionalArgsAndFlagError{"filter"},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("together with"),
		},
		{
			name:         "Good minimal args, filter instead of models",
			oldBookmarks: []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			args: []string{
				"bookmark",
				"remove",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			outputValidator: testCommon.ValidatorEqual("1\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags don't exist yet",
			args: []string{
				"bookmark",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 2, URL: "bar2"}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "All tags don't exist yet",
			args: []string{
				"bookmark",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 3, URL: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 4, URL: "bar2"}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			outputBuffer := testCommon.NewBufferString("")
			errorBuffer := testCommon.NewBufferString("")
			fs := afero.NewMemMapFs()
			cli := cmd.NewCli(cmd.WithStdErrOverride(errorBuffer), cmd.WithDbOverride(db), cmd.WithFsOverride(fs), cmd.WithAll())
			cli.RootCmd.SetOut(outputBuffer)

			cli.RootCmd.SetArgs(test.args)

			if test.oldBookmarks != nil {
				cli.BookmarkRemoveCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.BookmarkManager.Add(context.Background(), test.oldBookmarks)
					assert.NoError(t, err, test.name+", assert adding old tags")
				}
			}

			err = cli.Execute()

			stdout := outputBuffer.String()
			stderr := errorBuffer.String()

			if test.outputValidator != nil {
				test.outputValidator(t, stdout, test.name+", assert stdout matches")
			}
			if test.errorValidator != nil {
				test.errorValidator(t, stderr, test.name+", assert stderr matches")
			}

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorValidator(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdBookmarkFind(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tags            []*domain.Bookmark
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "Bad filter",
			args: []string{
				"bookmark",
				"find-first",
				"--filter",
				"foo",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "no tags",
			args: []string{
				"bookmark",
				"find-first",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags",
			args: []string{
				"bookmark",
				"find-first",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tags:            []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			outputValidator: testCommon.ValidatorContains("foo"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			outputBuffer := testCommon.NewBufferString("")
			errorBuffer := testCommon.NewBufferString("")
			fs := afero.NewMemMapFs()
			cli := cmd.NewCli(cmd.WithStdErrOverride(errorBuffer), cmd.WithDbOverride(db), cmd.WithFsOverride(fs), cmd.WithAll())
			cli.RootCmd.SetOut(outputBuffer)

			cli.RootCmd.SetArgs(test.args)

			if test.tags != nil {
				cli.BookmarkFindCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.BookmarkManager.Add(context.Background(), test.tags)
					assert.NoError(t, err, test.name+", assert adding old tags")
				}
			}

			err = cli.Execute()

			stdout := outputBuffer.String()
			stderr := errorBuffer.String()

			if test.outputValidator != nil {
				test.outputValidator(t, stdout, test.name+", assert stdout matches")
			}
			if test.errorValidator != nil {
				test.errorValidator(t, stderr, test.name+", assert stderr matches")
			}

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorValidator(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdBookmarkCount(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tags            []*domain.Bookmark
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No tags, no filter",
			args: []string{
				"bookmark",
				"count",
			},
			outputValidator: testCommon.ValidatorEqual("0\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Some tags, no filter",
			args: []string{
				"bookmark",
				"count",
			},
			tags:            []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			outputValidator: testCommon.ValidatorEqual("2\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"bookmark",
				"count",
				"--filter",
				"foo",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "no tags, using filter",
			args: []string{
				"bookmark",
				"count",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			outputValidator: testCommon.ValidatorEqual("0\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags, using filter",
			args: []string{
				"bookmark",
				"count",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tags:            []*domain.Bookmark{{ID: 1, URL: "foo"}, {ID: 2, URL: "bar"}},
			outputValidator: testCommon.ValidatorEqual("1\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			outputBuffer := testCommon.NewBufferString("")
			errorBuffer := testCommon.NewBufferString("")
			fs := afero.NewMemMapFs()
			cli := cmd.NewCli(cmd.WithStdErrOverride(errorBuffer), cmd.WithDbOverride(db), cmd.WithFsOverride(fs), cmd.WithAll())
			cli.RootCmd.SetOut(outputBuffer)

			cli.RootCmd.SetArgs(test.args)

			if test.tags != nil {
				cli.BookmarkCountCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.BookmarkManager.Add(context.Background(), test.tags)
					assert.NoError(t, err, test.name+", assert adding old tags")
				}
			}

			err = cli.Execute()

			stdout := outputBuffer.String()
			stderr := errorBuffer.String()

			if test.outputValidator != nil {
				test.outputValidator(t, stdout, test.name+", assert stdout matches")
			}
			if test.errorValidator != nil {
				test.errorValidator(t, stderr, test.name+", assert stderr matches")
			}

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorValidator(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdBookmarkDoesExist(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tag             *domain.Bookmark
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No tags",
			args: []string{
				"bookmark",
				"does-exist",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Some tags",
			args: []string{
				"bookmark",
				"does-exist",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo"}))),
			},
			tag:             &domain.Bookmark{ID: 1, URL: "foo"},
			outputValidator: testCommon.ValidatorContains("true\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"bookmark",
				"does-exist",
				"--filter",
				"foo",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "no tags, using filter",
			args: []string{
				"bookmark",
				"does-exist",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			outputValidator: testCommon.ValidatorEqual("false\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags, using filter",
			args: []string{
				"bookmark",
				"does-exist",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tag:             &domain.Bookmark{ID: 1, URL: "foo"},
			outputValidator: testCommon.ValidatorContains("true\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad model",
			args: []string{
				"bookmark",
				"does-exist",
				"foo",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Filter and model",
			args: []string{
				"bookmark",
				"does-exist",
				string(drop.From2To1(json.Marshal(domain.Bookmark{ID: 1, URL: "foo"}))),
				"--filter",
				string(drop.From2To1(json.Marshal(domain.BookmarkFilter{URL: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             cmd.ConflictingPositionalArgsAndFlagError{"filter"},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("together with"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, test.name)
			db, err := testCommon.GetDB()
			assert.NoError(t, err, test.name+", assert db creation")

			outputBuffer := testCommon.NewBufferString("")
			errorBuffer := testCommon.NewBufferString("")
			fs := afero.NewMemMapFs()
			cli := cmd.NewCli(cmd.WithStdErrOverride(errorBuffer), cmd.WithDbOverride(db), cmd.WithFsOverride(fs), cmd.WithAll())
			cli.RootCmd.SetOut(outputBuffer)

			cli.RootCmd.SetArgs(test.args)

			if test.tag != nil {
				cli.BookmarkDoesExistCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.BookmarkManager.Add(context.Background(), []*domain.Bookmark{test.tag})
					assert.NoError(t, err, test.name+", assert adding old tags")
				}
			}

			err = cli.Execute()

			stdout := outputBuffer.String()
			stderr := errorBuffer.String()

			if test.outputValidator != nil {
				test.outputValidator(t, stdout, test.name+", assert stdout matches")
			}
			if test.errorValidator != nil {
				test.errorValidator(t, stderr, test.name+", assert stderr matches")
			}

			if test.err != nil {
				assert.ErrorIs(t, err, test.err, test.name+", assert test error matches expected")
			} else if test.errorMatcher != nil {
				test.errorValidator(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}
