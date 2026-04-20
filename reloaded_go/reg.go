package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Read file content
func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return string(data)
}

// Write output to file
func WriteFile(filename, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}

// Core text processing using regex
func ProcessText(text string) string {
	lines := strings.Split(text, "\n")

	// Regex patterns
	multiSpace := regexp.MustCompile(`\s+`)
	spaceBeforePunct := regexp.MustCompile(`\s+([.,!?;:])`)
	spaceInsideSingleQuotes := regexp.MustCompile(`'\s*(.*?)\s*'`)
	spaceInsideDoubleQuotes := regexp.MustCompile(`"\s*(.*?)\s*"`)

	var result []string

	for _, line := range lines {
		// Step 1: Trim outer spaces
		line = strings.TrimSpace(line)

		// Step 2: Collapse multiple spaces
		line = multiSpace.ReplaceAllString(line, " ")

		// Step 3: Remove space before punctuation
		line = spaceBeforePunct.ReplaceAllString(line, "$1")

		// Step 4: Clean inside single quotes
		line = spaceInsideSingleQuotes.ReplaceAllString(line, "'$1'")

		// Step 5: Clean inside double quotes
		line = spaceInsideDoubleQuotes.ReplaceAllString(line, "\"$1\"")

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text := ReadFile(inputFile)
	result := ProcessText(text)
	WriteFile(outputFile, result)
}
