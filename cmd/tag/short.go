package cmd

import (
	"github.com/spf13/cobra"
)

var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	tagCmd.AddCommand(shortCmd)

}
