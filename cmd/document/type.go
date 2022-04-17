package cmd

import (
	"github.com/spf13/cobra"
)

var typeCmd = &cobra.Command{
	Use:   "type",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	documentCmd.AddCommand(typeCmd)

}
