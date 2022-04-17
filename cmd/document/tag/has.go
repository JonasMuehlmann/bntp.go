package cmd

import (
	"github.com/spf13/cobra"
)

var hasCmd = &cobra.Command{
	Use:   "has",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	tagCmd.AddCommand(hasCmd)

}
