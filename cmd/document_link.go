package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var documentLinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Manage links between bntp documents",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentLinkEditCmd = &cobra.Command{
	Use:   "edit OLD_NAME NEW_NAME",
	Short: "Change a link between bntp documents",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentLinkRemoveCmd = &cobra.Command{
	Use:   "remove LINK...",
	Short: "Remove links from a bntp documents",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentLinkAddCmd = &cobra.Command{
	Use:   "add LINK...",
	Short: "Add links to a bntp document",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentLinkListCmd = &cobra.Command{
	Use:   "list",
	Short: "List links from/to a bntp document",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}
