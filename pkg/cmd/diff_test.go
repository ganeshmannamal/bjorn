package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDiffRun(t *testing.T) {
	inputs := [][]string{
		{"image1.png","image2.png"},
		{"image1.png","image5.png"},
	}

	err := createInputCsv(inputs)

	if err != nil {
		t.Error(err)
	}

	opts := &Opts{
		csvFile: "../../testdata/input.csv",
	}

	err = opts.Run()

	if err != nil {
		t.Error(err)
	}

	f, err := os.Open(opts.outFile)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 2, len(lines), "Output CSV should have same no of lines as input")

	for i, line := range lines {
		assert.Equal(t, 4, len(line), "Each line should have 4 columns")
		assert.Equal(t, inputs[i],[]string{line[0],line[1]}, "First 2 columns should be same is input list")
	}

	if len(lines) == 2 {
		line1 := lines[0]
		assert.Equal(t, fmt.Sprintf("%.2f", 0.0), line1[2], "Score for first image should be 0")

		line2 := lines[1]
		assert.Equal(t, fmt.Sprintf("%.2f", 1.0), line2[2], "Score for first image should be 1")
	}
}

func createInputCsv(inputs [][]string) error {
	inFile, err := os.Create("../../testdata/input.csv")
	if err != nil {
		return err
	}

	writer := csv.NewWriter(inFile)

	for _, input := range inputs {
		err = writer.Write(input)
		if err != nil {
			return err
		}
	}

	writer.Flush()

	err = inFile.Close()
	if err != nil {
		return err
	}

	return nil
}
