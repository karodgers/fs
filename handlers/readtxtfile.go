package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	asciiArtHeight = 8
)

//the function returns a slice of strings with the ascii art representation in its corresponding index as in the asii chart
func ReadAsciiArt() []string {
	var args []string

	args = append(args, os.Args[1:]...)

	filename := args[len(args)-1]

	if len(os.Args) == 2 {
		filename = "standard.txt"
	}

	switch args[len(args)-1] {
	case "thinkertoy":
		filename = "thinkertoy.txt"
	case "shadow":
		filename = "shadow.txt"
	case "standard":
		filename = "standard.txt"
	default:
		if !strings.HasSuffix(filename, ".txt") {
			filename += ".txt"
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error; specified banner doees not have a corresponding .txt file")
		os.Exit(1)
	}
	defer file.Close()

	lineCount := 0

	scanner := bufio.NewScanner(file)

	var asciiArt []string
	var artLines []string

	for scanner.Scan() {

		lineCount++

		lines := scanner.Text()

		if len(lines) == 0 {
			continue
		}
		artLines = append(artLines, lines)

		if len(artLines) == asciiArtHeight {
			asciiArt = append(asciiArt, strings.Join(artLines, "\n"))
			artLines = nil
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading ASCII art file:", scanner.Err())
		os.Exit(1)
	}
	if len(asciiArt) == 0 {
		fmt.Println("Error: The ASCII art file is empty.")
		os.Exit(1)
	}

	if lineCount != 855 {
		fmt.Printf("Read Error in line %v: Do not change the content of the txt file", lineCount)
		os.Exit(1)
	}

	return asciiArt
}
