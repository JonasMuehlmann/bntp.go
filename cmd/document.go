package cmd

import (
	"github.com/spf13/cobra"
)

var documentCmd = &cobra.Command{
	Use:   "document",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	rootCmd.AddCommand(documentCmd)

}
