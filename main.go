package main

import (
	"fmt"
	"os"
	"term_screen/parser"
	"term_screen/processor"
	"term_screen/reader"
)

func main() {
	byteQueue := reader.Read(os.Stdin)

	parser := parser.New(byteQueue)
	cmdQueue := parser.ParseQueue()

	processor := processor.New(cmdQueue)
	err := processor.ProcessCommands()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing commands: %v", err)
		os.Exit(1)
	}

	fmt.Println("Processing completed gracefully")
}
