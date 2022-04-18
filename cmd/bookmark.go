package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var bookmarkAddCmd = &cobra.Command{
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

var bookmarkEditCmd = &cobra.Command{
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

var bookmarkExportCmd = &cobra.Command{
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

var bookmarkImportCmd = &cobra.Command{
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

var bookmarkListCmd = &cobra.Command{
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

var bookmarkRemoveCmd = &cobra.Command{
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
}
