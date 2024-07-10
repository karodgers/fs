package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/handlers"
)

// check the correct number of argument passed before processing to
//call the functions in handlers package
func main() {
	if len(os.Args) == 1 {
		handlers.ErrorMsg()
		os.Exit(0)
	}
	text := strings.Join(os.Args[1:len(os.Args)-1], " ")

	if len(os.Args) == 2 {
		text = os.Args[1]
	}
	if len(os.Args) > 3 {
		handlers.ErrorMsg()
		os.Exit(0)
	}

	asciiArt := handlers.ReadAsciiArt()

	if handlers.ContainsNonASCII(text) {
		fmt.Println("Error: Non-ASCII characters detected")
		os.Exit(1)
	}
	handlers.PrintAsciiArt(text, asciiArt)
}
