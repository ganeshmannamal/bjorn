package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"github.com/ganeshmannamal/bjorn/pkg/pair"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type Opts struct {
	csvFile string
	outFile string
}

func (opts *Opts) Run() error {

	csvRootPath, err := filepath.Abs(filepath.Dir(opts.csvFile))
	if err != nil {
		return err
	}

	// Open CSV file
	f, err := os.Open(opts.csvFile)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	if opts.outFile == "" {
		opts.outFile = filepath.Join(csvRootPath, "output.csv")
	}
	out, err := os.Create(opts.outFile)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(out)
	defer writer.Flush()

	for _, line := range lines {
		img1Path := resolvePath(csvRootPath, line[0])
		img2Path := resolvePath(csvRootPath, line[1])

		p, err := pair.NewImagePair(img1Path, img2Path)

		if err != nil{
			return err
		}

		p.Compare()
		err = writer.Write([]string{line[0], line[1], fmt.Sprintf("%.2f", p.Score), fmt.Sprintf("%f", p.Time)})
	}
	return nil
}

func NewRootCommand() *cobra.Command {
	opts := &Opts{}
	rootCmd := &cobra.Command{
		Use:   "bjorn",
		Short: "image comparison tool for Bjorn",
		Long: `Bjorn allow the users (Bjorn) to compare images that are
		provided as a list in a csv file`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("need atleast one argument for csv file")
			}

			opts.csvFile = args[0]

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			err := opts.Run()
			if err != nil {
				ExitWithError("", err)
			}
			color.New(color.Bold, color.FgGreen).Fprintf(os.Stdout, "\nCommand has completed\n")
		},
	}

	return rootCmd
}

func resolvePath(parent string, path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	pathBase := filepath.Dir(path)

	if pathBase == ".." {
		return filepath.Join(filepath.Dir(parent), filepath.Base(path))
	}

	if pathBase == "." {
		return filepath.Join(parent, filepath.Base(path))
	}

	return filepath.Join(parent, path)
}
