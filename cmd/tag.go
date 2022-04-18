package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags available for bntp entities",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var tagAddCmd = &cobra.Command{
	Use:   "add TAG...",
	Short: "Add new bntp tags",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

// TODO: If ambiguous, return ambiguous component
var tagAmbiguousCmd = &cobra.Command{
	Use:   "ambiguous TAG...",
	Short: "Check if bntp tag's leafs are ambiguous",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var tagEditCmd = &cobra.Command{
	Use:   "edit OLD_NAME NEW_NAME",
	Short: "Change a bntp tag",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var tagExportCmd = &cobra.Command{
	Use:   "export FILE",
	Short: "Export bntp tags",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var tagImportCmd = &cobra.Command{
	Use:   "import FILE",
	Short: "Import bntp tags",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var tagListCmd = &cobra.Command{
	Use:   "list",
	Short: "List bntp tags",
	Long:  `A longer description`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var tagRemoveCmd = &cobra.Command{
	Use:   "remove TAG...",
	Short: "Remove bntp tags",
	Long:  `A longer description`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var tagShortCmd = &cobra.Command{
	Use:   "short TAG...",
	Short: "Return shortened bntp tags",
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
	RootCmd.AddCommand(tagCmd)

	tagCmd.AddCommand(tagShortCmd)
	tagCmd.AddCommand(tagRemoveCmd)
	tagCmd.AddCommand(tagListCmd)
	tagCmd.AddCommand(tagImportCmd)
	tagCmd.AddCommand(tagExportCmd)
	tagCmd.AddCommand(tagEditCmd)
	tagCmd.AddCommand(tagAmbiguousCmd)
	tagCmd.AddCommand(tagAddCmd)

	tagListCmd.Flags().BoolP("short", "s", false, "Whetever to list shortened tags instead of fully qualified ones")
}
