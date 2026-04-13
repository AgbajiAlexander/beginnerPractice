package main

import (
	"fmt"
	"strings"
)

// ToUpperLastWord applies ToUpper to the previous word when it finds (up), and removes (up) from the output.
func ToUpperLastWord(words []string) []string { //what does this line do? it takes a slice of strings and checks for the presence of (up). If it finds (up), it converts the previous word to uppercase and removes (up) from the output.
	var result []string               // This line initializes an empty slice of strings called result, which will be used to store the modified list of words after processing.
	for i := 0; i < len(words); i++ { // This line starts a for loop that iterates over the input slice of strings (words) using an index variable i. The loop continues until it has processed all elements in the words slice.
		if words[i] == "(up)" { // This line checks if the current word (words[i]) is equal to the string "(up)". If it is, it means that we need to apply the ToUpper function to the previous word in the result slice.
			if len(result) > 0 { // This line checks if the result slice has at least one element before trying to access the last element. This is important to avoid an index out of range error in case (up) is the first word in the input.
				result[len(result)-1] = strings.ToUpper(result[len(result)-1]) // This line takes the last word in the result slice (result[len(result)-1]) and applies the strings.ToUpper function to it, converting it to uppercase. The modified word is then assigned back to the last position in the result slice.
			}
			continue // skip adding (up)
		}
		result = append(result, words[i]) // If the current word is not "(up)", this line appends it to the result slice as is, without modification.
	}
	return result // After the loop has processed all words, this line returns the modified result slice, which contains the original words with any necessary transformations applied based on the presence of "(up)".
}

func main() {
	// Example tests
	inputs := [][]string{ //what does this do? it takes a slice of strings and checks for the presence of (up). If it finds (up), it converts the previous word to uppercase and removes (up) from the output.
		{"hello", "world", "(up)"},
		{"go", "is", "fun", "(up)"},
		{"hello", "(up)"},
	}
	for _, input := range inputs {
		fmt.Printf("Input: %v\nOutput: %v\n", input, ToUpperLastWord(input))
	}
}
