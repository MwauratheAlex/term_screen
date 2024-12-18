package main

import (
	"fmt"
	"os"
	"term_screen/parser"
	"term_screen/processor"
	"term_screen/reader"
	"term_screen/ui"
)

func main() {
	byteQueue := reader.Read(os.Stdin)

	parser := parser.New(byteQueue)
	cmdQueue := parser.ParseQueue()

	screen := ui.NewScreen()
	processor := processor.New(cmdQueue, screen)

	err := processor.ProcessCommands()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing commands: %v", err)
		os.Exit(1)
	}
}
