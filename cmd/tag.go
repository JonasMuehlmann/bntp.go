package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var addCmd = &cobra.Command{
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

var ambiguousCmd = &cobra.Command{
	Use:   "ambiguous",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var editCmd = &cobra.Command{
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

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var listCmd = &cobra.Command{
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

var removeCmd = &cobra.Command{
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

var shortCmd = &cobra.Command{
	Use:   "short",
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
	RootCmd.AddCommand(tagCmd)

	tagCmd.AddCommand(shortCmd)
	tagCmd.AddCommand(removeCmd)
	tagCmd.AddCommand(listCmd)
	tagCmd.AddCommand(importCmd)
	tagCmd.AddCommand(exportCmd)
	tagCmd.AddCommand(editCmd)
	tagCmd.AddCommand(ambiguousCmd)
	tagCmd.AddCommand(addCmd)
}
