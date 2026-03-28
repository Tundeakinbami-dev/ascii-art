package ascii

import "strings"

// Render converts a text string into ASCII art using the provided banner map.
// Each character in the text is looked up in the banner and its 8-line
// ASCII representation is printed line-by-line.
func Render(text string, banner map[rune][]string) string {

	// strings.Builder is used for efficient string concatenation
	var result strings.Builder

	// ASCII art characters are 8 rows tall
	for row := 0; row < 8; row++ {

		// Iterate through each character in the input text
		for _, char := range text {

			// Check if the character exists in the banner map
			if asciiChar, ok := banner[char]; ok {

				// Append the corresponding row of the ASCII art character
				result.WriteString(asciiChar[row])
			}

		}

		// After finishing one row for all characters, move to the next line
		result.WriteString("\n")
	}

	// Return the final ASCII art string
	return result.String()
}