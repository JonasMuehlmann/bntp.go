package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	tagCmd.AddCommand(addCmd)

}
