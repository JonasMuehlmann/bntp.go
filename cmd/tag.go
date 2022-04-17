package cmd

import (
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	rootCmd.AddCommand(tagCmd)

}
