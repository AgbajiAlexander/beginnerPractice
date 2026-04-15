package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(data)

	result := process(text)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func process(text string) string {
	words := strings.Fields(text)

	// STEP 1: Apply transformations
	for i := 0; i < len(words); i++ {

		// HANDLE (hex)
		if words[i] == "(hex)" && i > 0 {
			num, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = remove(words, i)
			i--
			continue
		}

		// HANDLE (bin)
		if words[i] == "(bin)" && i > 0 {
			num, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = remove(words, i)
			i--
			continue
		}

		// HANDLE (up), (low), (cap)
		if strings.HasPrefix(words[i], "(up") || // what does this do? it checks if the word starts with (up, which means it can be (up) or (up,2) etc.
			strings.HasPrefix(words[i], "(low") ||
			strings.HasPrefix(words[i], "(cap") {

			mode, count := parseInstruction(words[i]) // this function will return the mode (up, low, cap) and the count (number of words to transform)

			start := i - count
			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				switch mode {
				case "up":
					words[j] = strings.ToUpper(words[j])
				case "low":
					words[j] = strings.ToLower(words[j])
				case "cap":
					words[j] = capitalize(words[j])
				}
			}

			words = remove(words, i)
			i--
		}
	}

	// STEP 2: Join words
	result := strings.Join(words, " ")

	// STEP 3: Fix punctuation
	result = fixPunctuation(result)

	// STEP 4: Fix quotes
	result = fixQuotes(result)

	// STEP 5: Fix "a" → "an"
	result = fixArticles(result)

	return result
}

func parseInstruction(s string) (string, int) {
	s = strings.Trim(s, "()")

	parts := strings.Split(s, ",")

	mode := parts[0]

	if len(parts) == 2 {
		n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err == nil {
			return mode, n
		}
	}

	return mode, 1
}

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func fixPunctuation(text string) string {
	// Remove space before punctuation
	punct := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punct {
		text = strings.ReplaceAll(text, " "+p, p)
	}

	// Fix spacing after punctuation
	for _, p := range punct {
		text = strings.ReplaceAll(text, p, p+" ")
	}

	// Fix multiple spaces
	text = strings.Join(strings.Fields(text), " ")

	// Handle special groups like ... !?
	text = strings.ReplaceAll(text, ". . .", "...")
	text = strings.ReplaceAll(text, "! ?", "!?")

	return text
}

func fixQuotes(text string) string {
	parts := strings.Split(text, "'")

	for i := 1; i < len(parts); i += 2 {
		parts[i] = strings.TrimSpace(parts[i])
	}

	return strings.Join(parts, "'")
}

func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			r := rune(words[i+1][0])
			if isVowel(r) {
				if unicode.IsUpper(rune(words[i][0])) {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiouhAEIOUH", r)
}
