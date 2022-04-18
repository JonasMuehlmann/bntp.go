package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var bookmarkTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags of bntp bookmarks",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkTagAddCmd = &cobra.Command{
	Use:   "add TAG...",
	Short: "Add tags to a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkTagEditCmd = &cobra.Command{
	Use:   "edit OLD_NAME NEW_NAME",
	Short: "Edit tags of a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkTagRemoveCmd = &cobra.Command{
	Use:   "remove TAG...",
	Short: "Remove tags from a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkTagList = &cobra.Command{
	Use:   "list",
	Short: "List tags of a bntp bookmark",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}
