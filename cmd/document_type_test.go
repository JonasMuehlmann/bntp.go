package cmd_test

import (
	"context"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/cmd"
	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	testCommon "github.com/JonasMuehlmann/bntp.go/test"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCmdDocumentTypeAdd(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldTypes        []string
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "no args",
			args: []string{
				"document",
				"type",
				"add",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name: "Good minimal arg",
			args: []string{
				"document",
				"type",
				"add",
				"foo",
				"bar",
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name:     "Add duplicate",
			oldTypes: []string{"bar"},
			args: []string{
				"document",
				"type",
				"add",
				"foo",
				"bar",
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

			if test.oldTypes != nil {
				err = cli.BNTPBackend.DocumentManager.AddType(context.Background(), test.oldTypes)
				assert.NoError(t, err, test.name+", assert adding old types")
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

func TestCmdDocumentTypeEdit(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldTypes        []string
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "No args",
			args: []string{
				"document",
				"type",
				"edit",
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("arg", "received"),
			errorMatcher:    testCommon.ValidatorContains("arg", "received"),
		},
		{
			name: "No args",
			args: []string{
				"document",
				"type",
				"edit",
				"foo",
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("arg", "received"),
			errorMatcher:    testCommon.ValidatorContains("arg", "received"),
		},
		{
			name:     "Good minimal args",
			oldTypes: []string{"foo"},
			args: []string{
				"document",
				"type",
				"edit",
				"foo",
				"bar",
			},
			outputValidator: testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Old does not exist",
			args: []string{
				"document",
				"type",
				"edit",
				"foo",
				"bar",
			},
			err:             helper.NonExistentPrimaryDataError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("primary"),
		},
		// FIX: This actually tests the manager, not the cli
		{
			name:     "Renaming to existing",
			oldTypes: []string{"foo", "bar"},
			args: []string{
				"document",
				"type",
				"edit",
				"foo",
				"bar",
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

			if test.oldTypes != nil {
				err = cli.BNTPBackend.DocumentManager.AddType(context.Background(), test.oldTypes)
				assert.NoError(t, err, test.name+", assert adding old types")
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

func TestCmdDocumentTypeRemove(t *testing.T) {
	tests := []struct {
		err             error
		errorMatcher    testCommon.OutputValidator
		name            string
		oldTypes        []string
		args            []string
		outputValidator testCommon.OutputValidator
		errorValidator  testCommon.OutputValidator
	}{
		{
			name: "no args",
			args: []string{
				"document",
				"type",
				"remove",
			},
			err:             helper.IneffectiveOperationError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("no effect"),
		},
		{
			name:     "Good minimal arg",
			oldTypes: []string{"foo", "bar"},
			args: []string{
				"document",
				"type",
				"remove",
				"foo",
				"bar",
			},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorEmpty,
		},
		// FIX: This actually tests the manager, not the cli
		{
			name: "Remove non-existent",
			args: []string{
				"document",
				"type",
				"remove",
				"foo",
				"bar",
			},
			err:             helper.NonExistentPrimaryDataError{},
			outputValidator: testCommon.ValidatorEmpty,
			errorValidator:  testCommon.ValidatorContains("primary"),
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

			if test.oldTypes != nil {
				err = cli.BNTPBackend.DocumentManager.AddType(context.Background(), test.oldTypes)
				assert.NoError(t, err, test.name+", assert adding old types")
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
