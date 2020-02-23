package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/ganeshmannamal/bjorn/pkg/cmd"
	"os"
	"strings"
)

// Set by build process
var (
	version string
)

//func init() {
//	rootCmd = cmd.NewRootCommand()
//	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
//}

func main() {
	// Look for version
	for _, v := range os.Args[1:] {
		v = strings.TrimLeft(v, "-")
		if v == "v" || v == "version" {
			if version == "" {
				version = "dev"
			}

			fmt.Printf("bjorn %s\n", version)
			os.Exit(0)
		}
	}
	Execute()
}

func Execute() {
	rootCmd := cmd.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stdout, color.RedString("❗️ %v\n", err))
		os.Exit(1)
	}
}
