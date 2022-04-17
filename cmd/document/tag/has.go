package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hasCmd = &cobra.Command{
	Use:   "has",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	tagCmd.AddCommand(hasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
