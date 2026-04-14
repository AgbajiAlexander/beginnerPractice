package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AllinOne(str string) string {
	ayz := strings.Fields(str)
	var char4 []string
	for _, char := range ayz {
		if strings.HasPrefix(char, "(") && strings.HasSuffix(char, ")") {
			char1 := char[1 : len(char)-1]
			words := strings.Split(char1, ",")
			count := 1
			char2 := strings.TrimSpace(words[0])
			if len(words) > 1 {
				char3, err := strconv.Atoi(strings.TrimSpace(words[1]))
				if err == nil && char3 > count {
					count = char3
				}

			}
			char5 := len(char4) - count
			if char5 < count {
				char5 = 0
			}
			for j := char5; j < len(char4); j++ {
				if char2 == "up" {
					char4[j] = strings.ToUpper(char4[j])
				}
				if char2 == "low" {
					char4[j] = strings.ToLower(char4[j])
				}
				if char4[j] == "hex" {
					strconv.ParseInt(char4[j], 16, 64)
				}
			}
		} else {

			char4 = append(char4, char)
		}
	}
	return strings.Join(char4, " ")
}

func main() {
	fmt.Println(AllinOne("learning 1E (hex) go (up) IS (low) fun"))
}
