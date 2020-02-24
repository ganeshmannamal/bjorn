package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"github.com/ganeshmannamal/bjorn/pkg/pair"
	"github.com/ganeshmannamal/bjorn/pkg/util"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type Opts struct {
	csvFile string
	outFile string
}

// NewDiffCommand generates a cobra command for executing the diff command.
// This command accepts the flags --file (-f path to csv file) and
// --out (-o path for output file). --out is optional, if not provided,
// out put file is created in same directory as input csv
func NewDiffCommand() *cobra.Command {
	opts := &Opts{}
	diffCmd := &cobra.Command{
		Use:   "diff",
		Short: "compare images listed in a csv file",
		Long:  `run bjorn diff to compare images listed in a csv file`,
		Run: func(cmd *cobra.Command, args []string) {
			err := opts.Run()
			if err != nil {
				ExitWithError("diff", err)
			}
			color.New(color.Bold, color.FgGreen).Fprintf(os.Stdout, "\nDIff written to %s\n", opts.outFile)
		},
	}

	diffCmd.Flags().StringVarP(&opts.csvFile, "file", "f", "", "CSV file to read image list")
	diffCmd.Flags().StringVarP(&opts.outFile, "out", "o", "", "Output file location")
	err := diffCmd.MarkFlagRequired("file")
	if err != nil {
		ExitWithError("diff", err)
	}

	return diffCmd
}

// Run reads the csv file and parses each line for image pairs for comparison.
// After comparison it writes the to output csv, with each line containing -
// "image1,image2,score,time taken
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
	defer out.Close()

	writer := csv.NewWriter(out)
	defer writer.Flush()

	for _, line := range lines {
		img1Path := util.ResolvePath(csvRootPath, line[0])
		img2Path := util.ResolvePath(csvRootPath, line[1])

		p, err := pair.NewImagePair(img1Path, img2Path)

		if err != nil {
			return err
		}

		p.Compare()
		err = writer.Write([]string{line[0], line[1], fmt.Sprintf("%.2f", p.Score), fmt.Sprintf("%f", p.Time)})
	}
	return nil
}
