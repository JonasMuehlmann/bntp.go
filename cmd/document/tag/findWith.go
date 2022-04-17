package cmd

import (
	"github.com/spf13/cobra"
)

var findWithCmd = &cobra.Command{
	Use:   "findWith",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	tagCmd.AddCommand(findWithCmd)

}
