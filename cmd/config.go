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
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helper"
	"github.com/JonasMuehlmann/datastructures.go/maps/hashmap"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func WithConfigCommand() CliOption {
	return func(cli *Cli) (err error) {

		cli.ConfigCmd = &cobra.Command{
			Use:   "config",
			Short: "Manage bntp configuration",
			Long:  `A longer description`,
			RunE: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					return helper.IneffectiveOperationError{Inner: helper.EmptyInputError{}}
				}

				return nil
			},
		}
		cli.ConfigPathsCmd = &cobra.Command{
			Use:   "paths",
			Short: "List the search paths for config files",
			Long:  `A longer description`,
			RunE: func(cmd *cobra.Command, args []string) error {
				for _, extension := range cli.ConfigManager.ConfigSearchPaths {
					fmt.Fprintln(cli.RootCmd.OutOrStdout(), extension)
				}

				return nil
			},
		}

		cli.ConfigExtensionsCmd = &cobra.Command{
			Use:   "extensions",
			Short: "List the allowed extensions for config files",
			Long:  `A longer description`,
			RunE: func(cmd *cobra.Command, args []string) error {
				for _, extension := range viper.SupportedExts {
					fmt.Fprintln(cli.RootCmd.OutOrStdout(), extension)
				}

				return nil
			},
		}

		cli.configBaseNameCmd = &cobra.Command{
			Use:   "base-name",
			Short: "Show the base name expected for config files",
			Long:  `A longer description`,
			RunE: func(cmd *cobra.Command, args []string) error {
				fmt.Fprintln(cli.RootCmd.OutOrStdout(), cli.ConfigManager.ConfigFileBaseName)

				return nil
			},
		}

		cli.exportConfigCmd = &cobra.Command{
			Use:   "export FILE",
			Short: "Export the current config state",
			Long:  `A longer description`,
			Args:  cobra.ExactArgs(1),
			RunE: func(cmd *cobra.Command, args []string) error {
				err := cli.ConfigManager.Viper.WriteConfigAs(args[0])

				return err
			},
		}

		cli.RootCmd.AddCommand(cli.ConfigCmd)
		cli.ConfigCmd.AddCommand(cli.ConfigPathsCmd)
		cli.ConfigCmd.AddCommand(cli.ConfigExtensionsCmd)
		cli.ConfigCmd.AddCommand(cli.configBaseNameCmd)
		cli.ConfigCmd.AddCommand(cli.exportConfigCmd)
	}
}
