package handlers

import (
	"fmt"
	"strings"
)

/*
This function is responsible for converting a given input text into ASCII art after checking for special characters in the input . The function  calls the other function `PrintLineByLine`, to do the actual ASCII printing.
*/
func PrintAsciiArt(text string, asciiArt []string) {
	if text == "" {
		return
	}

	specialChars := map[string]string{
		"\\t": "Tab",
		"\\b": "Backspace",
		"\\v": "Vertical Tab",
		"\\0": "Null",
		"\\f": "Form Feed",
		"\\r": "Carriage Return",
	}

	for spChar, description := range specialChars {
		if strings.Contains(text, spChar) {
			fmt.Printf("Print Error: Special ASCII character (%s) or (%s) detected \n", spChar, description)
			return

		}
	}

	if strings.Contains(text, "\\n") {
		input := strings.Split(text, "\\n")
		count := 0
		for _, word := range input {
			if word == "" {
				count++
				if count < len(input) {
					fmt.Println()
				}
			} else if len(word) > 0 {
				PrintLineByLine(word, asciiArt)
			}
		}
	} else {

		lines := strings.Split(text, "\n")
		for _, line := range lines {
			if len(line) > 0 {
				PrintLineByLine(line, asciiArt)
			} else {
				fmt.Println()
			}
		}
	}
}
