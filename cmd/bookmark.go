package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "A brief description of your command",
	Long:  `A longeer description`,
}

func init() {
	rootCmd.AddCommand(bookmarkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bookmarkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bookmarkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
