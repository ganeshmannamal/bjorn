package main

import (
	"encoding/csv"
	"github.com/ganeshmannamal/bjorn/pkg/cmd"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	for i, arg := range os.Args {
		if arg == "-test.main" {
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
			main()
			return
		}
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCommand(t *testing.T)  {
	if os.Getenv("RUNME") == "1" {
		cmd.Execute()
		return
	}

	inputs := [][]string{
		{"image1.png","image2.png"},
		{"image1.png","image5.png"},
	}

	inFile, err := os.Create("../../testdata/input.csv")
	if err != nil {
		t.Error(err)
	}
	defer inFile.Close()

	writer := csv.NewWriter(inFile)
	defer writer.Flush()

	for _, input := range inputs {
		err = writer.Write(input)
		if err != nil {
			t.Error(err)
		}
	}

	args := []string{
		"-test.run=TestCommand",
		"diff",
		"--file",
		"../../testdata/test.csv",
	}
	c := exec.Command(os.Args[0], args...)
	c.Env = append(os.Environ(), "RUNME=1")
	out, err := c.CombinedOutput()
	if err != nil {
		t.Log(string(out))
		t.Error(err)
	}

	if !fileExists("../../testdata/output.csv") {
		t.Errorf("output file was not generated")
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err) // some other race condition, we panic!!
	}
	return true
}
