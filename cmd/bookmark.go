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
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Manage bntp bookmarks",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkAddCmd = &cobra.Command{
	Use:   "add DATA",
	Short: "Add a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkEditCmd = &cobra.Command{
	Use:   "edit NEW_DATA",
	Short: "Edit a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkExportCmd = &cobra.Command{
	Use:   "export PATH",
	Short: "Export bntp bookmarks",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkImportCmd = &cobra.Command{
	Use:   "import PATH",
	Short: "Import bntp bookmarks",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkListCmd = &cobra.Command{
	Use:   "list",
	Short: "List bntp bookmarks",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkRemoveCmd = &cobra.Command{
	Use:   "remove BOOKMARK...",
	Short: "Remove a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	RootCmd.AddCommand(bookmarkCmd)
	bookmarkCmd.AddCommand(bookmarkListCmd)
	bookmarkCmd.AddCommand(bookmarkImportCmd)
	bookmarkCmd.AddCommand(bookmarkExportCmd)
	bookmarkCmd.AddCommand(bookmarkEditCmd)
	bookmarkCmd.AddCommand(bookmarkAddCmd)
	bookmarkCmd.AddCommand(bookmarkRemoveCmd)

	bookmarkCmd.AddCommand(bookmarkTagCmd)
	bookmarkTagCmd.AddCommand(bookmarkTagRemoveCmd)
	bookmarkTagCmd.AddCommand(bookmarkTagEditCmd)
	bookmarkTagCmd.AddCommand(bookmarkTagAddCmd)
	bookmarkTagCmd.AddCommand(bookmarkTagList)

	bookmarkCmd.AddCommand(bookmarkTypeCmd)
	bookmarkTypeCmd.AddCommand(bookmarkTypeAddCmd)
	bookmarkTypeCmd.AddCommand(bookmarkTypeEditCmd)
	bookmarkTypeCmd.AddCommand(bookmarkTypeRemoveCmd)

	for _, subcommand := range bookmarkCmd.Commands() {
		if !slices.Contains([]*cobra.Command{bookmarkAddCmd, bookmarkListCmd, bookmarkRemoveCmd, bookmarkExportCmd, bookmarkImportCmd}, subcommand) {
			subcommand.PersistentFlags().StringP("bookmark", "b", "", "The bookmark to work with")
		}
	}
}
