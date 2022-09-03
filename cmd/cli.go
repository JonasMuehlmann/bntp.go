// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package cmd

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/bntp/backend"
	"github.com/JonasMuehlmann/bntp.go/internal/config"
	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CliOption func(*Cli)

func WithAll() CliOption {
	return func(cli *Cli) {
		WithBookmarkCommand()(cli)
		WithBookmarkTypeCommand()(cli)
		WithDocumentCommand()(cli)
		WithDocumentTypeCommand()(cli)
		WithTagCommand()(cli)
		WithConfigCommand()(cli)
		WithConfigManager()(cli)
		WithBNTPBackend()(cli)
	}
}

func WithBNTPBackend() CliOption {
	return func(cli *Cli) {
		cli.BNTPBackend = cli.ConfigManager.NewBackendFromConfig()
	}
}

func WithConfigManager() CliOption {
	return func(cli *Cli) {
		cli.ConfigManager = config.NewConfigManager(cli.StdErr, cli.DBOverride, cli.FsOverride)
	}
}

func WithStdErrOverride(stderrToUse io.Writer) CliOption {
	return func(cli *Cli) {
		cli.StdErr = stderrToUse
	}
}

func WithDbOverride(dbToUse *sql.DB) CliOption {
	return func(cli *Cli) {
		cli.DBOverride = dbToUse
	}
}

func WithFsOverride(fsToUse afero.Fs) CliOption {
	return func(cli *Cli) {
		cli.FsOverride = fsToUse
		cli.Fs = fsToUse
	}
}
func NewCli(options ...CliOption) *Cli {
	cli := &Cli{
		StdErr: os.Stderr,
		Fs:     afero.NewOsFs(),
	}

	cli.RootCmd = &cobra.Command{
		Use:   "bntp.go",
		Short: "bntp.go - the all in one productivity system.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
			}

			return nil
		},
	}

	// Make sure e.g. db is set before being used
	for _, option := range options {
		optionName := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if !strings.Contains(optionName, "Command") {
			option(cli)
		}
	}

	for _, option := range options {
		optionName := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if strings.Contains(optionName, "Command") {
			option(cli)
		}
	}

	cli.RootCmd.PersistentFlags().StringVarP(
		&cli.ConfigManager.PassedConfigPath,
		"config",
		"c",
		"",
		"The config file to use instead of ones found in search paths",
	)

	cli.RootCmd.PersistentFlags().BoolVar(
		&cli.DebugMode,
		"debug",
		false,
		"Whetever to enable debug mode for extra logic/logging",
	)

	cli.RootCmd.PersistentPreRun = cli.PersistentPreRun
	cli.RootCmd.PersistentFlags().String(
		strcase.KebabCase(config.ConsoleLogLevel),
		cli.ConfigManager.GetDefaultSettings()[config.ConsoleLogLevel].(string),
		fmt.Sprintf("The minimum log level to display on the console (Allowed values: %v)", log.AllLevels),
	)

	cli.RootCmd.PersistentFlags().String(
		strcase.KebabCase(config.FileLogLevel),
		cli.ConfigManager.GetDefaultSettings()[config.FileLogLevel].(string),
		fmt.Sprintf("The minimum log level to log to the log files (Allowed values: %v)", log.AllLevels),
	)

	for _, setting := range []string{config.ConsoleLogLevel, config.FileLogLevel} {
		err := cli.ConfigManager.Viper.BindPFlag(setting, cli.RootCmd.PersistentFlags().Lookup(strcase.KebabCase(setting)))
		if err != nil {
			panic(err)
		}
	}

	cli.Logger = cli.ConfigManager.Logger

	return cli
}

func (cli *Cli) Execute() error {
	cli.RootCmd.SilenceUsage = true
	cli.RootCmd.SilenceErrors = true

	err := cli.RootCmd.Execute()

	if err != nil {
		cli.Logger.Error(err)
	}

	return err
}

func (cli *Cli) PersistentPreRun(cmd *cobra.Command, args []string) {
	if cli.DebugMode {
		boil.DebugMode = true
	}
}

type Cli struct {
	ConfigManager *config.ConfigManager
	BNTPBackend   *backend.Backend
	InFormat      string
	OutFormat     string
	FilterRaw     string
	UpdaterRaw    string
	PathFormat    bool
	ShortFormat   bool
	DebugMode     bool
	StdErr        io.Writer
	DBOverride    *sql.DB
	Logger        *log.Logger
	FsOverride    afero.Fs
	Fs            afero.Fs

	BookmarkAddCmd        *cobra.Command
	BookmarkCmd           *cobra.Command
	BookmarkCountCmd      *cobra.Command
	BookmarkDoesExistCmd  *cobra.Command
	BookmarkEditCmd       *cobra.Command
	BookmarkFindCmd       *cobra.Command
	BookmarkListCmd       *cobra.Command
	BookmarkRemoveCmd     *cobra.Command
	BookmarkReplaceCmd    *cobra.Command
	BookmarkTypeAddCmd    *cobra.Command
	BookmarkTypeCmd       *cobra.Command
	BookmarkTypeEditCmd   *cobra.Command
	BookmarkTypeRemoveCmd *cobra.Command
	BookmarkTypeListCmd   *cobra.Command
	BookmarkUpsertCmd     *cobra.Command
	configBaseNameCmd     *cobra.Command
	ConfigCmd             *cobra.Command
	ConfigExtensionsCmd   *cobra.Command
	ConfigPathsCmd        *cobra.Command
	DocumentAddCmd        *cobra.Command
	DocumentCmd           *cobra.Command
	DocumentCountCmd      *cobra.Command
	DocumentDoesExistCmd  *cobra.Command
	DocumentEditCmd       *cobra.Command
	DocumentFindCmd       *cobra.Command
	DocumentListCmd       *cobra.Command
	DocumentRemoveCmd     *cobra.Command
	DocumentReplaceCmd    *cobra.Command
	DocumentTypeAddCmd    *cobra.Command
	DocumentTypeCmd       *cobra.Command
	DocumentTypeEditCmd   *cobra.Command
	DocumentTypeRemoveCmd *cobra.Command
	DocumentTypeListCmd   *cobra.Command
	DocumentUpsertCmd     *cobra.Command
	exportConfigCmd       *cobra.Command
	RootCmd               *cobra.Command
	TagAddCmd             *cobra.Command
	TagAmbiguousCmd       *cobra.Command
	TagCmd                *cobra.Command
	TagCountCmd           *cobra.Command
	TagDoesExistCmd       *cobra.Command
	TagEditCmd            *cobra.Command
	TagExportCmd          *cobra.Command
	TagFindCmd            *cobra.Command
	TagImportCmd          *cobra.Command
	TagListCmd            *cobra.Command
	TagRemoveCmd          *cobra.Command
	TagReplaceCmd         *cobra.Command
	TagShortCmd           *cobra.Command
	TagUpsertCmd          *cobra.Command
}
