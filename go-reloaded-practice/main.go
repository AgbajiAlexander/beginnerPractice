package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check if correct number of arguments are passed
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	// Get input and output file names
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read file content
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert file content to string
	text := string(content)

	// Process the text
	result := processText(text)

	// Write result to output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func processText(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		switch {
		// HEX conversion
		case words[i] == "(hex)" && i > 0:
			num, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = remove(words, i)
			i--

		// BIN conversion
		case words[i] == "(bin)" && i > 0:
			num, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = remove(words, i)
			i--

		// UPPERCASE
		case words[i] == "(up)" && i > 0:
			words[i-1] = strings.ToUpper(words[i-1])
			words = remove(words, i)
			i--

		// LOWERCASE
		case words[i] == "(low)" && i > 0:
			words[i-1] = strings.ToLower(words[i-1])
			words = remove(words, i)
			i--

		// CAPITALIZE
		case words[i] == "(cap)" && i > 0:
			words[i-1] = capitalize(words[i-1])
			words = remove(words, i)
			i--
		}
	}

	// Fix punctuation spacing
	result := strings.Join(words, " ")
	result = fixPunctuation(result)

	// Fix "a" to "an"
	result = fixArticles(result)

	return result
}

// Remove a word from slice
func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// Capitalize first letter
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// Fix punctuation spacing
func fixPunctuation(text string) string {
	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuations {
		text = strings.ReplaceAll(text, " "+p, p)
	}

	return text
}

// Fix "a" to "an"
func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			first := rune(words[i+1][0])
			if isVowel(first) {
				words[i] = "an"
			}
		}
	}

	return strings.Join(words, " ")
}

// Check if vowel
func isVowel(r rune) bool {
	vowels := "aeiouhAEIOUH"
	return strings.ContainsRune(vowels, r)
}
