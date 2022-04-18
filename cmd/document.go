package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var documentCmd = &cobra.Command{
	Use:   "document",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentAddCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var documentRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long:  `A longer description`,
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
}
