package cmd

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	typeCmd.AddCommand(removeCmd)

}
