package ascii

import (
	"fmt"
	"os"
	"strings"
)

// LoadBanner reads a banner file and converts it into a map.
// Each printable ASCII character (32–126) is mapped to its
// corresponding 8-line ASCII-art representation.
func LoadBanner(path string) (map[rune][]string, error) {

	// Read the entire banner file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Split the file into individual lines
	lines := strings.Split(string(data), "\n")

	// Basic validation: banner files should contain many lines
	if len(lines) < 800 {
		return nil, fmt.Errorf("banner file seems invalid")
	}

	// Create a map to store ASCII characters and their 8-line drawings
	banner := make(map[rune][]string)

	// Start reading from line 1 because the first line is empty
	index := 1

	// Loop through printable ASCII characters (space to ~)
	for i := 32; i <= 126; i++ {

		// Each character occupies 8 lines in the banner file
		charLines := lines[index : index+8]

		// Store the character and its ASCII-art lines in the map
		banner[rune(i)] = charLines

		// Move index to the next character block (8 lines + 1 empty line)
		index += 9
	}

	// Return the completed banner map
	return banner, nil
}