package handlers

import "fmt"

func ErrorMsg() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . something standard")
}
