package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	RootCmd = &cobra.Command{
		Use:   "bntp.go",
		Short: "bntp.go - the all in one productivity system.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// TODO: Mark quiet and verbose flags as mutually exclusive when the next cobra version gets released.
	RootCmd.PersistentFlags().BoolP("quiet", "q", false, "Disable all logging")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable full logging")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/bntp/")
		viper.AddConfigPath("$HOME/.config/")
		viper.AddConfigPath("$HOME/")
		viper.SetConfigName("bntprc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Fprintf(os.Stderr, "%v, allowed extensions/formats: %v\n", err.Error(), viper.SupportedExts)
		os.Exit(0)
	}
}
