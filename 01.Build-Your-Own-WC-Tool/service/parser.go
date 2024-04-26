package service

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// FetchInput gathers input from the user, parses and validates it.
func FetchInput() ([]string, error) {

	input, err := fetchInput()
	if err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}
	parsedString := parseInput(input)
	if err := validate(parsedString); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	return parsedString, nil
}

// validate checks the correctness of the parsed input according to the business rules.
func validate(parsedString []string) error {

	if len(parsedString) == 0 || parsedString[0] != "vwc" {
		return fmt.Errorf("first element must be 'vwc'")
	}

	if len(parsedString) > 1 {
		if err := validateFlags(parsedString[1]); err != nil {
			return err
		}
	}

	return validateCharacters(parsedString)
}

// validateFlags checks if the second element of parsedString is a valid flag or path.
func validateFlags(second string) error {
	allowedFlags := map[string]bool{"-c": true, "-l": true, "-m": true, "-w": true}
	if _, ok := allowedFlags[second]; !ok && !strings.Contains(second, ".") {
		return fmt.Errorf("invalid flag or path")
	}
	return nil
}

// validateCharacters ensures all strings in parsedString contain only valid characters.
func validateCharacters(strings []string) error {
	validChars := regexp.MustCompile(`^[a-zA-Z0-9./-]*$`)
	for _, str := range strings {
		if !validChars.MatchString(str) {
			return fmt.Errorf("invalid characters detected in %s", str)
		}
	}
	return nil
}

// fetchInput reads a line of input from stdin and handles possible I/O errors.
func fetchInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// parseInput splits the input string into words based on whitespace.
func parseInput(input string) []string {
	return strings.Fields(input)
}
