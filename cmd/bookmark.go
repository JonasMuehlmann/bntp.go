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
