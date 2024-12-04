package service

import (
	"fmt"
	"os"
	"strings"
)

// CommandFunction represents a function type that takes an array of strings as input and returns nothing.
type CommandFunction func([]string) string

// countCharacters counts the number of characters in the specified file.
func countCharacters(input []string) string {
	filePath := input[2]
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("Error reading file: %v", err)
	}
	return fmt.Sprintf("Number of characters in file '%s': %d", filePath, len(content))
}

// countLines counts the number of lines in the specified file.
func countLines(input []string) string {
	filePath := input[2]
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("Error reading file: %v", err)
	}
	lines := strings.Split(string(content), "\n")
	return fmt.Sprintf("Number of lines in file '%s': %d", filePath, len(lines))
}

// countWords counts the number of words in the specified file.
func countWords(input []string) string {
	filePath := input[2]
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("Error reading file: %v", err)
	}
	words := strings.Fields(string(content))
	return fmt.Sprintf("Number of words in file '%s': %d", filePath, len(words))
}

// countBytes counts the number of bytes in the specified file.
func countBytes(input []string) string {
	filePath := input[2]
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("Error reading file: %v", err)
	}
	return fmt.Sprintf("Number of bytes in file '%s': %d", filePath, len(content))
}

func printAll(input []string) string {
	input = append(input, input[1])
	return countCharacters(input) + "\n" + countBytes(input) + "\n" + countLines(input) + "\n" + countWords(input)
}
