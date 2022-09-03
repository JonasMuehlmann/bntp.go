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

func TestCmdTagAdd(t *testing.T) {
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
				"tag",
				"add",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad args",
			args: []string{
				"tag",
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
				"tag",
				"add",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Add duplicate",
			args: []string{
				"tag",
				"add",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "bar"}))),
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagReplace(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldTags         []*domain.Tag
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name:            "No args",
			args:            []string{"tag", "replace"},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad args",
			args: []string{
				"tag",
				"replace",
				"foo",
				"bar",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name:    "Good minimal arg",
			oldTags: []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			args: []string{
				"tag",
				"replace",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name:    "Some tags don't exist yet",
			oldTags: []*domain.Tag{{ID: 1, Tag: "foo"}},
			args: []string{
				"tag",
				"replace",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar2"}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "All tags don't exist yet",
			args: []string{
				"tag",
				"replace",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 3, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 4, Tag: "bar2"}))),
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

			if test.oldTags != nil {
				cli.TagReplaceCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.TagManager.Add(context.Background(), test.oldTags)
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagUpsert(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldTags         []*domain.Tag
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name:            "No args",
			args:            []string{"tag", "upsert"},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad args",
			args: []string{
				"tag",
				"upsert",
				"foo",
				"bar",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name:    "Good minimal args",
			oldTags: []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			args: []string{
				"tag",
				"upsert",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name:    "Updating and adding tags",
			oldTags: []*domain.Tag{{ID: 1, Tag: "foo"}},
			args: []string{
				"tag",
				"upsert",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Updating tags",
			args: []string{
				"tag",
				"upsert",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 3, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 4, Tag: "bar2"}))),
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

			if test.oldTags != nil {
				cli.TagUpsertCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.TagManager.Add(context.Background(), test.oldTags)
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagEdit(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldTags         []*domain.Tag
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "Only updater",
			args: []string{
				"tag",
				"edit",
				"--updater",
				string(drop.From2To1(json.Marshal(domain.TagUpdater{Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad models",
			args: []string{
				"tag",
				"edit",
				"foo",
				"bar",
				"--updater",
				string(drop.From2To1(json.Marshal(domain.TagUpdater{Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Minimal args, missing updater",
			args: []string{
				"tag",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar2"}))),
			},
			errorMatcher:    testCommon.ValidatorContains("required", "updater"),
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("required"),
		},
		{
			name: "Minimal args, bad updater",
			args: []string{
				"tag",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar2"}))),
				"--updater",
				"foo",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name:    "Good minimal args",
			oldTags: []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			args: []string{
				"tag",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar"}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.TagUpdater{Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.NumAffectedRecords{2}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"tag",
				"edit",
				"--filter",
				"foo",
				"--updater",
				string(drop.From2To1(json.Marshal(domain.TagUpdater{Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Filter and model",
			args: []string{
				"tag",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.TagUpdater{Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             cmd.ConflictingPositionalArgsAndFlagError{"filter"},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("together with"),
		},
		{
			name:    "Good minimal args, filter instead of models",
			oldTags: []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			args: []string{
				"tag",
				"edit",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.TagUpdater{Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.NumAffectedRecords{1}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags don't exist yet",
			args: []string{
				"tag",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar2"}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.TagUpdater{Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "All tags don't exist yet",
			args: []string{
				"tag",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 3, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 4, Tag: "bar2"}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.TagUpdater{Tag: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
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

			if test.oldTags != nil {
				cli.TagEditCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.TagManager.Add(context.Background(), test.oldTags)
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagExport(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		tags            []*domain.Tag
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No args",
			args: []string{
				"tag",
				"export",
			},
			errorMatcher:    testCommon.ValidatorContains("arg", "received"),
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("accepts", "received"),
		},
		{
			name: "No tags",
			args: []string{
				"tag",
				"export",
				"out",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Some tags",
			tags: []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			args: []string{
				"tag",
				"export",
				"out",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
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
				cli.TagEditCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.TagManager.Add(context.Background(), test.tags)
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagImport(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		fileContent     string
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No args",
			args: []string{
				"tag",
				"import",
			},
			errorMatcher:    testCommon.ValidatorContains("arg", "received"),
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("accepts", "received"),
		}, {
			name: "File does not exist",
			args: []string{
				"tag",
				"import",
				"in",
			},
			errorMatcher:    testCommon.ValidatorContains("does not exist"),
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		{
			name: "No tags",
			args: []string{
				"tag",
				"import",
				"in",
			},
			fileContent:     " ",
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Some tags",
			args: []string{
				"tag",
				"import",
				"in",
			},
			fileContent:     string(drop.From2To1(json.Marshal([]domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}}))),
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

			if test.fileContent != "" {
				err = afero.WriteFile(fs, "in", []byte(test.fileContent), 0o644)
				assert.NoError(t, err, test.name+", assert writing tags to import")
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagList(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tags            []*domain.Tag
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No tags",
			args: []string{
				"tag",
				"list",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Some tags",
			args: []string{
				"tag",
				"list",
			},
			tags:            []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			outputValidator: testCommon.ValidatorContains("foo", "bar"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"tag",
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
				"tag",
				"list",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags, using filter",
			args: []string{
				"tag",
				"list",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tags:            []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
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
				cli.TagListCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.TagManager.Add(context.Background(), test.tags)
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagRemove(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldTags         []*domain.Tag
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "Bad models",
			args: []string{
				"tag",
				"remove",
				"foo",
				"bar",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name:    "Good minimal args",
			oldTags: []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			args: []string{
				"tag",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar"}))),
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.NumAffectedRecords{2}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"tag",
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
				"tag",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             cmd.ConflictingPositionalArgsAndFlagError{"filter"},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("together with"),
		},
		{
			name:    "Good minimal args, filter instead of models",
			oldTags: []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			args: []string{
				"tag",
				"remove",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.NumAffectedRecords{1}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags don't exist yet",
			args: []string{
				"tag",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 2, Tag: "bar2"}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "All tags don't exist yet",
			args: []string{
				"tag",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 3, Tag: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 4, Tag: "bar2"}))),
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

			if test.oldTags != nil {
				cli.TagRemoveCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.TagManager.Add(context.Background(), test.oldTags)
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagFind(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tags            []*domain.Tag
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "Bad filter",
			args: []string{
				"tag",
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
				"tag",
				"find-first",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags",
			args: []string{
				"tag",
				"find-first",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tags:            []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
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
				cli.TagFindCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.TagManager.Add(context.Background(), test.tags)
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagCount(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tags            []*domain.Tag
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No tags, no filter",
			args: []string{
				"tag",
				"count",
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.Count{0}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Some tags, no filter",
			args: []string{
				"tag",
				"count",
			},
			tags:            []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.Count{2}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"tag",
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
				"tag",
				"count",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.Count{0}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags, using filter",
			args: []string{
				"tag",
				"count",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tags:            []*domain.Tag{{ID: 1, Tag: "foo"}, {ID: 2, Tag: "bar"}},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.Count{1}))) + "\n"),
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
				cli.TagCountCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.TagManager.Add(context.Background(), test.tags)
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}

func TestCmdTagDoesExist(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tag             *domain.Tag
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No tags",
			args: []string{
				"tag",
				"does-exist",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Some tags",
			args: []string{
				"tag",
				"does-exist",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo"}))),
			},
			tag:             &domain.Tag{ID: 1, Tag: "foo"},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.DoesExist{true}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad filter",
			args: []string{
				"tag",
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
				"tag",
				"does-exist",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.DoesExist{false}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags, using filter",
			args: []string{
				"tag",
				"does-exist",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tag:             &domain.Tag{ID: 1, Tag: "foo"},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.DoesExist{true}))) + "\n"),
			errorValidator:  testCommon.ValidatorEmpty,
		},
		{
			name: "Bad model",
			args: []string{
				"tag",
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
				"tag",
				"does-exist",
				string(drop.From2To1(json.Marshal(domain.Tag{ID: 1, Tag: "foo"}))),
				"--filter",
				string(drop.From2To1(json.Marshal(domain.TagFilter{Tag: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
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
				cli.TagDoesExistCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.TagManager.Add(context.Background(), []*domain.Tag{test.tag})
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
				test.errorMatcher(t, err.Error(), test.name+", assert error string matches")
			} else {
				assert.NoError(t, err, test.name+", assert test does not error unexpectedly")
			}
		})
	}
}
