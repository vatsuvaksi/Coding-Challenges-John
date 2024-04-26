package service

// commandFactory takes an array of strings as input and returns a function based on the command flag.
func CommandFactory(input []string) CommandFunction {
	if len(input) < 2 {
		return func([]string) string {
			return "Insufficient arguments"
		}
	}

	switch input[1] {

	case "-c":
		return countBytes
	case "-l":
		return countLines
	case "-m":
		return countCharacters
	case "-w":
		return countWords
	default:
		return func([]string) string {
			return "Invalid command flag"
		}
	}
}
