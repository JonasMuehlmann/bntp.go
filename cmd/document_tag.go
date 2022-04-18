package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var documentTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags of bntp documents",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentTagAddCmd = &cobra.Command{
	Use:   "add TAG...",
	Short: "Add tags to a bntp document",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentTagEditCmd = &cobra.Command{
	Use:   "edit OLD_NAME NEW_NAME",
	Short: "Change a tag in a bntp document",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentTagFindWithCmd = &cobra.Command{
	Use:   "find-with TAG...",
	Short: "Find bntp documents with specific tags",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentTagHasCmd = &cobra.Command{
	Use:   "has TAG...",
	Short: "Check if a bntp document has specific tags",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentTagListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tags of a bntp document",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentTagRemoveCmd = &cobra.Command{
	Use:   "remove TAG...",
	Short: "Remove a tag from a bntp document",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}
