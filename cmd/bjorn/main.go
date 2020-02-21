package main

import (
	"fmt"
	"os"
	"strings"
)

// Set by build process
var (
	version string
)

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

func Execute() {}
