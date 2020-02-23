package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

type commandInterface interface {
	Execute() error
}

func ExitWithError(cmd string, err error) {
	fmt.Fprintf(os.Stdout, color.RedString("\n❗️ %v\n", err))
	fmt.Fprintf(os.Stdout, color.RedString("❗️ %s command has failed\n", cmd))
	os.Exit(1)
}
