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

	"github.com/JonasMuehlmann/bntp.go/bntp/backend"
	"github.com/JonasMuehlmann/bntp.go/internal/config"
	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stoewer/go-strcase"
)

var BNTPBackend *backend.Backend
var stderrToUse io.Writer
var dbToUse *sql.DB

func WithAll(cli *Cli) {
	WithBookmark(cli)
	WithBookmarkType(cli)
	WithDocument(cli)
	WithDocumentType(cli)
	WithTag(cli)
	WithConfig(cli)
}

func NewCli(options ...func(*Cli)) *Cli {
	cli := &Cli{}
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

	cli.RootCmd.PersistentFlags().StringVarP(&config.PassedConfigPath, "config", "c", "", "The config file to use instead of ones found in search paths")

	cli.RootCmd.PersistentFlags().String(
		strcase.KebabCase(config.ConsoleLogLevel),
		config.GetDefaultSettings()[config.ConsoleLogLevel].(string),
		fmt.Sprintf("The minimum log level to display on the console (Allowed values: %v)", log.AllLevels),
	)

	cli.RootCmd.PersistentFlags().String(
		strcase.KebabCase(config.FileLogLevel),
		config.GetDefaultSettings()[config.FileLogLevel].(string),
		fmt.Sprintf("The minimum log level to log to the log files (Allowed values: %v)", log.AllLevels),
	)

	for _, option := range options {
		option(cli)
	}

	cobra.OnInitialize(func() { config.InitConfig(stderrToUse, dbToUse); BNTPBackend = config.NewBackendFromConfig() }, cli.bindFlagsToConfig)

	return cli
}

func (cli *Cli) Execute(backend *backend.Backend, stderr io.Writer, testDB ...*sql.DB) error {
	BNTPBackend = backend

	cli.RootCmd.SilenceUsage = true
	cli.RootCmd.SilenceErrors = true

	stderrToUse = stderr
	if len(testDB) > 0 {
		dbToUse = testDB[0]
	}

	err := cli.RootCmd.Execute()

	if err != nil {
		log.Error(err)
	}

	return err
}
func (cli *Cli) bindFlagsToConfig() {
	for _, setting := range []string{config.ConsoleLogLevel, config.FileLogLevel} {
		err := viper.BindPFlag(setting, cli.RootCmd.Flags().Lookup(strcase.KebabCase(setting)))
		if err != nil {
			log.Fatal(err)
		}
	}
}

type Cli struct {
	Format     string
	FilterRaw  string
	UpdaterRaw string

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
	BookmarkUpsertCmd     *cobra.Command
	configBaseNameCmd     *cobra.Command
	configCmd             *cobra.Command
	configExtensionsCmd   *cobra.Command
	configPathsCmd        *cobra.Command
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
