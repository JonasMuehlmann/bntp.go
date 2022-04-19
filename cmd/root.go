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
	"log"
	"os"

	"github.com/JonasMuehlmann/bntp.go/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "bntp.go",
	Short: "bntp.go - the all in one productivity system.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// TODO: Mark quiet and verbose flags as mutually exclusive when the next cobra version gets released.
	RootCmd.PersistentFlags().BoolP(config.QUIET, config.QUIET[0:1], false, "Disable all logging")
	RootCmd.PersistentFlags().BoolP(config.VERBOSE, config.VERBOSE[0:1], false, "Enable full logging")
	RootCmd.PersistentFlags().StringVarP(&config.ConfigPath, "config", "c", "", "The config file to use instead of ones found in search paths")

	cobra.OnInitialize(config.InitConfig, bindFlagsToConfig)
}

func bindFlagsToConfig() {
	for _, setting := range []string{config.QUIET, config.VERBOSE} {
		err := viper.BindPFlag(setting, RootCmd.Flags().Lookup(setting))
		if err != nil {
			log.Fatal(err)
		}
	}
}
