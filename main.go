package main

import (
	"fmt"
	"strings"
)

func CleanInput(text string) (cleanedInput []string) {
	// No

	lowerCase := strings.ToLower(text)
	words := strings.Fields(lowerCase)
	for _, word := range words {
		cleanedInput = append(cleanedInput, strings.TrimSpace(word))
	}
	return cleanedInput
}

func main() {
	fmt.Println("Hello, World!")
}
