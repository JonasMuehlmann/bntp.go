package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var documentTypeCmd = &cobra.Command{
	Use:   "type",
	Short: "Manage types of bntp documents",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentTypeAddCmd = &cobra.Command{
	Use:   "add TYPE...",
	Short: "Add new bntp document types",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentTypeEditCmd = &cobra.Command{
	Use:   "edit OLD_NAME NEW_NAME",
	Short: "Change a bntp document type",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentTypeRemoveCmd = &cobra.Command{
	Use:   "remove TYPE...",
	Short: "Remove bntp document types",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}
