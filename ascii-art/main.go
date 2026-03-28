package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/internal/ascii"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"text\"")
		return
	}

	input := os.Args[1]

	// Handle newline characters
	lines := strings.Split(input, "\\n")

	//call our standard function from standard.txt
	banner, err := ascii.LoadBanner("banners/thinkertoy.txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}
	//range index and value
	for _, line := range lines {

		if line == " " {
			fmt.Println()
			continue
		}

		output := ascii.Render(line, banner)
		fmt.Print(output)
	}
}
