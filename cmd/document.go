package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var documentCmd = &cobra.Command{
	Use:   "document",
	Short: "Manage bntp documents",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentAddCmd = &cobra.Command{
	Use:   "add FILE...",
	Short: "Add bntp documents",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentEditCmd = &cobra.Command{
	Use:   "edit OLD_NAME NEW_NAME",
	Short: "Edit a bntp document",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List bntp documents",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentRemoveCmd = &cobra.Command{
	Use:   "remove FILE...",
	Short: "Remove bntp documents",
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
	RootCmd.AddCommand(documentCmd)

	documentCmd.AddCommand(documentAddCmd)
	documentCmd.AddCommand(documentEditCmd)
	documentCmd.AddCommand(documentListCmd)
	documentCmd.AddCommand(documentRemoveCmd)

	documentCmd.AddCommand(documentTypeCmd)
	documentTypeCmd.AddCommand(documentTypeAddCmd)
	documentTypeCmd.AddCommand(documentTypeEditCmd)
	documentTypeCmd.AddCommand(documentTypeRemoveCmd)

	documentCmd.AddCommand(documentTagCmd)
	documentTagCmd.AddCommand(documentTagAddCmd)
	documentTagCmd.AddCommand(documentTagEditCmd)
	documentTagCmd.AddCommand(documentTagFindWithCmd)
	documentTagCmd.AddCommand(documentTagHasCmd)
	documentTagCmd.AddCommand(documentTagListCmd)
	documentTagCmd.AddCommand(documentTagRemoveCmd)

	documentCmd.AddCommand(documentLinkCmd)
	documentLinkCmd.AddCommand(documentLinkAddCmd)
	documentLinkCmd.AddCommand(documentLinkEditCmd)
	documentLinkCmd.AddCommand(documentLinkRemoveCmd)
	documentLinkCmd.AddCommand(documentLinkListCmd)

	for _, subcommand := range documentCmd.Commands() {
		if !slices.Contains([]*cobra.Command{documentAddCmd, documentListCmd, documentRemoveCmd}, subcommand) {
			subcommand.PersistentFlags().StringP("document", "d", "", "The document to work with")
		}
	}

	documentLinkListCmd.Flags().BoolP("back", "b", false, "Whetever to list backlinks instead of links")
	documentTagHasCmd.Flags().BoolP("or", "o", false, "Whetever to require any instead of all tags to match")
	documentTagFindWithCmd.Flags().BoolP("or", "o", false, "Whetever to require any instead of all tags to match")
}
