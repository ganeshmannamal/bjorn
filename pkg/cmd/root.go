package cmd

import (
	goflag "flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/golang/glog"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "bjorn",
	Short: "image comparison tool for Bjorn",
	Long: `bjorn allow the users (like Bjorn) to compare images that are
		provided as a list in a csv file`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// For cobra + glog flags. Available to all subcommands.
		goflag.Parse()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bjorn.yaml)")

	rootCmd.AddCommand(NewDiffCommand())
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stdout, color.RedString("❗️ %v\n", err))
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			glog.Error(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bjorn" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bjorn")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		glog.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
