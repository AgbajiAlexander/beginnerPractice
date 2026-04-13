package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile((inputFile))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(content)

	result := processLexText(text)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("error writing files:", err)
		return
	}
}

func processLexText(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		switch {
		case words[i] == "(hex)" && i > 0:
			num, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = remove(words, i)
			i--

		case words[i] == "(bin)" && i > 0:
			num, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = strconv.Itoa(int(num))
			words = remove(words, i)
			i--

		case words[i] == "(up)" && i > 0:
			words[i-1] = strings.ToUpper(words[i-1])
			words = remove(words, i)
			i--

		case words[i] == "(low)" && i > 0:
			words[i-1] = strings.ToLower(words[i-1])
			words = remove(words, i)
			i--

		case words[i] == "(cap)" && i > 0:
			words[i-1] = strings.Title(words[i-1])
			words = remove(words, i)
			i--
		}
	}
	result := strings.Join(words, " ")
	result = fixPunctuation(result)

	result = fixArticles(result)
	return result

}

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func fixPunctuation(text string) string {
	punctuations := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range punctuations {
		text = strings.ReplaceAll(text, " "+p, p)
	}
	return text
}

func fixArticles(text string) string {
	text = strings.ReplaceAll(text, " a ", " an ")
	text = strings.ReplaceAll(text, " A ", " An ")
	return text
}
