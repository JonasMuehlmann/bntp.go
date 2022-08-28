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

func getDocumentSkeleton() string {
	return "# Tags\n\n# Links\n\n# Backlinks\n"
}

func TestCmdDocumentAdd(t *testing.T) {
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
				"document",
				"add",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad args",
			args: []string{
				"document",
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
				"document",
				"add",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Add duplicate",
			args: []string{
				"document",
				"add",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "bar"}))),
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

			err = afero.WriteFile(fs, "foo", []byte(getDocumentSkeleton()), 0644)
			assert.NoError(t, err, test.name+", assert document file creation")

			err = afero.WriteFile(fs, "bar", []byte(getDocumentSkeleton()), 0644)
			assert.NoError(t, err, test.name+", assert document file creation")

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

func TestCmdDocumentReplace(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldDocuments    []*domain.Document
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name:            "No args",
			args:            []string{"document", "replace"},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad args",
			args: []string{
				"document",
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
			oldDocuments: []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			args: []string{
				"document",
				"replace",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name:         "Some tags don't exist yet",
			oldDocuments: []*domain.Document{{ID: 1, Path: "foo"}},
			args: []string{
				"document",
				"replace",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar2"}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "All tags don't exist yet",
			args: []string{
				"document",
				"replace",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 3, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 4, Path: "bar2"}))),
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

			if test.oldDocuments != nil {
				cli.DocumentReplaceCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.DocumentManager.Add(context.Background(), test.oldDocuments)
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

func TestCmdDocumentUpsert(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldDocuments    []*domain.Document
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name:            "No args",
			args:            []string{"document", "upsert"},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad args",
			args: []string{
				"document",
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
			oldDocuments: []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			args: []string{
				"document",
				"upsert",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name:         "Updating and adding tags",
			oldDocuments: []*domain.Document{{ID: 1, Path: "foo"}},
			args: []string{
				"document",
				"upsert",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Updating tags",
			args: []string{
				"document",
				"upsert",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 3, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 4, Path: "bar2"}))),
			},
			outputValidator: testCommon.ValidatorEmpty,
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

			if test.oldDocuments != nil {
				cli.DocumentUpsertCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.DocumentManager.Add(context.Background(), test.oldDocuments)
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

func TestCmdDocumentEdit(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldDocuments    []*domain.Document
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "Only updater",
			args: []string{
				"document",
				"edit",
				"--updater",
				string(drop.From2To1(json.Marshal(domain.DocumentUpdater{Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Bad models",
			args: []string{
				"document",
				"edit",
				"foo",
				"bar",
				"--updater",
				string(drop.From2To1(json.Marshal(domain.DocumentUpdater{Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Minimal args, missing updater",
			args: []string{
				"document",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar2"}))),
			},
			errorMatcher:    testCommon.ValidatorContains("required", "updater"),
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("required"),
		},
		{
			name: "Minimal args, bad updater",
			args: []string{
				"document",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar2"}))),
				"--updater",
				"foo",
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name:         "Good minimal args",
			oldDocuments: []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			args: []string{
				"document",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar"}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.DocumentUpdater{Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			outputValidator: testCommon.ValidatorEqual("2\n"),
		},
		{
			name: "Bad filter",
			args: []string{
				"document",
				"edit",
				"--filter",
				"foo",
				"--updater",
				string(drop.From2To1(json.Marshal(domain.DocumentUpdater{Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             cmd.EntityMarshallingError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("marshalling"),
		},
		{
			name: "Filter and model",
			args: []string{
				"document",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.DocumentUpdater{Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             cmd.ConflictingPositionalArgsAndFlagError{"filter"},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("together with"),
		},
		{
			name:         "Good minimal args, filter instead of models",
			oldDocuments: []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			args: []string{
				"document",
				"edit",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.DocumentUpdater{Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			outputValidator: testCommon.ValidatorEqual("1\n"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags don't exist yet",
			args: []string{
				"document",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar2"}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.DocumentUpdater{Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "All tags don't exist yet",
			args: []string{
				"document",
				"edit",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 3, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 4, Path: "bar2"}))),
				"--updater",
				string(drop.From2To1(json.Marshal(domain.DocumentUpdater{Path: optional.Make(model.UpdateOperation[string]{Operator: model.UpdateAppend, Operand: "foo"})}))),
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

			if test.oldDocuments != nil {
				cli.DocumentEditCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.DocumentManager.Add(context.Background(), test.oldDocuments)
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

func TestCmdDocumentList(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tags            []*domain.Document
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No tags",
			args: []string{
				"document",
				"list",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Some tags",
			args: []string{
				"document",
				"list",
			},
			tags:            []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			outputValidator: testCommon.ValidatorContains("foo", "bar"),
		},
		{
			name: "Bad filter",
			args: []string{
				"document",
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
				"document",
				"list",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags, using filter",
			args: []string{
				"document",
				"list",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tags:            []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			outputValidator: testCommon.ValidatorContains("foo"),
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
				cli.DocumentListCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.DocumentManager.Add(context.Background(), test.tags)
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

func TestCmdDocumentRemove(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldDocuments    []*domain.Document
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "Bad models",
			args: []string{
				"document",
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
			oldDocuments: []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			args: []string{
				"document",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar"}))),
			},
			outputValidator: testCommon.ValidatorEqual("2\n"),
		},
		{
			name: "Bad filter",
			args: []string{
				"document",
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
				"document",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             cmd.ConflictingPositionalArgsAndFlagError{"filter"},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("together with"),
		},
		{
			name:         "Good minimal args, filter instead of models",
			oldDocuments: []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			args: []string{
				"document",
				"remove",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			outputValidator: testCommon.ValidatorEqual("1\n"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags don't exist yet",
			args: []string{
				"document",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 2, Path: "bar2"}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("does not exist"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "All tags don't exist yet",
			args: []string{
				"document",
				"remove",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 3, Path: "foo2"}))),
				string(drop.From2To1(json.Marshal(domain.Document{ID: 4, Path: "bar2"}))),
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

			if test.oldDocuments != nil {
				cli.DocumentRemoveCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.DocumentManager.Add(context.Background(), test.oldDocuments)
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

func TestCmdDocumentFind(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tags            []*domain.Document
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "Bad filter",
			args: []string{
				"document",
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
				"document",
				"find-first",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags",
			args: []string{
				"document",
				"find-first",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tags:            []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			outputValidator: testCommon.ValidatorContains("foo"),
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
				cli.DocumentFindCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.DocumentManager.Add(context.Background(), test.tags)
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

func TestCmdDocumentCount(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tags            []*domain.Document
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No tags, no filter",
			args: []string{
				"document",
				"count",
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.Count{0}))) + "\n"),
		},
		{
			name: "Some tags, no filter",
			args: []string{
				"document",
				"count",
			},
			tags:            []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.Count{2}))) + "\n"),
		},
		{
			name: "Bad filter",
			args: []string{
				"document",
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
				"document",
				"count",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.Count{0}))) + "\n"),
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags, using filter",
			args: []string{
				"document",
				"count",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tags:            []*domain.Document{{ID: 1, Path: "foo"}, {ID: 2, Path: "bar"}},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.Count{1}))) + "\n"),
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
				cli.DocumentCountCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.DocumentManager.Add(context.Background(), test.tags)
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

func TestCmdDocumentDoesExist(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		args            []string
		tag             *domain.Document
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No tags",
			args: []string{
				"document",
				"does-exist",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Some tags",
			args: []string{
				"document",
				"does-exist",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo"}))),
			},
			tag:             &domain.Document{ID: 1, Path: "foo"},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.DoesExist{true}))) + "\n"),
		},
		{
			name: "Bad filter",
			args: []string{
				"document",
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
				"document",
				"does-exist",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.DoesExist{false}))) + "\n"),
		},

		// FIX: This actually tests the manager, not the cli
		{
			name: "Some tags, using filter",
			args: []string{
				"document",
				"does-exist",
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
			},
			tag:             &domain.Document{ID: 1, Path: "foo"},
			outputValidator: testCommon.ValidatorEqual(string(drop.From2To1(json.Marshal(cmd.DoesExist{true}))) + "\n"),
		},
		{
			name: "Bad model",
			args: []string{
				"document",
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
				"document",
				"does-exist",
				string(drop.From2To1(json.Marshal(domain.Document{ID: 1, Path: "foo"}))),
				"--filter",
				string(drop.From2To1(json.Marshal(domain.DocumentFilter{Path: optional.Make(model.FilterOperation[string]{Operator: model.FilterEqual, Operand: model.ScalarOperand[string]{Operand: "foo"}})}))),
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
				cli.DocumentDoesExistCmd.PreRun = func(_ *cobra.Command, _ []string) {
					err = cli.BNTPBackend.DocumentManager.Add(context.Background(), []*domain.Document{test.tag})
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
