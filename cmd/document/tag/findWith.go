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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findWithCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findWithCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
