package cmd

import (
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	documentCmd.AddCommand(linkCmd)

}
