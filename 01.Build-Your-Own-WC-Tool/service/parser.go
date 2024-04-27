package service

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// FetchInput gathers input from the user, parses and validates it.
func FetchInput() ([]string, bool, error) {

	var isStepFinal bool = false
	input, err := fetchInput()
	if err != nil {
		return nil, isStepFinal, fmt.Errorf("error reading input: %w", err)
	}
	parsedString := parseInput(input)

	// This is for handling final step
	if !(strings.Compare(parsedString[0], "cat") == 0) {
		if err := validate(parsedString); err != nil {
			return nil, isStepFinal, fmt.Errorf("validation error: %w", err)
		}
	} else if strings.Compare(parsedString[0], "cat") == 0 {
		isStepFinal = true
		if err := validateForFinalStep(parsedString); err != nil {
			return nil, isStepFinal, fmt.Errorf("validation error: %w", err)
		}
	}

	return parsedString, isStepFinal, nil
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

func validateForFinalStep(parsedString []string) error {
	// Check if parsedString length is sufficient for the final step
	if len(parsedString) < 4 {
		return fmt.Errorf("insufficient input for final step")
	}

	// Check if input[1] contains a dot (.)
	if !strings.Contains(parsedString[1], ".") {
		return fmt.Errorf("second element must contain a dot")
	}

	// Check if input[2] is a pipe (|)
	if parsedString[2] != "|" {
		return fmt.Errorf("third element must be a pipe (|)")
	}

	// Check if input[3] is 'vwc'
	if parsedString[3] != "vwc" {
		return fmt.Errorf("fourth element must be 'vwc'")
	}

	// handling for no flag case in final step

	if len(parsedString) == 4 {
		parsedString = append(parsedString, "place.holder")
	}
	// Check if input[4] is a valid flag or path
	if err := validateFlags(parsedString[4]); err != nil {
		return fmt.Errorf("invalid flag or path: %w", err)
	}

	return nil
}
