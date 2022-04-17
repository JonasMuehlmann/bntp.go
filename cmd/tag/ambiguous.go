package cmd

import (
	"github.com/spf13/cobra"
)

var ambiguousCmd = &cobra.Command{
	Use:   "ambiguous",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	tagCmd.AddCommand(ambiguousCmd)

}
