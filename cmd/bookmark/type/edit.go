package cmd

import (
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	typeCmd.AddCommand(editCmd)

}
