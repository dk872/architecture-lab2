package main

import (
	"flag"
	"fmt"
	lab2 "github.com/dk872/architecture-lab2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	file            = flag.String("f", "", "Input file containing the expression")
	output          = flag.String("o", "", "Output file for the result")
)

func main() {
	flag.Parse()

	if *inputExpression != "" && *file != "" {
		fmt.Fprintln(os.Stderr, "ERROR: Cannot use both -e and -f")
		os.Exit(1)
	}

	var input io.Reader
	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to open file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		input = f
	} else {
		fmt.Fprintln(os.Stderr, "ERROR: No input provided (-e or -f required)")
		os.Exit(1)
	}

	var outputWriter io.Writer
	var outputBuffer *strings.Builder

	if *output != "" {
		f, err := os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: Failed to create output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		outputWriter = f
	} else {
		outputBuffer = &strings.Builder{}
		outputWriter = outputBuffer
	}

	handler := &lab2.ComputeHandler{Input: input, Output: outputWriter}
	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}

	if output == nil || *output == "" {
		fmt.Println(outputWriter.(*strings.Builder).String())
	}
}
