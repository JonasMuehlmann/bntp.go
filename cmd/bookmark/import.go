package cmd

import (
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	bookmarkCmd.AddCommand(importCmd)

}
