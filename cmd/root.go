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

var RootCmd = &cobra.Command{
	Use:   "bntp.go",
	Short: "bntp.go - the all in one productivity system.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return helper.IneffectiveOperationError{Inner: helper.EmptyInputError}
		}

		return nil
	},
}

func Execute(backend *backend.Backend, stderr io.Writer, testDB ...*sql.DB) error {
	BNTPBackend = backend

	RootCmd.SilenceUsage = true
	RootCmd.SilenceErrors = true

	stderrToUse = stderr
	if len(testDB) > 0 {
		dbToUse = testDB[0]
	}

	err := RootCmd.Execute()

	if err != nil {
		log.Error(err)
	}

	return err
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&config.PassedConfigPath, "config", "c", "", "The config file to use instead of ones found in search paths")

	RootCmd.PersistentFlags().String(
		strcase.KebabCase(config.ConsoleLogLevel),
		config.GetDefaultSettings()[config.ConsoleLogLevel].(string),
		fmt.Sprintf("The minimum log level to display on the console (Allowed values: %v)", log.AllLevels),
	)

	RootCmd.PersistentFlags().String(
		strcase.KebabCase(config.FileLogLevel),
		config.GetDefaultSettings()[config.FileLogLevel].(string),
		fmt.Sprintf("The minimum log level to log to the log files (Allowed values: %v)", log.AllLevels),
	)

	cobra.OnInitialize(func() { config.InitConfig(stderrToUse, dbToUse); BNTPBackend = config.NewBackendFromConfig() }, bindFlagsToConfig)
}

func bindFlagsToConfig() {
	for _, setting := range []string{config.ConsoleLogLevel, config.FileLogLevel} {
		err := viper.BindPFlag(setting, RootCmd.Flags().Lookup(strcase.KebabCase(setting)))
		if err != nil {
			log.Fatal(err)
		}
	}
}
